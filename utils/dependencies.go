package utils

import (
	"fmt"
	"reflect"
)

var messageError = "there's no dependency with token %s"

func InitDependencies(object interface{}) {
	element := reflect.ValueOf(object).Elem()
	for i := 0; i < element.NumField(); i++ {
		field := element.Field(i)
		if field.CanSet() {
			token := field.Type().Name()

			var dependency interface{}
			dependency, err := Get(token, SERVICES)
			if err != nil {
				dependency, err = Get(token, REPOSITORIES)
				if err != nil {
					panic(fmt.Errorf(messageError, token))
				}
			}
			field.Set(reflect.ValueOf(dependency))
		}
	}
}
