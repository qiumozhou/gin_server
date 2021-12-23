package student

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine){
	v1 := e.Group("/api/v1")
	{
		v1.GET("/student/:id",getStudent)
		v1.POST("/student",addStudent)
		v1.GET("/students",getStudentList)
		v1.PUT("/student",putStudent)
		v1.DELETE("/student/:id",deleteStudent)
	}

}