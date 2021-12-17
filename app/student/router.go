package student

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine){
	v1 := e.Group("/api/v1")
	{
		v1.GET("/student",indexHandler)
	}

}