package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Validate() gin.HandlerFunc{
	return func(c *gin.Context){
		token := c.DefaultQuery("token","")
		if token != ""{
			//c.JSON(http.StatusOK,gin.H{"message":"身份验证成功"})
			c.Next()
		}else{
			c.JSON(http.StatusOK,gin.H{"message":"请登录"})
			c.Abort()
		}
	}
}