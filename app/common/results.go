package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func POSTHandleResult(res bool,c *gin.Context){
	if res != true{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":"fail",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":"success",
		})
	}
}

func GETHandleResult(res interface{},c *gin.Context){
	data := map[string]interface{}{
		"code":0,
		"msg":"ok",
		"data":res,
	}
	c.JSON(http.StatusOK,data)

}


func PUTHandleResult(res bool,c *gin.Context){
	if res != true{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":"fail",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":"success",
		})
	}
}

func DELETEHandleResult(res bool,c *gin.Context){
	if res != true{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":"fail",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":"success",
		})
	}
}



