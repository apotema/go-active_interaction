package active_interaction

type ActiveInteraction[T any] interface {
	Run() T
	AddError(field string, value string)
	AddValidatorError(err error) InteractionError
	GetError() InteractionError
}

type InteractionUtils struct {
	Error InteractionError
}

func (m *InteractionUtils) AddValidatorError(err error) InteractionError {
	return m.Error.AddValidatorError(err)
}

func (m *InteractionUtils) AddError(field string, value string) {
	m.Error.AddError(field, value)
}

func (m InteractionUtils) GetError() InteractionError {
	return m.Error
}

type ValidateHooks struct {
}

type ExecuteHooks struct {
}
