package services

type TodosService struct {
	Test string
}

func (service *TodosService) GetAll() (string, error) {
	return "TodoService:GetAll", nil
}
