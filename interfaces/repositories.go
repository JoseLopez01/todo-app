package interfaces

import "todo-app/models"

type TodosRepository interface {
	GetAll() []models.Todo
}
