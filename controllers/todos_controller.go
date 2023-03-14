package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/interfaces"
	"todo-app/utils"
)

type TodosController struct {
	TodoService interfaces.TodosService
}

func (controller *TodosController) InitRoutes(router *gin.RouterGroup) {
	todoGroup := router.Group("/todos")
	todoGroup.GET("", controller.GetAll)
}

func (controller *TodosController) InitDependencies() {
	utils.InitDependencies(controller)
}

func (controller *TodosController) GetAll(ctx *gin.Context) {
	message, _ := controller.TodoService.GetAll()
	ctx.JSON(http.StatusOK, gin.H{"data": message})
}
