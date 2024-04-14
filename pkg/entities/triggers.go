package entities

import (
	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
)

// TriggerConfig represents the configuration of a trigger.
type TriggerConfig struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
}

// Trigger represents a trigger.
type Trigger struct {
	*TriggerConfig
}

// TriggerCategory represents a trigger category.
type TriggerCategory string

const (
	// TriggerCategorySchedule represents a schedule trigger category.
	TriggerCategorySchedule TriggerCategory = "schedule"
	// TriggerCategoryWebhook represents a webhook trigger category.
	TriggerCategoryWebhook TriggerCategory = "webhook"
)

// NewTrigger creates a new Trigger with the given configuration and default values.
func NewTrigger(config *TriggerConfig) (*Trigger, error) {
	var err error

	if config.ID == uuid.Nil {
		config.ID, err = utils.GenerateID()
		if err != nil {
			return nil, err
		}
	}

	return &Trigger{
		TriggerConfig: config,
	}, nil
}
