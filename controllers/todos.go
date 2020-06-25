package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"opentodo/models"
)

func TodosController(router *gin.Engine, connection *gorm.DB) {
	todosRouter := router.Group("/todos")
	var todoInput struct {
		Title *string `json:"title" binding:"required"`
		Done  *bool   `json:"done" binding:"required"`
	}

	todosRouter.GET("/", func(ctx *gin.Context) {
		var todos []models.Todo

		connection.Find(&todos)

		ctx.JSON(http.StatusOK, todos)
	})

	todosRouter.POST("/", func(ctx *gin.Context) {
		err := ctx.BindJSON(&todoInput)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
		} else {
			todo := models.Todo{
				Title: *todoInput.Title,
				Done:  *todoInput.Done,
			}

			connection.Create(&todo)

			ctx.Status(http.StatusCreated)
		}
	})

	todosRouter.DELETE("/:id", func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")

		connection.Delete(models.Todo{}, "id = ?", id)

		ctx.Status(http.StatusOK)
	})

	todosRouter.PUT("/:id", func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		err := ctx.BindJSON(&todoInput)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
		} else {
			connection.Model(new(models.Todo)).Where("id = ?", id).Updates(models.Todo{
				Title: *todoInput.Title,
				Done:  *todoInput.Done,
			})

			ctx.Status(http.StatusOK)
		}
	})
}
