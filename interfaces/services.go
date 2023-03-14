package interfaces

type TodosService interface {
	GetAll() (string, error)
}
