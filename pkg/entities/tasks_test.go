package entities_test

import (
	"testing"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	task := entities.Task{
		&entities.TaskConfig{
			ID:    uuid.New(),
			Title: "Test Task",
		},
	}

	assert.Equal(t, "Test Task", task.Title)
}

func TestTaskID(t *testing.T) {
	// Test task with input ID
	task, err := entities.NewTask(
		&entities.TaskConfig{
			ID:    uuid.New(),
			Title: "Test Task",
		},
	)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, task.ID)

	// Test task without ID
	task2, err := entities.NewTask(
		&entities.TaskConfig{
			Title: "Test Task without ID",
		},
	)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, task2.ID)
}
