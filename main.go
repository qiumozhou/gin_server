package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/student"
	"server/app/teacher"
	"server/routers"
)

func index(r *gin.Engine){
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
}

func main()  {
	// 加载多个APP的路由配置
	routers.Include(
		student.Routers,
		teacher.Routers,
		index,
		)
	// 初始化路由
	r := routers.Init()
	fmt.Println("http://127.0.0.1:8000")
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

}