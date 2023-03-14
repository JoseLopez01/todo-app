package repositories

import "todo-app/utils"

type Repository interface {
}

var appRepositories = []Repository{
	&TodosRepository{},
}

func InitRepositories() {
	for _, repository := range appRepositories {
		utils.Register(repository, utils.REPOSITORIES)
	}
}
