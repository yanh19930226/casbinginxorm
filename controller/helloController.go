package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct{}

// Hello godoc
//
// @Summary Hello
// @Tags hello
// @Description Hello
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Router /api/v1/hello [get]
func (r *HelloController) Hello(c *gin.Context) {
	fmt.Println("Hello 接收到GET请求..")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": "Hello 接收到GET请求..",
	})
}
