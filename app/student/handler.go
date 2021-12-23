package student

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/app/common"
	"strconv"
)


func addStudent(c *gin.Context){
	var stu Student
	c.BindJSON(&stu)
	res := AddStu(stu)
	common.POSTHandleResult(res,c)
}

func getStudent(c *gin.Context){
	id := c.Param("id")
	i,_ := strconv.Atoi(id)
	stu := GetStu(i)
	common.GETHandleResult(stu,c)
}

func getStudentList(c *gin.Context){
	page,_ := strconv.Atoi(c.Query("page"))
	limit,_ := strconv.Atoi(c.Query("limit"))
	stuList := GetStuList(page,limit)
	common.GETHandleResult(stuList,c)
}


func putStudent(c *gin.Context){
	var stu Student
	c.BindJSON(&stu)
	fmt.Println(stu)
	res := UpdateStu(stu)
	common.PUTHandleResult(res,c)
}

func deleteStudent(c *gin.Context){
	id := c.Param("id")
	i,_ := strconv.Atoi(id)
	stu := DeleteStu(i)
	common.DELETEHandleResult(stu,c)
}



