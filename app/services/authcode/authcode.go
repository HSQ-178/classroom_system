package authcode

import (
	"classroom-system/app/models"
	"classroom-system/database"
	"errors"
	"time"
)

func check(authcodeReq *models.AuthcodeReq) error {
	if authcodeReq.TeacherCard == "" {
		return errors.New("老师编号不能为空")
	}
	if authcodeReq.Grade == "" {
		return errors.New("年级不能为空")
	}
	if authcodeReq.College == "" {
		return errors.New("学院不能为空")
	}
	if authcodeReq.Major == "" {
		return errors.New("专业不能为空")
	}
	if authcodeReq.Class == "" {
		return errors.New("班级不能为空")
	}
	if authcodeReq.CourseId == 0 {
		return errors.New("课程编号不能为空")
	}
	if authcodeReq.Authcode == 0 {
		return errors.New("验证码不能为空")
	}
	return nil
}

// 添加签到码
func AddAuthcode(authcodeReq *models.AuthcodeReq) error {

	err := check(authcodeReq)
	if err != nil {
		return err
	}

	db := database.GetMysql().Table("authcodes")
	authcode := models.Authcode{
		TeacherCard: authcodeReq.TeacherCard,
		Grade:       authcodeReq.Grade,
		College:     authcodeReq.College,
		Major:       authcodeReq.Major,
		Class:       authcodeReq.Class,
		CourseId:    authcodeReq.CourseId,
		Authcode:    authcodeReq.Authcode,
		CreateTime:  time.Now(),
	}

	err = db.Create(&authcode).Error
	if err != nil {
		return err
	}
	return nil
}

// 查找签到码
func GetAuthcode(authcodeReq *models.AuthcodeReq) (models.AuthcodeResp, error) {
	var authcodeResp models.AuthcodeResp

	err := check(authcodeReq)
	if err != nil {
		return authcodeResp, err
	}

	db := database.GetMysql().Table("authcodes")
	if authcodeReq.TeacherCard != "" {
		db = db.Where("teacher_card = ?", authcodeReq.TeacherCard)
	}
	if authcodeReq.Grade != "" {
		db = db.Where("grade = ?", authcodeReq.Grade)
	}
	if authcodeReq.College != "" {
		db = db.Where("college = ?", authcodeReq.College)
	}
	if authcodeReq.Major != "" {
		db = db.Where("major = ?", authcodeReq.Major)
	}
	if authcodeReq.Class != "" {
		db = db.Where("class = ?", authcodeReq.Class)
	}
	if authcodeReq.CourseId != 0 {
		db = db.Where("course_id = ?", authcodeReq.CourseId)
	}
	if authcodeReq.Authcode != 0 {
		db = db.Where("authcode = ?", authcodeReq.Authcode)
	}
	//if authcodeReq.CreateTime.IsZero() {
	//	db = db.Where("DATE(create_time) = ?", authcodeReq.CreateTime)
	//}
	err = db.First(&authcodeResp).Error
	if err != nil {
		return authcodeResp, err
	}
	return authcodeResp, nil
}
