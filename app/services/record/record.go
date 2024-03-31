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

	if r.Status == 0 {
		return errors.New("状态不能为空")
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

	db := database.GetMysql().Model(&models.Student{}).
		Not("student_card in (?)",
			database.GetMysql().
				Table("records").
				Select("student_card").
				Where("teacher_card = ? ", r.TeacherCard).
				Where("course_id = ? ", r.CourseId).
				Where("create_time = ? ", r.CreateTime)).
		Where("grade = ? AND college = ? AND major = ? AND class = ? AND status = ?", r.Grade, r.College, r.Major, r.Class, r.Status)

	// 获取未签到学生的数量
	err = db.Count(&notSignInList.TotalCount).Error
	if err != nil {
		return notSignInList, errors.New("获取未签到学生数量失败")
	}

	//分页
	if r.Paginate.Page > 0 && r.Paginate.PageSize > 0 {
		db.Scopes(utils.Paginate(&r.Paginate))
	}

	// 获取未签到学生列表
	err = db.Scan(&notSignInList.StudentList).Error
	if err != nil {
		return notSignInList, errors.New("获取未签到学生列表失败")
	}

	return notSignInList, nil

}

// GetSignInAndAbsence 查询签到或缺勤
func GetSignInAndAbsence(r *models.RecordReq) (signInAndAbsenceList models.RecordSignInRespList, err error) {

	err = check(r)
	if err != nil {
		return signInAndAbsenceList, err
	}

	db := database.GetMysql().Table("records")

	if r.TeacherCard != "" {
		db = db.Where("teacher_card = ?", r.TeacherCard)
	}

	if r.Grade != "" {
		db = db.Where("grade = ?", r.Grade)
	}

	if r.College != "" {
		db = db.Where("college = ?", r.College)
	}

	if r.Major != "" {
		db = db.Where("major = ?", r.Major)
	}

	if r.Class != "" {
		db = db.Where("class = ?", r.Class)
	}

	if r.CourseId != 0 {
		db = db.Where("course_id = ?", r.CourseId)
	}

	if r.Status != 0 {
		db = db.Where("status = ?", r.Status)
	}

	if r.CreateTime != "" {
		db = db.Where("create_time = ?", r.CreateTime)
	}

	err = db.Count(&signInAndAbsenceList.TotalCount).Error
	if err != nil {
		return signInAndAbsenceList, err
	}

	// 分页
	if r.Paginate.Page > 0 && r.Paginate.PageSize > 0 {
		db = db.Scopes(utils.Paginate(&r.Paginate))
	}

	err = db.Find(&signInAndAbsenceList.RecordResp).Error
	if err != nil {
		return signInAndAbsenceList, err
	}

	return signInAndAbsenceList, nil

}
