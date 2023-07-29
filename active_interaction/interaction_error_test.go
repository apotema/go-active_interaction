package active_interaction_test

import (
	"testing"

	"github.com/apotema/go-active_interaction/active_interaction"
	"github.com/stretchr/testify/assert"
)

func TestAddError(t *testing.T) {
	interaction_error := active_interaction.InteractionError{}
	interaction_error.AddError("field", "message")
	assert.Equal(t, interaction_error.Error(), "\"field\": [\"message\"]")
}

func TestAddErrorMultipleMessages(t *testing.T) {
	interaction_error := active_interaction.InteractionError{}
	interaction_error.AddError("field", "message")
	interaction_error.AddError("field", "another message")
	assert.Equal(t,
		interaction_error.Error(),
		"\"field\": [\"message\",\"another message\"]",
	)
}

func TestAddErrorMultipleFields(t *testing.T) {
	interaction_error := active_interaction.InteractionError{}
	interaction_error.AddError("field", "message")
	interaction_error.AddError("another field", "another message")
	assert.Contains(t,
		interaction_error.Error(),
		"\"field\": [\"message\"]",
	)
	assert.Contains(t,
		interaction_error.Error(),
		"\"another field\": [\"another message\"]",
	)
}
