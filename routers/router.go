package routers

import (
	"casbinginxorm/controller"
	"casbinginxorm/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {

	//获取router路由对象
	r := gin.New()
	//文件记录到文件
	// r.Use(utils.LoggerToFile())
	//注意 Recover 要尽量放在第一个被加载
	//如不是的话，在recover前的中间件或路由，将不能被拦截到
	//程序的原理是：
	//1.请求进来，执行recover
	//2.程序异常，抛出panic
	//3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
	r.Use(middleware.Recover())

	//Cors
	r.Use(middleware.Cors())

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//Group
	apiv1 := r.Group("/api/v1")

	//No verification required
	authController := new(controller.AuthController)
	apiv1.POST("/auth/createtoken", authController.CreateToken)
	apiv1.GET("/auth/validtoken/{token}", authController.ValidToken)
	apiv1.POST("/auth/createrefreshtoken", authController.CreateRefreshToken)
	apiv1.GET("/auth/validrefreshtoken/{refreshtoken}", authController.ValidRefreshToken)

	apiv1.GET("/pingerror", func(c *gin.Context) {
		// 无意抛出 panic
		var slice = []int{1, 2, 3, 4, 5}
		slice[6] = 6
	})

	//JwtAuthorize
	apiv1.Use(middleware.JwtAuthorize())
	{
		helloController := new(controller.HelloController)
		apiv1.GET("/hello", helloController.Hello)

		userController := new(controller.UserController)
		apiv1.GET("/users", userController.ListUsers)
		apiv1.POST("/users", userController.AddUser)
		apiv1.PATCH("/users/{id}", userController.UpdateUser)
		apiv1.DELETE("/users/{id}", userController.DeleteUser)

		roleController := new(controller.RoleController)
		apiv1.GET("/roles", roleController.ListRoles)
		apiv1.POST("/roles", roleController.AddRole)
		apiv1.PATCH("/roles/{id}", roleController.UpdateRole)
		apiv1.DELETE("/roles/{id}", roleController.DeleteRole)
	}

	return r
}
