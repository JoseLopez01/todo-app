package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	InitRoutes(router *gin.RouterGroup)
	InitDependencies()
}

var AppControllers = []Controller{
	&TodosController{},
}

func InitControllers(group *gin.RouterGroup) {
	for _, controller := range AppControllers {
		controller.InitRoutes(group)
		controller.InitDependencies()
	}
}
