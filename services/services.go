package services

import "todo-app/utils"

type Service interface {
	InitDependencies()
}

var AppServices = []Service{
	&TodosService{},
}

func InitServices() {
	for _, service := range AppServices {
		service.InitDependencies()
		utils.Register(service, utils.SERVICES)
	}
}
