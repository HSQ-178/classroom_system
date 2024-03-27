package teacher

import (
	"classroom-system/app/models"
	"classroom-system/database"
	"classroom-system/utils"
	"classroom-system/utils/jwt"
	"errors"

	"gorm.io/gorm"
)

func check(t *models.TeachersRegisterReq) error {
	if t.College == "" {
		return errors.New("学院不能为空")
	}
	if t.Major == "" {
		return errors.New("专业不能为空")
	}

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
		College:     t.College,
		Major:       t.Major,
		TeacherCard: t.TeacherCard,
		Name:        t.Name,
		Phone:       t.Phone,
		Password:    utils.MD5(t.Password),
	}

	//插入数据库
	return database.GetMysql().Create(&Teacher).Error
}

// TeacherLogin 老师登录
func TeacherLogin(t *models.TeachersLoginReq) (models.TeachersLoginResp, error) {
	var loginResp models.TeachersLoginResp
	var teacherInfo models.Teachers

	err := database.GetMysql().Table("teachers").
		Where("teacher_card = ? AND password = ?", t.TeacherCard, utils.MD5(t.Password)).
		First(&teacherInfo).Error
	if err != nil {
		return loginResp, errors.New("该老师不存在")
	}

	if teacherInfo.TeacherCard != t.TeacherCard {
		return loginResp, errors.New("教师编号错误")
	}

	if teacherInfo.Password != utils.MD5(t.Password) {
		return loginResp, errors.New("密码错误")
	}

	token, err := jwt.GenerateToken(teacherInfo.Id, teacherInfo.TeacherCard, teacherInfo.Password)
	if err != nil {
		return loginResp, errors.New("生成token失败")
	}

	//返回老师信息
	loginResp.Token = token
	loginResp.Teacher = models.Teachers{
		Id:          teacherInfo.Id,
		College:     teacherInfo.College,
		Major:       teacherInfo.Major,
		TeacherCard: teacherInfo.TeacherCard,
		Name:        teacherInfo.Name,
		Phone:       teacherInfo.Phone,
		Password:    teacherInfo.Password,
	}

	return loginResp, nil
}
