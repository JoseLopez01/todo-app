package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/interfaces"
)

type TodoController struct {
	TodoService interfaces.TodoService
}

func (controller TodoController) InitRoutes(router *gin.RouterGroup) {
	todoGroup := router.Group("/todos")
	todoGroup.GET("", controller.GetAll)
}

func (controller TodoController) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "list_of_todos"})
}
