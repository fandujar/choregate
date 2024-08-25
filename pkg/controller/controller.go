package controller

import (
	"context"

	"github.com/fandujar/choregate/pkg/providers/tektoncd"
	tektonAPI "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"

	"github.com/rs/zerolog/log"
)

type Controller struct {
	*ControllerConfig
}

type ControllerConfig struct {
	TektonCD tektoncd.TektonClient
}

func NewController() (*Controller, error) {
	client, err := tektoncd.NewTektonClient()
	if err != nil {
		return nil, err
	}
	return &Controller{
		ControllerConfig: &ControllerConfig{
			TektonCD: client,
		},
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
			switch event.Type {
			case "ADDED":
				log.Info().Msgf("Task %s added", event.Object.(*tektonAPI.Task).Name)
			case "MODIFIED":
				log.Info().Msgf("Task %s modified", event.Object.(*tektonAPI.Task).Name)
			case "DELETED":
				log.Info().Msgf("Task %s deleted", event.Object.(*tektonAPI.Task).Name)
			default:
				log.Error().Msgf("Unknown event type: %s", event.Type)
			}
		}
	}
}
