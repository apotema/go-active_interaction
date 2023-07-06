package active_interaction

type ActiveInteraction[T any] interface {
	Run() T
}

type ValidateHooks struct {
	BeforeValidate func()
	Functions      []func()
}

func (BV *ValidateHooks) Execute() {

}
