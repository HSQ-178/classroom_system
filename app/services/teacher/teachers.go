package teacher

import (
	"classroom-system/app/models"
	"classroom-system/database"
	"errors"

	"gorm.io/gorm"
)

func check(t *models.TeachersRegisterReq) error {

	if t.TeacherCard == "" {
		return errors.New("老师编号不能为空")
	}

	if t.Password == "" {
		return errors.New("密码不能为空")
	}

	if t.Name == "" {
		return errors.New("姓名不能为空")
	}

	return nil
}

// TeacherRegister 老师注册
func TeacherRegister(t *models.TeachersRegisterReq) error {

	var tempTeacherCard models.Teachers

	err := check(t)
	if err != nil {
		return err
	}

	//判断老师编号是否存在
	err = database.GetMysql().Table("teachers").Where("teacher_card = ?", t.TeacherCard).First(&tempTeacherCard).Error

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
