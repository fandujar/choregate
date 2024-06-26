package entities

import (
	"time"

	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

type TaskConfig struct {
	ID          uuid.UUID          `json:"id"`
	Title       string             `json:"title"`
	Namespace   string             `json:"namespace"`
	Description string             `json:"description"`
	Timeout     time.Duration      `json:"timeout"`
	Steps       []tekton.Step      `json:"steps"`
	Params      []tekton.ParamSpec `json:"params"`
}

type Task struct {
	*TaskConfig
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

	return &Task{
		TaskConfig: config,
	}, nil
}
