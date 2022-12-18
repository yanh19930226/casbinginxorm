package service

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Enforcer *casbin.Enforcer

// 初始化casbin
func CasbinSetup() {
	a, err := xormadapter.NewAdapter("mysql", "root:66^^66@tcp(121.43.34.54:3306)/rbac_db?charset=utf8", true)
	if err != nil {
		log.Printf("连接数据库错误: %v", err)
		return
	}
	e, err := casbin.NewEnforcer("conf/rbac_models.conf", a)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
		return
	}

	//获取router路由对象
	r := gin.New()
	//增加policy
	r.POST("/api/v1/add", func(c *gin.Context) {
		fmt.Println("增加Policy")
		if ok, _ := e.AddPolicy("admin", "/api/v1/world", "GET"); !ok {
			fmt.Println("Policy已经存在")
		} else {
			fmt.Println("增加成功")
		}
	})
	//删除policy
	r.DELETE("/api/v1/delete", func(c *gin.Context) {
		fmt.Println("删除Policy")
		if ok, _ := e.RemovePolicy("admin", "/api/v1/world", "GET"); !ok {
			fmt.Println("Policy不存在")
		} else {
			fmt.Println("删除成功")
		}
	})
	//获取policy
	r.GET("/api/v1/get", func(c *gin.Context) {
		fmt.Println("查看policy")
		list := e.GetPolicy()
		for _, vlist := range list {
			for _, v := range vlist {
				fmt.Printf("value: %s, ", v)
			}
		}
	})

	Enforcer = e
}
