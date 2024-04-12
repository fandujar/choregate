package entities

import (
	"github.com/google/uuid"
)

type TaskConfig struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type Task struct {
	*TaskConfig
}

// NewTask creates a new Task with the given configuration and default values.
func NewTask(config *TaskConfig) (*Task, error) {
	var err error

	if config.ID == uuid.Nil {
		config.ID, err = generateID()
		if err != nil {
			return nil, err
		}
	}

	return &Task{
		TaskConfig: config,
	}, nil
}
