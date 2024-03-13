package models

// Teachers 老师模型
type Teachers struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	TeacherCard string `json:"teacher_card"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}

// TeachersRegisterReq  老师注册请求参数
type TeachersRegisterReq struct {
	TeacherCard string `json:"teacher_card"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}

// TeachersLoginReq 老师登录请求参数
type TeachersLoginReq struct {
	TeacherCard string `json:"teacher_card"`
	Password    string `json:"password"`
}

// TeachersLoginRes 老师登录响应参数
type TeachersLoginRes struct {
	Token string `json:"token"`
}
