package utils

import (
	"fmt"
	"reflect"
)

const (
	REPOSITORIES = "REPOSITORIES"
	SERVICES     = "SERVICES"
)

type Registered map[string]interface{}

var appServices = make(Registered)
var appRepositories = make(Registered)

func Register(object interface{}, registerType string) {
	name := reflect.TypeOf(object).Name()
	appServices[name] = object
}

func Get(token string, registerType string) (interface{}, error) {
	var source Registered
	if registerType == REPOSITORIES {
		source = appRepositories
	} else if registerType == SERVICES {
		source = appServices
	}

	object, ok := source[token]
	if !ok {
		return nil, fmt.Errorf("no register with token %s", token)
	}

	return object, nil
}
