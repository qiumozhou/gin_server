package student

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/common"
	"strconv"
)


func addStudent(c *gin.Context){
	var stu Student
	c.BindJSON(&stu)
	res := AddStu(stu,c)
	common.POSTHandleResult(res,c)
}

func getStudent(c *gin.Context){
	id := c.Param("id")
	i,_ := strconv.Atoi(id)
	stu := GetStu(i)
	var newStudent Student
	res := GetCache("student")
	if res!= ""{
		json.Unmarshal([]byte(res),&newStudent)
		common.GETHandleResult(newStudent,c)
	}else{
		SetCache(stu)
		common.GETHandleResult(stu,c)
	}



}

func getStudentList(c *gin.Context){
	page,_ := strconv.Atoi(c.DefaultQuery("page","1"))
	limit,_ := strconv.Atoi(c.DefaultQuery("limit","10"))
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

func uploadIMG(c *gin.Context)  {
	file,_ := c.FormFile("imgfile")
	filePath := common.StringSplice("/home/qiumozhou/Desktop/gin_server/media/img/",file.Filename)
	fmt.Println("接收到文件: ", filePath)
	c.SaveUploadedFile(file, filePath)
	c.JSON(http.StatusOK, gin.H{
		"code":0,
		"msg":"ok",
	})
}



