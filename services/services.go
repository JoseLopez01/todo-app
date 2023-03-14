package services

import "todo-app/utils"

type Service interface {
}

var AppServices = []Service{
	&TodosService{
		Test: "test_test",
	},
}

func InitServices() {
	for _, service := range AppServices {
		utils.Register(service, utils.SERVICES)
	}
}
