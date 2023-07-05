package active_interaction_test

import (
	"fmt"
	"testing"

	"github.com/apotema/go-active_interaction/active_interaction"
	"github.com/stretchr/testify/assert"
)

type Subject struct {
	A int
}

func (s Subject) Run() int {
	return 1
}

func TestReturnValue(t *testing.T) {
	value, _ := active_interaction.Execute[int](Subject{A: 2})
	assert.Equal(t, *value, 1)
}

type SubjectValidate struct {
	A int `validate:"gte=4"`
}

func (s SubjectValidate) Run() int {
	return 1
}

func Test_Validate_invalid(t *testing.T) {
	_, error := active_interaction.Execute[int](SubjectValidate{A: 2})
	assert.Error(t, error)
}

type SubjectBeforeValidate struct {
	A int
}

func (s SubjectBeforeValidate) BeforeValidate() []func() {
	fmt.Println("called by name")
	var setA = func() {
		s.A = 4
	}
	return []func(){setA}
}

func (s SubjectBeforeValidate) Run() int {
	return s.A
}

func TestBeforeBeforeValidate(t *testing.T) {
	value, _ := active_interaction.Execute[int](SubjectBeforeValidate{A: 2})
	assert.Equal(t, 4, *value)
}
