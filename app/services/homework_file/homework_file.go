package homework_file

import (
	"classroom-system/app/models"
	"classroom-system/database"
	"errors"
)

func CreateHomeworkFile(homeworkFile *models.HomeworkFiles) error {
	homeworkFile.Status = 1

	err := database.GetMysql().Create(homeworkFile).Error
	if err != nil {
		return errors.New("创建作业文件失败")
	}

	return nil
}
