package course

import (
	"classroom-system/app/models"
	"classroom-system/database"
	"errors"
)

// GetAllCourse 获取该老师所有课程
func GetAllCourse(c *models.CourseReq) (models.CourseListResp, error) {
	var courseList models.CourseListResp
	var err error

	//if c.TeacherCard == "" {
	//	return courseList, errors.New("老师编号不能为空")
	//}

	db := database.GetMysql().Table("courses")

	if c.Id != 0 {
		db.Where("id = ?", c.Id)
	}

	if c.College != "" {
		db.Where("college = ?", c.College)
	}

	if c.Major != "" {
		db.Where("major = ?", c.Major)
	}

	if c.Class != "" {
		db.Where("class = ?", c.Class)
	}

	if c.Grade != "" {
		db.Where("grade = ?", c.Grade)
	}

	if c.TeacherCard != "" {
		db.Where("teacher_card = ?", c.TeacherCard)
	}

	if c.StudentCard != "" {
		db.Where("student_card = ?", c.StudentCard)
	}

	if c.StudentName != "" {
		db.Where("student_name = ?", c.StudentName)
	}

	err = db.Find(&courseList.CourseResp).Count(&courseList.TotalCount).Error

	if err != nil {
		return courseList, errors.New("查询失败")
	}

	return courseList, nil
}
