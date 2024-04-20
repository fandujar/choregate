package main

import (
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
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Configure the logger level and format
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = time.RFC3339Nano

	log.Info().Msg("starting choregate")

	r := chi.NewRouter()

	// Create repositories
	// Check which type of repository is being used
	// implemented types: memory

	var taskRepository repositories.TaskRepository
	var taskRunRepository repositories.TaskRunRepository
	var userRepository repositories.UserRepository
	var triggerRepository repositories.TriggerRepository
	var teamRepository repositories.TeamRepository

	if true {
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
	tektonClient := providers.NewTektonClient()

	// Create services
	taskService := services.NewTaskService(taskRepository, taskRunRepository, tektonClient)
	triggerService := services.NewTriggerService(triggerRepository)

	// Register the routes
	transport.RegisterTasksRoutes(r, *taskService)
	transport.RegisterTriggersRoutes(r, *triggerService)

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
