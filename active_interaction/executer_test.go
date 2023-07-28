package active_interaction_test

import (
	"testing"

	"github.com/apotema/go-active_interaction/active_interaction"
	. "github.com/apotema/go-active_interaction/active_interaction"
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
	assert.Equal(t, value, 1)
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
	A             int
	ValidateHooks `before:"SetA"`
}

func (s *SubjectBeforeValidate) SetA() {
	s.A = 4
}

func (s SubjectBeforeValidate) Run() int {
	return s.A
}

func TestBeforeValidateHook(t *testing.T) {
	value, _ := Execute[int](&SubjectBeforeValidate{A: 2})
	assert.Equal(t, 4, value)
}

type SubjectMultipleBeforeValidate struct {
	A             int `validate:"gte=4"`
	ValidateHooks `before:"SetA|SetB"`
}

func (s *SubjectMultipleBeforeValidate) SetA() {
	s.A += 4
}

func (s *SubjectMultipleBeforeValidate) SetB() {
	s.A += 4
}

func (s SubjectMultipleBeforeValidate) Run() int {
	return s.A
}

func TestMultipleBeforeValidateHooks(t *testing.T) {
	value, _ := Execute[int](&SubjectMultipleBeforeValidate{A: 2})
	assert.Equal(t, 10, value)
}

type SubjectAfterValidate struct {
	A             int
	ValidateHooks `after:"SetA"`
}

func (s *SubjectAfterValidate) SetA() {
	s.A += 4
}

func (s SubjectAfterValidate) Run() int {
	return s.A
}

func TestAfterValidateHooks(t *testing.T) {
	value, _ := Execute[int](&SubjectAfterValidate{A: 2})
	assert.Equal(t, 6, value)
}

type SubjectAfterValidateWithValidation struct {
	A             int `validate:"gte=4"`
	ValidateHooks `after:"SetA"`
}

func (s *SubjectAfterValidateWithValidation) SetA() {
	s.A += 4
}

func (s SubjectAfterValidateWithValidation) Run() int {
	return s.A
}

func AfterValidateHookExecuteAfterValidation(t *testing.T) {
	_, error := Execute[int](&SubjectAfterValidateWithValidation{A: 2})
	assert.Error(t, error)
}
