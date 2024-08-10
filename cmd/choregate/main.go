package main

import (
	"embed"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fandujar/choregate/pkg/providers"
	"github.com/fandujar/choregate/pkg/repositories"
	"github.com/fandujar/choregate/pkg/repositories/memory"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/fandujar/choregate/pkg/transport"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed ui/*
var choregateUIFS embed.FS

// choregateUI is a handler that serves the UI for Choregate static files
func choregateUI(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(choregateUIFS)).ServeHTTP(w, r)
}

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

	// Serve the UI
	r.Get("/*", choregateUI)

	// Create repositories
	// Check which type of repository is being used
	// implemented types: memory

	var taskRepository repositories.TaskRepository
	var taskRunRepository repositories.TaskRunRepository
	var userRepository repositories.UserRepository
	var triggerRepository repositories.TriggerRepository
	var teamRepository repositories.TeamRepository
	var sessionsRepository repositories.SessionsRepository

	if true {
		// Create a memory repository
		taskRepository = memory.NewInMemoryTaskRepository()
		taskRunRepository = memory.NewInMemoryTaskRunRepository()
		userRepository = memory.NewInMemoryUserRepository()
		triggerRepository = memory.NewInMemoryTriggerRepository()
		teamRepository = memory.NewInMemoryTeamRepository()
		sessionsRepository = memory.NewInMemorySessionsRepository()
	}

	// Print the type of each repository being used
	log.Debug().Msgf("type of taskRepository: %T", taskRepository)
	log.Debug().Msgf("type of userRepository: %T", userRepository)
	log.Debug().Msgf("type of triggerRepository: %T", triggerRepository)
	log.Debug().Msgf("type of teamRepository: %T", teamRepository)
	log.Debug().Msgf("type of sessionsRepository: %T", sessionsRepository)

	// Initialize tekton client
	tektonClient, err := providers.NewTektonClient()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create Tekton client")
	}

	// Create services
	taskService := services.NewTaskService(taskRepository, taskRunRepository, tektonClient)
	triggerService := services.NewTriggerService(triggerRepository)
	userService := services.NewUserService(userRepository)

	// Register the routes
	r.Route("/api", func(r chi.Router) {
		r.Use(
			middleware.RequestID,
			CustomLogger(),
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
