package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"opentodo/models"
)

func TodosController(router *gin.Engine) {
	todosRouter := router.Group("/todos")
	todos := []models.Todo{
		{
			Title: "Buy Eggs",
			Done:  false,
		},
	}

	todosRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, todos)
	})
}
