package student

import (
	"classroom-system/app/models"
	"classroom-system/database"
	"classroom-system/utils"
	"errors"
)

// GetStudent 获取该老师所有学生
func GetStudent(s *models.StudentReq) (models.StudentListResp, error) {
	var studentList models.StudentListResp

	//根据教师编号查询所有学生
	db := database.GetMysql().Table("students s").
		Select("s.grade, s.college, s.major, s.class, s.student_card, s.name").
		Joins("JOIN majors m ON s.major = m.major").
		Joins("JOIN courses c ON m.id = c.major_id").
		Joins("JOIN teachers t ON c.teacher_card = t.teacher_card")
	//Distinct().
	//Find(&studentList.Response).Error

	if s.Id != 0 {
		db.Where("s.id = ?", s.Id)
	}

	if s.TeacherCard != "" {
		db.Where("t.teacher_card = ?", s.TeacherCard)
	}

	if s.Grade != 0 {
		db.Where("s.grade = ?", s.Grade)
	}

	if s.College != "" {
		db.Where("s.college Like ?", "%"+s.College+"%")
	}

	if s.Major != "" {
		db.Where("s.major Like ?", "%"+s.Major+"%")
	}

	if s.Class != "" {
		db.Where("s.class Like ?", "%"+s.Class+"%")
	}

	if s.StudentCard != "" {
		db.Where("s.student_card Like ?", "%"+s.StudentCard+"%")
	}

	if s.Name != "" {
		db.Where("s.name Like ?", "%"+s.Name+"%")
	}

	err := db.Distinct("student_card").Count(&studentList.TotalCount).Error
	if err != nil {
		return studentList, errors.New("查询总数失败")
	}

	if s.Pagination.Page > 0 && s.Pagination.PageSize > 0 {
		db.Scopes(utils.Paginate(&s.Pagination))
	}

	err = db.Distinct("student_card").First(&studentList.Response).Error

	if err != nil {
		return studentList, errors.New("查询失败")
	}

	return studentList, nil
}
