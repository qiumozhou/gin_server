package teacher

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"welcome teacher",
	})
}