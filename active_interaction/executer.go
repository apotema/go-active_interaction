package active_interaction

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func Execute[T any](interaction ActiveInteraction[T]) (*T, error) {
	err := validate.Struct(interaction)
	if err != nil {
		return nil, err
	}

	var i interface{} = interaction

	// if _, ok := i.(BeforeValidate); ok {
	// 	// fmt.Println("before validate")
	// 	values := reflect.ValueOf(interaction).MethodByName("BeforeValidate").Call(nil)
	// 	// fmt.Println("before validate")
	// 	a := values[0].Interface().([]func())
	// 	a[0]()
	// } else {
	// 	fmt.Println("no before validate")
	// }

	val := i.(ActiveInteraction[T]).Run()
	return &val, nil
}
