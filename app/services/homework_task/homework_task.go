package homework_task

import (
	"classroom-system/app/models"
	"classroom-system/database"
	"classroom-system/utils"
	"errors"
	"gorm.io/gorm/clause"
)

func CreateHomeworkTask(homeworkTask *models.HomeworkTask) (id int, err error) {
	homeworkTask.Status = 1

	result := database.GetMysql().Create(homeworkTask)
	if result.Error != nil {
		return 0, errors.New("创建作业任务失败")
	}

	return homeworkTask.Id, nil
}

func ListHomeworkTask(req *models.HomeworkTaskReq) (models.HomeworkTaskListResp, error) {
	var homeworkTaskList models.HomeworkTaskListResp

	db := database.GetMysql().Model(models.HomeworkTask{})

	if req.Id != 0 {
		db = db.Where("id = ?", req.Id)
	}

	if req.TeacherCard != "" {
		db = db.Where("teacher_card = ?", req.TeacherCard)
	}

	if req.CourseId != 0 {
		db = db.Where("course_id = ?", req.CourseId)
	}

	if req.Title != "" {
		db = db.Where("title like ?", "%"+req.Title+"%")
	}

	if req.Subscribe != "" {
		db = db.Where("subscribe like ?", "%"+req.Subscribe+"%")
	}

	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}

	db.Preload(clause.Associations)

	if req.Paginate.Page > 0 && req.Paginate.PageSize > 0 {
		db = db.Scopes(utils.Paginate(&req.Paginate))
	}

	err := db.Count(&homeworkTaskList.Total).Error
	if err != nil {
		return homeworkTaskList, errors.New("获取作业任务数量失败")
	}

	err = db.Find(&homeworkTaskList.Records).Error
	if err != nil {
		return homeworkTaskList, errors.New("获取作业任务列表失败")
	}

	return homeworkTaskList, nil
}
