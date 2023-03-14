package services

import (
	"todo-app/interfaces"
	"todo-app/models"
	"todo-app/utils"
)

type TodosService struct {
	Repository interfaces.TodosRepository
}

func (service *TodosService) InitDependencies() {
	utils.InitDependencies(service)
}

func (service *TodosService) GetAll() []models.Todo {
	return service.Repository.GetAll()
}
