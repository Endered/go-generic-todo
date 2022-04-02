package server

import (
	"net/http"
	"todo/src/controller"
	"todo/src/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Router(db *sqlx.DB) *gin.Engine {

	tr := repository.NewTodoRepositoryImpl(db)
	tc := controller.NewTodoControllerImpl(tr)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/todo", tc.GetAllTodo)
	r.GET("/todo/:id", tc.GetTodo)
	r.POST("/todo", tc.AddTodo)
	r.DELETE("/todo/:id", tc.DeleteTodo)
	return r
}
