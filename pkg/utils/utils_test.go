package utils_test

import (
	"testing"

	"github.com/fandujar/choregate/pkg/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestGenerateID() tests the GenerateID() function.
func TestGenerateID(t *testing.T) {
	id, err := utils.GenerateID()

	assert.NoError(t, err)
	assert.NotEqual(t, 0, id)
	assert.NotPanics(t, func() { uuid.MustParse(id.String()) })

}
