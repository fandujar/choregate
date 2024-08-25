package controller

import (
	"context"
	"fmt"

	"github.com/fandujar/choregate/pkg/providers/tektoncd"
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
			// TODO: handle event to create task on choregate and add choregate labels to tektonCD task
			log.Info().Msg(fmt.Sprintf("event received: %v, event type: %t", event, event))
		}
	}
}
