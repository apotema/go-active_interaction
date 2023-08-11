package active_interaction_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/apotema/go-active_interaction"
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
	interactionError := &active_interaction.InteractionError{}
	assert.False(t, interactionError.HasError(), "HasError() should return false when there are no errors")

	// Test AddError and HasError together
	field := "name"
	message := "Name is required"

	interactionError.AddError(field, message)
	assert.True(t, interactionError.HasError(), "HasError() should return true after adding an error")

	// Check if the error was added correctly
	errorMessages, ok := interactionError.ErrorMap()[field]
	assert.True(t, ok, "Error should be added to the correct field")
	assert.Contains(t, errorMessages, message, "Error message should be added to the correct field")
}

func TestAppend(t *testing.T) {
	interactionErrorToAppend := &active_interaction.InteractionError{}
	field := "appendedError"
	message := "This is an appended error"
	interactionErrorToAppend.AddError(field, message)

	initialInteractionError := &active_interaction.InteractionError{}
	field = "initialError"
	message = "This is an initial error"
	initialInteractionError.AddError(field, message)

	initialInteractionError.Append(interactionErrorToAppend, "testPrefix")

	expectedErrors := map[string][]string{
		"initialError":             {"This is an initial error"},
		"testPrefix.appendedError": {"This is an appended error"},
	}

	actualErrors := initialInteractionError.ErrorMap()
	assert.Equal(t, expectedErrors, actualErrors)
}

func TestAddValidatorError(t *testing.T) {
	interaction_error := active_interaction.InteractionError{}
	errorMessage := `Key: 'SubjectValidate.A' Error:Field validation for 'A' failed on the 'gte' tag
Key: 'SubjectValidate.B' Error:Field validation for 'B' failed on the 'gte' tag`
	err := errors.New(errorMessage)
	interaction_error.AddValidatorError(err)
	assert.Equal(t, interaction_error.Error(), "\"A\": [\"Field validation for 'A' failed on the 'gte' tag\"], \"B\": [\"Field validation for 'B' failed on the 'gte' tag\"]")
}

func TestAddValidatorErrorHash(t *testing.T) {
	errStr := "'object.field1' Error:This is field1 error\n" +
		"'object.field2' Error:This is field2 error"
	err := fmt.Errorf(errStr)

	m := active_interaction.InteractionError{}
	m = m.AddValidatorError(err)

	expectedErrors := map[string][]string{
		"field1": {"This is field1 error"},
		"field2": {"This is field2 error"},
	}

	actualErrors := m.ErrorMap()
	assert.Equal(t, expectedErrors, actualErrors)
}
