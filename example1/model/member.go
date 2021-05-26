package model

// Member 成员
type Member struct {
	Avatar     string `json:"avatar"`     // 头像
	Name       string `json:"name"`       // 名字
	Profession string `json:"profession"` // 专业
	Describe   string `json:"describe"`   // 描述
}
