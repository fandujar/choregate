package main

import (
	"net/http"
	"time"

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
	var userRepository repositories.UserRepository
	var triggerRepository repositories.TriggerRepository
	var teamRepository repositories.TeamRepository

	if true {
		// Create a memory repository
		taskRepository = memory.NewInMemoryTaskRepository()
		userRepository = memory.NewInMemoryUserRepository()
		triggerRepository = memory.NewInMemoryTriggerRepository()
		teamRepository = memory.NewInMemoryTeamRepository()
	}

	// Print the type of each repository being used
	log.Debug().Msgf("type of taskRepository: %T", taskRepository)
	log.Debug().Msgf("type of userRepository: %T", userRepository)
	log.Debug().Msgf("type of triggerRepository: %T", triggerRepository)
	log.Debug().Msgf("type of teamRepository: %T", teamRepository)

	// Create services
	taskService := services.NewTaskService(taskRepository)

	// Register the routes
	transport.RegisterTasksRoutes(r, *taskService)

	http.ListenAndServe(":8080", r)
}
