package models

// Teachers 老师模型
type Teachers struct {
	Id          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	TeacherCard string `form:"teacherCard" json:"teacherCard"`
	Name        string `json:"name"`
	Phone       string `json:"phone" gorm:"default: null"`
	Password    string `json:"password"`
}

// TeachersRegisterReq  老师注册请求参数
type TeachersRegisterReq struct {
	TeacherCard string `json:"teacherCard"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}

// TeachersLoginReq 老师登录请求参数
type TeachersLoginReq struct {
	TeacherCard string `json:"teacherCard"`
	Password    string `json:"password"`
}

// TeachersLoginResp 老师登录响应参数
type TeachersLoginResp struct {
	Token   string   `json:"token"`
	Teacher Teachers `json:"teacher"`
}
