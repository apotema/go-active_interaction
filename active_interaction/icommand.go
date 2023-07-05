package active_interaction

type ActiveInteraction[T any] interface {
	Run() T
}

type BeforeValidate interface {
	BeforeValidate() []func()
}
