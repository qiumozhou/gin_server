package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/student"
	"server/app/teacher"
	"server/routers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func index(r *gin.Engine){
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
}

func main()  {
	//初始化数据库链接
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local")

	if err!= nil{
		panic(err)
	}
	db.AutoMigrate(&student.Student{},&teacher.Teacher{})
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	defer db.Close()
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