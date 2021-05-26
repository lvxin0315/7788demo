package model

// Course 课程
type Course struct {
	Title    string `json:"title"`    // 课程标题
	Describe string `json:"describe"` // 课程描述
	IsTop    bool   `json:"is_top"`   // 是否首页top显示
}
