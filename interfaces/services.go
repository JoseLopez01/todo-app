package interfaces

import "todo-app/models"

type TodosService interface {
	GetAll() []models.Todo
}
