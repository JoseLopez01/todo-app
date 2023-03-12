package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	InitRoutes(router *gin.RouterGroup)
}

var AppControllers = []Controller{
	TodoController{},
}

func InitControllers(group *gin.RouterGroup) {
	for _, controller := range AppControllers {
		controller.InitRoutes(group)
	}
}
