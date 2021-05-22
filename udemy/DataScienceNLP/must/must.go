package must

import (
	"log"
	"reflect"
)

func ReturnElseLogFatal(params ...interface{}) interface{} {
	f := reflect.ValueOf(params[0])

	inputs := []reflect.Value{}
	for _, param := range params[1:] {
		inputs = append(inputs, reflect.ValueOf(param))
	}
	e := f.Call(inputs)

	if e[1].Interface() != nil {
		err := e[1].Interface().(error)
		log.Fatal(err)
	}

	return e[0].Interface()

}
