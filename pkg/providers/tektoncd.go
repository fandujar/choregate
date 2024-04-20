package providers

import (
	"context"
	"fmt"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
)

// TektonClient is a client for interacting with Tekton.
type TektonClient interface {
	// GetTaskRun returns a task run by name.
	GetTaskRun(ctx context.Context, id uuid.UUID) (*entities.TaskRun, error)
	// RunTask runs a task.
	RunTask(ctx context.Context, task *entities.TaskRun) error
}

type TektonClientImpl struct{}

// NewTektonClient creates a new TektonClient.
func NewTektonClient() TektonClient {
	return &TektonClientImpl{}
}

func (c *TektonClientImpl) GetTaskRun(ctx context.Context, id uuid.UUID) (*entities.TaskRun, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *TektonClientImpl) RunTask(ctx context.Context, task *entities.TaskRun) error {
	return fmt.Errorf("not implemented")
}
