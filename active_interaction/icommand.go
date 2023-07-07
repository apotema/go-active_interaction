package active_interaction

type ActiveInteraction[T any] interface {
	Run() T
}

type ValidateHooks struct {
}
