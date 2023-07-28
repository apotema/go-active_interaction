package active_interaction_test

import (
	"testing"

	. "github.com/apotema/go-active_interaction/active_interaction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Subject struct {
	A int
}

func (s Subject) Run() int {
	return 1
}

func TestReturnValue(t *testing.T) {
	value, _ := Execute[int](Subject{A: 2})
	assert.Equal(t, value, 1)
}

type SubjectValidate struct {
	A int `validate:"gte=4"`
}

func (s SubjectValidate) Run() int {
	return 1
}

func Test_Validate_invalid(t *testing.T) {
	_, error := Execute[int](SubjectValidate{A: 2})
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

func TestAfterValidateHookExecuteAfterValidation(t *testing.T) {
	_, error := Execute[int](&SubjectAfterValidateWithValidation{A: 2})
	assert.Error(t, error)
}

type SubjectBeforeExecute struct {
	A            int
	ExecuteHooks `before:"SetA"`
}

func (s *SubjectBeforeExecute) SetA() {
	s.A += 4
}

func (s SubjectBeforeExecute) Run() int {
	return s.A
}

func TestBeforeExecuteHook(t *testing.T) {
	value, _ := Execute[int](&SubjectBeforeExecute{A: 2})
	assert.Equal(t, 6, value)
}

type SubjectMultipleBeforeExecutes struct {
	A            int
	ExecuteHooks `before:"SetA|SetB"`
}

func (s *SubjectMultipleBeforeExecutes) SetA() {
	s.A += 4
}

func (s *SubjectMultipleBeforeExecutes) SetB() {
	s.A += 4
}

func (s SubjectMultipleBeforeExecutes) Run() int {
	return s.A
}

func TestMultipleBeforeExecuteHook(t *testing.T) {
	value, _ := Execute[int](&SubjectMultipleBeforeExecutes{A: 2})
	assert.Equal(t, 10, value)
}

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

type SubjectAfterExecute struct {
	A            int
	ExecuteHooks `after:"SetA"`
	mock         *MyMockedObject
}

func (s *SubjectAfterExecute) SetA() {
	s.mock.DoSomething(10)
	s.A += 4
}

func (s *SubjectAfterExecute) Run() int {
	return s.A
}

func TestAfterExecuteHookIsCalled(t *testing.T) {
	testObj := new(MyMockedObject)
	testObj.On("DoSomething", mock.Anything).Return(true, nil)

	value, _ := Execute[int](&SubjectAfterExecute{A: 2, mock: testObj})
	testObj.AssertExpectations(t)
	assert.Equal(t, 2, value)
}
