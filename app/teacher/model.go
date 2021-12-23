package teacher

import (
	"server/component/mysql"
	"time"
)

type Teacher struct {
	ID int `gorm:"primary_key"`
	Name  string `gorm:"type:varchar(20);not null"`
	Age   int   `gorm:"type:int;not null"`
	Gender int   `gorm:"type:int;not null"`
	Class  string `gorm:"type:varchar(20);not null"`
	Number int
	CTM   time.Time  `gorm:"column:ctm;default:CURRENT_TIMESTAMP" json:"ctm"`
	UTM  time.Time   `gorm:"column:utm;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"utm"`
}

func (*Teacher) TableName() string{
	return "tb_teacher"
}

func AddTeacher(teacher Teacher)bool{
	mysql.GetDb("gin").Exec(`insert tb_teacher(name,age,class,gender) values(?,?,?,?)`,
		teacher.Name,teacher.Age,teacher.Class,teacher.Gender)
	return true
}

func GetTeacher(id int)(teacher Teacher){
	mysql.GetDb("gin").Raw("select * from tb_teacher where id=?",id).Scan(&teacher)
	return
}

func PutTeacher(teacher Teacher)bool{
	mysql.GetDb("gin").Exec("update tb_teacher set name=?,gender=?,age=?,class=? where id=?",
		teacher.Name,teacher.Gender,teacher.Age,teacher.Class,teacher.ID)
	return true
}

func DeleteTeacher(id int)bool{
	mysql.GetDb("gin").Exec("delete from tb_teacher where id=?",id)
	return true
}
