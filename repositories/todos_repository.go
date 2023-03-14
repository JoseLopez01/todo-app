package repositories

import "todo-app/models"

var todos = []models.Todo{
	{Name: "First todo", Description: "This is a first todo"},
	{Name: "Second todo", Description: "This is a second todo", Completed: true},
}

type TodosRepository struct {
}

func (repository *TodosRepository) GetAll() []models.Todo {
	return todos
}
