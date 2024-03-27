package record

import (
	"classroom-system/app/models"
	"classroom-system/database"
	"classroom-system/utils"
	"errors"
)

func check(r *models.RecordReq) error {
	if r.TeacherCard == "" {
		return errors.New("教师编号不能为空")
	}

	if r.Grade == "" {
		return errors.New("年级不能为空")
	}

	if r.College == "" {
		return errors.New("学院不能为空")
	}

	if r.Major == "" {
		return errors.New("专业不能为空")
	}

	if r.Class == "" {
		return errors.New("班级不能为空")
	}

	if r.CourseId == 0 {
		return errors.New("课程编号不能为空")
	}
	return nil
}

// GetNotSignIn  查询未签到学生
func GetNotSignInStudentList(r *models.RecordReq) (models.NotSignInResp, error) {
	var notSignInList models.NotSignInResp

	err := check(r)
	if err != nil {
		return notSignInList, err
	}

	db := database.GetMysql().Model(&models.Student{}).Not("student_card = (?)",
		database.GetMysql().Table("records").Select("student_card").
			Where("teacher_card = ? AND course_id = ? AND (create_time = ?", r.TeacherCard, r.CourseId, r.CreateTime)).
		Where("grade = ? AND college = ? AND major = ? AND class = ?", r.Grade, r.College, r.Major, r.Class)

	// 获取未签到学生的数量
	var totalCount int64
	err = db.Count(&totalCount).Error
	if err != nil {
		return notSignInList, errors.New("获取未签到学生数量失败")
	}
	notSignInList.TotalCount = totalCount

	//分页
	if r.Paginate.Page > 0 && r.Paginate.PageSize > 0 {
		db.Scopes(utils.Paginate(&r.Paginate))
	}

	// 获取未签到学生列表
	err = db.Scan(&notSignInList.StudentList).Error
	if err != nil {
		return notSignInList, errors.New("获取未签到学生列表失败")
	}

	notSignInList.TotalSize = len(notSignInList.StudentList)

	return notSignInList, nil

}
