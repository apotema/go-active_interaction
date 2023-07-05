package active_interaction

import (
	"github.com/go-playground/validator/v10"
)

type BeforeValidate interface {
	BeforeValidate() []func()
}

var validate *validator.Validate = validator.New()

func Execute[T any](interaction ActiveInteraction[T]) (*T, error) {
	err := validate.Struct(interaction)
	if err != nil {
		return nil, err
	}

	if bf, ok := interaction.(BeforeValidate); ok {
		for _, beforeValidate := range bf.BeforeValidate() {
			beforeValidate()
		}
	}
	val := interaction.Run()
	return &val, nil
}
