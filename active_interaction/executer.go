package active_interaction

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

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

func Execute[T any](interaction ActiveInteraction[T]) (*T, error) {
	x := interaction
	err := validate.Struct(interaction)
	if err != nil {
		return nil, err
	}

	CallMethod(&x, "SetA")

	// var i interface{} = interaction

	// t := reflect.TypeOf(i)
	// fmt.Println(t.Name())
	// fmt.Println("went here")
	// t.Elem().MethodByName("SetA")
	// reflect.ValueOf(t.Elem()).MethodByName("SetA").Call([]reflect.Value{})
	// for i := 0; i < t.NumField(); i++ {
	// 	f := t.Field(i)
	// 	if value, ok := f.Tag.Lookup("before"); ok {
	// 		fmt.Println("Tag found")
	// 		fmt.Println(value)
	// 		reflect.ValueOf(i).MethodByName("SetA").Call([]reflect.Value{})
	// 	} else {
	// 		fmt.Println("Tag not found")
	// 	}
	// }

	val := x.Run()
	return &val, nil
}
