package main

import (
	_ "casbinginxorm/docs"
	"casbinginxorm/middleware"
	"casbinginxorm/models"
	"casbinginxorm/routers"
	"casbinginxorm/service"
	"casbinginxorm/utils"
)

func init() {
	models.Setup()
	service.CasbinSetup()
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9000
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description                 Description for what is this security definition being used

func main() {

	//init log
	utils.InitLogger()

	r := routers.InitRouter()

	//文件记录到文件
	r.Use(middleware.ReqLog(), middleware.Recover())

	//Cors
	r.Use(middleware.Cors())

	r.Run(":9000") //参数为空 默认监听8080端口
}
