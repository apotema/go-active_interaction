package active_interaction

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func Execute[T any](interaction ActiveInteraction[T]) (*T, error) {
	err := validate.Struct(interaction)
	if err != nil {
		return nil, err
	}

	var i interface{} = interaction

	if _, ok := i.(BeforeValidate); ok {
		// fmt.Println("before validate")
		values := reflect.ValueOf(interaction).MethodByName("BeforeValidate").Call(nil)
		// fmt.Println("before validate")
		for _, beforeValidate := range values {
			beforeValidate()
		}
	} else {
		fmt.Println("no before validate")
	}

	val := interaction.Run()
	return &val, nil
}
