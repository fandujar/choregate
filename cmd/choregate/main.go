package main

import (
	"context"
	"embed"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fandujar/choregate/pkg/providers/auth"
	"github.com/fandujar/choregate/pkg/providers/tektoncd"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/fandujar/choregate/pkg/repositories/memory"
	"github.com/fandujar/choregate/pkg/repositories/postgres"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/fandujar/choregate/pkg/transport"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed index.html assets/*
var choregateUIFS embed.FS

func main() {
	// Configure the logger level and format
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = time.RFC3339Nano

	log.Info().Msg("starting choregate")

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool { return true },
		AllowedMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	// Create repositories
	// Check which type of repository is being used
	// implemented types: memory

	var taskRepository repositories.TaskRepository
	var taskRunRepository repositories.TaskRunRepository
	var userRepository repositories.UserRepository
	var triggerRepository repositories.TriggerRepository
	var teamRepository repositories.TeamRepository

	repositoryType := os.Getenv("CHOREGATE_REPOSITORY_TYPE")
	if repositoryType == "postgres" {
		var err error
		ctx := context.Background()

		// Create a postgres repository
		taskRepository, err = postgres.NewPostgresTaskRepository(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create postgres task repository")
		}
	}

	if repositoryType == "memory" {
		// Create a memory repository
		taskRepository = memory.NewInMemoryTaskRepository()
		taskRunRepository = memory.NewInMemoryTaskRunRepository()
		userRepository = memory.NewInMemoryUserRepository()
		triggerRepository = memory.NewInMemoryTriggerRepository()
		teamRepository = memory.NewInMemoryTeamRepository()
	}

	// Print the type of each repository being used
	log.Debug().Msgf("type of taskRepository: %T", taskRepository)
	log.Debug().Msgf("type of userRepository: %T", userRepository)
	log.Debug().Msgf("type of triggerRepository: %T", triggerRepository)
	log.Debug().Msgf("type of teamRepository: %T", teamRepository)

	// Initialize tekton client
	tektonClient, err := tektoncd.NewTektonClient()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create Tekton client")
	}

	// Create services
	taskService := services.NewTaskService(taskRepository, taskRunRepository, tektonClient)
	triggerService := services.NewTriggerService(triggerRepository)
	userService := services.NewUserService(userRepository)

	// Create the auth provider
	authProvider, err := auth.NewAuthProvider(userService)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create auth provider")
	}
	tokenAuth := authProvider.NewTokenAuth()

	// Serve the UI
	r.Route("/", func(r chi.Router) {
		r.Use(
			CustomLogger(),
			middleware.RequestID,
			middleware.RealIP,
			middleware.Recoverer,
		)
		r.Handle("/*", http.FileServer(http.FS(choregateUIFS)))
	})

	// Handle login
	r.Route("/user", func(r chi.Router) {
		r.Use(
			CustomLogger(),
			middleware.RequestID,
			middleware.RealIP,
			middleware.Recoverer,
		)
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			token, err := authProvider.HandleLogin(
				r.Context(),
				r.FormValue("username"),
				r.FormValue("password"),
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"token": token,
			})
		})
		// r.Post("/logout", func(w http.ResponseWriter, r *http.Request) {

		// })
	})

	// Register the routes
	r.Route("/api", func(r chi.Router) {
		r.Use(
			CustomLogger(),
			middleware.RequestID,
			jwtauth.Verifier(tokenAuth),
			jwtauth.Authenticator(tokenAuth),
			middleware.RealIP,
			middleware.Recoverer,
		)
		transport.RegisterTasksRoutes(r, *taskService)
		transport.RegisterTriggersRoutes(r, *triggerService)
		transport.RegisterUsersRoutes(r, *userService)
	})

	// Prepare to handle signals
	// Start the HTTP server
	shutdown := make(chan bool, 1)
	signalHandler := make(chan os.Signal, 1)
	signal.Notify(signalHandler, os.Interrupt, syscall.SIGTERM)

	go func() {
		s := <-signalHandler
		log.Info().Msgf("received signal: %v", s)
		shutdown <- true
	}()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("server error")
			signalHandler <- syscall.SIGTERM
		}
	}()

	// Wait for a signal
	<-shutdown
	log.Info().Msg("shutting down server")
}

func CustomLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()
			defer func() {
				t2 := time.Now()
				log.Info().Msgf(
					"%s - url: %s - from: %s - %d - %d in %s",
					r.Method, r.URL, r.RemoteAddr, ww.Status(), ww.BytesWritten(), t2.Sub(t1),
				)

			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
