package entities

import (
	"github.com/google/uuid"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

type TaskRunConfig struct {
	ID     string
	TaskID uuid.UUID
	*tekton.TaskRun
}

type TaskRun struct {
	*TaskRunConfig
}

// NewTaskRun creates a new TaskRun with the given configuration and default values.
func NewTaskRun(config *TaskRunConfig) (*TaskRun, error) {
	if config.ID == "" {
		config.ID = config.TaskRun.Name
	}

	if config.TaskRun.Labels == nil {
		config.TaskRun.Labels = make(map[string]string, 2)
		// Add the task ID and taskRun ID to the labels.
		config.TaskRun.Labels["choregate.fandujar.dev/task-id"] = config.TaskID.String()
		config.TaskRun.Labels["choregate.fandujar.dev/taskrun-id"] = config.ID
	}

	return &TaskRun{
		TaskRunConfig: config,
	}, nil
}
