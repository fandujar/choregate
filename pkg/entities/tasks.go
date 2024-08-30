package entities

import (
	"time"

	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

type TaskConfig struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Namespace   string        `json:"namespace"`
	Description string        `json:"description"`
	Timeout     time.Duration `json:"timeout"`
	*TaskScope
	*tekton.TaskSpec
}

type Task struct {
	*TaskConfig
}

type TaskScope struct {
	Owner         uuid.UUID       `json:"owner"`         // Owner of the task
	Organizations []*Organization `json:"organizations"` // Organization that can access the task
	Teams         []*Team         `json:"teams"`         // Team that can access the task
	Users         []*User         `json:"users"`         // User that can access the task
}

// NewTask creates a new Task with the given configuration and default values.
func NewTask(config *TaskConfig) (*Task, error) {
	var err error

	if config.ID == uuid.Nil {
		config.ID, err = utils.GenerateID()
		if err != nil {
			return nil, err
		}
	}

	if config.Namespace == "" {
		config.Namespace = "choregate"
	}

	if config.Timeout == 0 {
		config.Timeout = 5 * time.Minute
	}

	if config.TaskSpec == nil {
		config.TaskSpec = &tekton.TaskSpec{}
	}

	return &Task{
		TaskConfig: config,
	}, nil
}
