package entities

import (
	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
)

type TaskRunConfig struct {
	ID     uuid.UUID     `json:"id"`
	TaskID uuid.UUID     `json:"task_id"`
	Status TaskRunStatus `json:"status"`
}

type TaskRun struct {
	*TaskRunConfig
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
