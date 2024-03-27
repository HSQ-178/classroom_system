package models

// Major 专业模型
type Major struct {
	Id        int    `json:"id"`
	CollegeId int    `json:"collegeId"`
	Name      string `json:"name"`
}
