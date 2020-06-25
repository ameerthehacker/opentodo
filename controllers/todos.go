package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"opentodo/models"
)

func TodosController(router *gin.Engine, connection *gorm.DB) {
	todosRouter := router.Group("/todos")
	var todos []models.Todo

	connection.Find(&todos)

	todosRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, todos)
	})
}
