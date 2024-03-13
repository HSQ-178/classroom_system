package teacher

import (
	"classroom-system/app/models"
	"classroom-system/database"
	"errors"
	"gorm.io/gorm"
)

// TeacherRegister 老师注册
func TeacherRegister(t *models.TeachersRegisterReq) error {

	//判断老师编号是否存在
	var tempTeacherCard models.Teachers
	err := database.GetMysql().Table("teachers").Where("teacher_card = ?", t.TeacherCard).First(tempTeacherCard).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该老师编号已存在")
	}

	Teacher := models.Teachers{
		TeacherCard: t.TeacherCard,
		Name:        t.Name,
		Phone:       t.Phone,
		Password:    t.Password,
	}

	//插入数据库
	return database.GetMysql().Create(&Teacher).Error
}
