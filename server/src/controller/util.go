package controller

import (
	"todo/src/util"

	"github.com/gin-gonic/gin"
)

func ShouldBindJSON[T any](c *gin.Context) util.Result[T] {
	var res T
	err := c.ShouldBindJSON(&res)
	return util.NewResult(res, err)
}
