package active_interaction

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func CallMethod(i interface{}, methodName string) interface{} {
	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(i)

	// if we start with a pointer, we need to get value pointed to
	// if we start with a value, we need to get a pointer to that value
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(i))
		temp := ptr.Elem()
		temp.Set(value)
	}

	// check for method on value
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	// check for method on pointer
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	if finalMethod.IsValid() {
		return finalMethod.Call([]reflect.Value{})
	}

	// return or panic, method not found of either type
	return ""
}

func Compose[T any](interaction ActiveInteraction[T]) (T, InteractionError) {
	result, error := Execute[T](interaction)
	return result, error
}

func Execute[T any](interaction ActiveInteraction[T]) (T, InteractionError) {
	var validate *validator.Validate = validator.New()
	var ptr reflect.Value
	var value reflect.Value

	ptr = reflect.New(reflect.TypeOf(interaction))
	temp := ptr.Elem()
	value = reflect.ValueOf(interaction)
	temp.Set(value)

	var t reflect.Type
	if reflect.TypeOf(interaction).Kind() == reflect.Ptr {
		t = reflect.TypeOf(interaction).Elem()
	} else {
		t = reflect.TypeOf(interaction)
	}

	f, _ := t.FieldByName("ValidateHooks")
	if value, ok := f.Tag.Lookup("before"); ok {
		for _, s := range strings.Split(value, "|") {
			CallMethod(interaction, s)
		}
	}

	err := validate.Struct(interaction)
	if err != nil {
		var result T
		interactionError := interaction.AddValidatorError(err)
		return result, interactionError
	}

	if value, ok := f.Tag.Lookup("after"); ok {
		for _, s := range strings.Split(value, "|") {
			CallMethod(interaction, s)
		}
	}

	beforeExecuteHook, _ := t.FieldByName("ExecuteHooks")
	if value, ok := beforeExecuteHook.Tag.Lookup("before"); ok {
		for _, s := range strings.Split(value, "|") {
			CallMethod(interaction, s)
		}
	}

	val := interaction.Run()

	afterExecuteHook, _ := t.FieldByName("ExecuteHooks")
	if value, ok := afterExecuteHook.Tag.Lookup("after"); ok {
		for _, s := range strings.Split(value, "|") {
			CallMethod(interaction, s)
		}
	}

	return val, interaction.GetError()
}
