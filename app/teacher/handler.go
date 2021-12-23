package teacher

import (
	"github.com/gin-gonic/gin"
	"server/app/common"
	"strconv"
)

func getTeacher(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	teacher:=GetTeacher(id)
	common.GETHandleResult(teacher,c)
}

func addTeacher(c *gin.Context){
	var teacher Teacher
	c.BindJSON(&teacher)
	res := AddTeacher(teacher)
	common.POSTHandleResult(res,c)
}

func putTeacher(c *gin.Context){
	var teacher Teacher
	c.BindJSON(&teacher)
	res := PutTeacher(teacher)
	common.PUTHandleResult(res,c)
}

func deleteTeacher(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	res := DeleteTeacher(id)
	common.DELETEHandleResult(res,c)
}