package controller

import (
	"net/http"
	"strconv"
	"todo/src/domain"
	"todo/src/repository"
	"todo/src/util"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	GetAllTodo(*gin.Context)
	GetTodo(*gin.Context)
	AddTodo(*gin.Context)
	DeleteTodo(*gin.Context)
}

type TodoControllerImpl struct {
	tr repository.TodoRepository
}

func NewTodoControllerImpl(tr repository.TodoRepository) *TodoControllerImpl {
	return &TodoControllerImpl{
		tr: tr,
	}
}

func (con *TodoControllerImpl) GetAllTodo(c *gin.Context) {
	todos := con.tr.GetAllTodo()
	if todos.Success {
		c.JSON(http.StatusOK, todos.Value)
	} else {
		c.String(500, todos.Error.Error())
	}
}

func (con *TodoControllerImpl) GetTodo(c *gin.Context) {
	idstr := c.Param("id")
	id := util.NewResult(strconv.Atoi(idstr))
	todo := util.FlatMapResult(id, func(id int) util.Result[domain.Todo] {
		return con.tr.GetTodo(id)
	})
	if todo.Success {
		c.JSON(http.StatusOK, todo.Value)
	} else {
		c.String(500, todo.Error.Error())
	}

}

func (con *TodoControllerImpl) AddTodo(c *gin.Context) {
	content := ShouldBindJSON[TodoContent](c)
	todo := util.FlatMapResult(content, func(content TodoContent) util.Result[domain.Todo] {
		return con.tr.AddTodo(content.Content)
	})
	if todo.Success {
		c.JSON(http.StatusOK, todo.Value)
	} else {
		c.String(500, todo.Error.Error())
	}
}

func (con *TodoControllerImpl) DeleteTodo(c *gin.Context) {
	idstr := c.Param("id")
	rid := util.NewResult(strconv.Atoi(idstr))
	if !rid.Success {
		c.String(400, rid.Error.Error())
		return
	}
	id := rid.Value
	err := con.tr.DeleteTodo(id)
	if err == nil {
		c.String(http.StatusOK, "Successful deleted")
	} else {
		c.String(500, err.Error())
	}
}
