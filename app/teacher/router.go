package teacher

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine){
	v1 := e.Group("/api/v1")
	{
		v1.GET("/teacher/:id",getTeacher)
		v1.POST("/teacher",addTeacher)
		v1.PUT("/teacher",putTeacher)
		v1.DELETE("/teacher/:id",deleteTeacher)
	}
}