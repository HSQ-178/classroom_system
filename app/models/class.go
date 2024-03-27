package models

// Class 班级模型
type Class struct {
	Id      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Grade   string `json:"grade"`
	College string `json:"college"`
	Major   string `json:"major"`
	Name    string `json:"name"`
}
