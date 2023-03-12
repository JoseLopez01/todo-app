package main

import (
	"github.com/gin-gonic/gin"

	"todo-app/controllers"
)

func main() {
	engine := gin.Default()
	apiGroup := engine.Group("/api")

	controllers.InitControllers(apiGroup)

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
