package entities

import (
	"github.com/fandujar/choregate/pkg/utils"
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

type TaskRunConfig struct {
	ID     uuid.UUID     `json:"id"`
	TaskID uuid.UUID     `json:"task_id"`
	Status TaskRunStatus `json:"status"`
}

type TaskRun struct {
	*TaskRunConfig
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

	return &Task{
		TaskConfig: config,
	}, nil
}

// NewTaskRun creates a new TaskRun with the given configuration and default values.
func NewTaskRun(config *TaskRunConfig) (*TaskRun, error) {
	var err error

	if config.ID == uuid.Nil {
		config.ID, err = utils.GenerateID()
		if err != nil {
			return nil, err
		}
	}

	return &TaskRun{
		TaskRunConfig: config,
	}, nil
}
