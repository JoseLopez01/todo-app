package main

import (
	"github.com/gin-gonic/gin"
	"todo-app/services"

	"todo-app/controllers"
)

func main() {
	engine := gin.Default()
	apiGroup := engine.Group("/api")

	services.InitServices()
	controllers.InitControllers(apiGroup)

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
