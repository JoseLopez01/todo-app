package main

import (
	"github.com/gin-gonic/gin"
	"todo-app/controllers"
	"todo-app/repositories"
	"todo-app/services"
)

func main() {
	engine := gin.Default()
	apiGroup := engine.Group("/api")

	repositories.InitRepositories()
	services.InitServices()
	controllers.InitControllers(apiGroup)

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
