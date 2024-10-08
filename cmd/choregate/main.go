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

	"github.com/fandujar/choregate/pkg/controller"
	"github.com/fandujar/choregate/pkg/entities"
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

func getLogLevel() zerolog.Level {
	switch os.Getenv("CHOREGATE_LOG_LEVEL") {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}

func main() {
	// Configure the logger level and format
	logLevel := getLogLevel()
	zerolog.SetGlobalLevel(logLevel)
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
	var organizationRepository repositories.OrganizationRepository

	repositoryType := os.Getenv("CHOREGATE_REPOSITORY_TYPE")
	if repositoryType == "" {
		repositoryType = "memory"
	}

	if repositoryType == "postgres" {
		var err error
		ctx := context.Background()

		// Create a postgres repository
		taskRepository, err = postgres.NewPostgresTaskRepository(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create postgres task repository")
		}
		taskRunRepository, err = postgres.NewPostgresTaskRunRepository(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create postgres task run repository")
		}
		userRepository, err = postgres.NewPostgresUserRepository(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create postgres user repository")
		}
		triggerRepository = memory.NewInMemoryTriggerRepository()
		teamRepository = memory.NewInMemoryTeamRepository()
		organizationRepository = memory.NewInMemoryOrganizationRepository()
	}

	if repositoryType == "memory" {
		// Create a memory repository
		taskRepository = memory.NewInMemoryTaskRepository()
		taskRunRepository = memory.NewInMemoryTaskRunRepository()
		userRepository = memory.NewInMemoryUserRepository()
		triggerRepository = memory.NewInMemoryTriggerRepository()
		teamRepository = memory.NewInMemoryTeamRepository()
		organizationRepository = memory.NewInMemoryOrganizationRepository()
	}

	// Print the type of each repository being used
	log.Debug().Msgf("type of taskRepository: %T", taskRepository)
	log.Debug().Msgf("type of userRepository: %T", userRepository)
	log.Debug().Msgf("type of triggerRepository: %T", triggerRepository)
	log.Debug().Msgf("type of teamRepository: %T", teamRepository)
	log.Debug().Msgf("type of organizationRepository: %T", organizationRepository)

	// Initialize tekton client
	tektonClient, err := tektoncd.NewTektonClient()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create Tekton client")
	}

	// Create services
	taskService := services.NewTaskService(
		taskRepository,
		taskRunRepository,
		tektonClient,
	)
	triggerService := services.NewTriggerService(triggerRepository)
	organizationService := services.NewOrganizationService(
		organizationRepository,
		teamRepository,
		userRepository,
	)

	// Create the auth provider
	authProvider, err := auth.NewAuthProvider(organizationService)
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

	// Health check
	r.Group(func(r chi.Router) {
		r.Get("/liveness", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.Get("/readiness", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
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
			user, token, err := authProvider.HandleLogin(
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
				"user_id":     user.ID.String(),
				"username":    user.Email,
				"system_role": user.SystemRole,
				"email":       user.Email,
				"token":       token,
			})
		})
		r.Post("/validate", func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			token = token[7:]
			if token == "" {
				http.Error(w, "missing token", http.StatusBadRequest)
				return
			}

			valid, err := authProvider.ValidateToken(r.Context(), token)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if !valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			json.NewEncoder(w).Encode(map[string]bool{
				"valid": valid,
			})
		})
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
		transport.RegisterUsersRoutes(r, *organizationService)
		transport.RegisterTeamsRoutes(r, *organizationService)
		transport.RegisterOrganizationsRoutes(r, *organizationService)
	})

	// Setup SuperUser if environment variable is set
	err = SetupSuperUser(*organizationService)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to setup superuser")
	}

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

	controller, err := controller.NewController(
		&controller.ControllerConfig{
			Service:  taskService,
			TektonCD: tektonClient,
		},
	)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to create controller")
	}

	go func() {
		if err := controller.Run(context.Background()); err != nil {
			log.Error().Err(err).Msg("controller error")
			signalHandler <- syscall.SIGTERM
		}
	}()

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

func SetupSuperUser(service services.OrganizationService) error {
	superUserEmail := os.Getenv("CHOREGATE_SUPERUSER_EMAIL")
	superUserPassword := os.Getenv("CHOREGATE_SUPERUSER_PASSWORD")

	if superUserEmail == "" || superUserPassword == "" {
		log.Info().Msg("no superuser credentials provided")
		return nil
	}

	_, err := service.GetUserByEmail(context.Background(), superUserEmail)
	if err == nil {
		log.Info().Msg("superuser already exists")
		return nil
	}

	user, err := entities.NewUser(
		&entities.UserConfig{
			Email:      superUserEmail,
			Password:   superUserPassword,
			SystemRole: "admin",
		},
	)

	if err != nil {
		return err
	}

	err = service.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}

	team, err := entities.NewTeam(
		&entities.TeamConfig{
			Name: "superuser",
		},
	)

	if err != nil {
		return err
	}

	err = service.CreateTeam(context.Background(), team)
	if err != nil {
		return err
	}

	organization, err := entities.NewOrganization(
		&entities.OrganizationConfig{
			Name: "superuser",
		},
	)

	if err != nil {
		return err
	}

	err = service.CreateOrganization(context.Background(), organization)
	if err != nil {
		return err
	}

	err = service.AddOrganizationMember(context.Background(), organization.ID, user.ID, "admin")
	if err != nil {
		return err
	}

	err = service.AddTeamMember(context.Background(), team.ID, user.ID, "admin")
	if err != nil {
		return err
	}

	err = service.AddTeam(context.Background(), organization.ID, team.ID)
	if err != nil {
		return err
	}

	return nil
}
