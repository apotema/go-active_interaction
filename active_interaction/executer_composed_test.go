package active_interaction_test

import (
	"fmt"
	"testing"

	. "github.com/apotema/go-active_interaction/active_interaction"
	"github.com/stretchr/testify/assert"
)

type SubjectComposedInteraction struct {
	InteractionUtils
}

func (s SubjectComposedInteraction) Run() int {
	val, _ := Compose[int, int](&s, &AnotherInteraction{})
	return val
}

type AnotherInteraction struct {
	InteractionUtils
}

func (s AnotherInteraction) Run() int {
	return 4
}

func TestComposedExecute(t *testing.T) {
	value, _ := Execute[int](&SubjectComposedInteraction{})
	assert.Equal(t, 4, value)
}

type SubjectComposedInteractionMain struct {
	InteractionUtils
}

type AnotherInteractionWithError struct {
	A int `validate:"gte=4"`
	InteractionUtils
}

func (s *SubjectComposedInteractionMain) Run() int {
	val, ee := Compose[int, int](s, &AnotherInteractionWithError{})
	fmt.Println("error inside compose test", ee)
	return val
}

func (s AnotherInteractionWithError) Run() int {
	return 4
}

func TestComposedExecuteWithError(t *testing.T) {
	_, error := Execute[int](&SubjectComposedInteractionMain{})
	assert.NotNil(t, error, "Error is not Nil")
}

func TestComposedExecuteWithErrorReturnstThSubInteractionError(t *testing.T) {
	_, error := Execute[int](&SubjectComposedInteractionMain{})
	assert.Equal(t, error.Error(), "\"AnotherInteractionWithError.A\": [\"Field validation for 'A' failed on the 'gte' tag\"]")
}
