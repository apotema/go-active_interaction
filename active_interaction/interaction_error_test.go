package active_interaction_test

import (
	"errors"
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

func TestAddValidatorError(t *testing.T) {
	interaction_error := active_interaction.InteractionError{}
	errorMessage := `Key: 'SubjectValidate.A' Error:Field validation for 'A' failed on the 'gte' tag
Key: 'SubjectValidate.B' Error:Field validation for 'B' failed on the 'gte' tag`
	err := errors.New(errorMessage)
	interaction_error.AddValidatorError(err)
	assert.Equal(t, interaction_error.Error(), "\"A\": [\"Field validation for 'A' failed on the 'gte' tag\"], \"B\": [\"Field validation for 'B' failed on the 'gte' tag\"]")
}
