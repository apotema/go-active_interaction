package active_interaction

type ActiveInteraction[T any] interface {
	Run() T
	AddError(error InteractionError)
}

type InteractionUtils struct {
	Errors []InteractionError
}

func (m *InteractionUtils) AddError(error InteractionError) {
	m.Errors = append(m.Errors, error)
}

type ValidateHooks struct {
}

type ExecuteHooks struct {
}
