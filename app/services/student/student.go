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
		Joins("JOIN courses c ON s.major = c.major")

	if s.Id != 0 {
		db.Where("s.id = ?", s.Id)
	}

	if s.TeacherCard != "" {
		db.Where("c.teacher_card = ?", s.TeacherCard)
	}

	if s.Grade != "" {
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

	// 计算去重后的数据总数
	var totalCount int64
	err := db.Model(&models.Student{}).Distinct("s.id").Count(&totalCount).Error
	if err != nil {
		return studentList, errors.New("查询总数失败")
	}
	studentList.TotalCount = totalCount

	if s.Pagination.Page > 0 && s.Pagination.PageSize > 0 {
		db.Scopes(utils.Paginate(&s.Pagination))
	}

	err = db.Select("s.id, s.grade, s.college, s.major, s.class, s.student_card, s.name").Distinct().Find(&studentList.Response).Error

	if err != nil {
		return studentList, errors.New("查询失败")
	}

	studentList.TotalSize = len(studentList.Response)

	return studentList, nil
}
