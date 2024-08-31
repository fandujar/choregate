package controller

import (
	"context"
	"fmt"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/providers/tektoncd"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/google/uuid"
	tektonAPI "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/rs/zerolog/log"
)

type Controller struct {
	*ControllerConfig
}

type ControllerConfig struct {
	TektonCD tektoncd.TektonClient
	Service  *services.TaskService
}

func NewController(config *ControllerConfig) (*Controller, error) {
	return &Controller{
		ControllerConfig: config,
	}, nil
}

// Run starts the controller and watchs for tektonCD tasks events
func (c *Controller) Run(ctx context.Context) error {
	events, err := c.TektonCD.WatchTasks(ctx, "choregate")
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-events:
			if err := c.HandleEvent(ctx, event); err != nil {
				log.Error().Err(err).Msg("failed to handle event")
			}
		}
	}
}

func (c *Controller) HandleEvent(ctx context.Context, event watch.Event) error {
	switch event.Type {
	case "ADDED":
		tektonTask := event.Object.(*tektonAPI.Task)
		log.Info().Msgf("task %s added", tektonTask.Name)

		task, err := entities.NewTask(
			&entities.TaskConfig{
				Name:     tektonTask.Name,
				TaskSpec: &tektonTask.Spec,
			},
		)
		if err != nil {
			return err
		}

		err = c.Service.Create(ctx, task)
		if err != nil {
			return err
		}

		c.TektonCD.SetTaskLabels(ctx, tektonTask, map[string]string{
			"choregate.fandujar.dev/task-id": task.ID.String(),
		})
	case "MODIFIED":
		tektonTask := event.Object.(*tektonAPI.Task)
		log.Info().Msgf("task %s modified", tektonTask.Name)

		taskID := tektonTask.Labels["choregate.fandujar.dev/task-id"]
		if taskID == "" {
			return fmt.Errorf("task %s has no task-id label", tektonTask.Name)
		}

		id := uuid.MustParse(taskID)

		task, err := entities.NewTask(
			&entities.TaskConfig{
				ID:       id,
				Name:     tektonTask.Name,
				TaskSpec: &tektonTask.Spec,
			},
		)

		if err != nil {
			return err
		}

		err = c.Service.Update(ctx, task)
		if err != nil {
			return err
		}
	case "DELETED":
		tektonTask := event.Object.(*tektonAPI.Task)
		log.Info().Msgf("task %s deleted", tektonTask.Name)

		taskID := tektonTask.Labels["choregate.fandujar.dev/task-id"]
		if taskID == "" {
			return fmt.Errorf("task %s has no task-id label", tektonTask.Name)
		}

		id := uuid.MustParse(taskID)

		task, err := entities.NewTask(
			&entities.TaskConfig{
				ID:       id,
				Name:     tektonTask.Name,
				TaskSpec: &tektonTask.Spec,
			},
		)

		if err != nil {
			return err
		}

		err = c.Service.Delete(ctx, task.ID)
		if err != nil {
			return err
		}
	default:
		log.Error().Msgf("unknown event type: %s", event.Type)
	}

	return nil
}
