package middleware

import (
	"casbinginxorm/utils"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//打印错误堆栈信息
				utils.GetLogger().Error(r)
				debug.PrintStack()
				c.JSON(http.StatusOK, utils.Fail(utils.ERROR))
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()
			}
		}()

		//加载完 defer recover，继续后续接口调用
		c.Next()
	}
}
