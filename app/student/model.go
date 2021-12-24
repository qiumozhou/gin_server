package student

import (
	"server/component/mysql"
	"time"
)

type Student struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"type:varchar(20);not null"`
	Age        int `json:"age" gorm:"type:int;not null;"`
	Gender     int `json:"gender" gorm:"type:int;not null;"`
	TeacherID      int `json:"teacher_id" gorm:"type:int;not null;"`
	CreatedAt time.Time `json:"ctm"`
	UpdatedAt time.Time `json:"utm"`
}


func (stu Student) TableName() string{
	return "tb_student"
}

func (Student)GetStuByName(name string)(stu Student){
	mysql.GetDb("gin").Where("name = ?", name).First(&stu)
	return
}

func AddStu(stu Student)bool{
	mysql.GetDb("gin").Create(&stu)
	return true
}

func GetStu(id int)(stu Student){
	mysql.GetDb("gin").Where("id = ?", id).First(&stu)
	return
}

func GetStuList(page int,limit int)(stulist []Student){
	mysql.GetDb("gin").Limit(limit).Offset(page*limit-limit).Find(&stulist)
	return
}

func UpdateStu(stu Student)bool{
	mysql.GetDb("gin").Model(&Student{}).Updates(stu)
	return true
}

func DeleteStu(id int)bool{
	mysql.GetDb("gin").Where("id=?",id).Delete(&Student{})
	return true
}

