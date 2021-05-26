package model

// Comment 评论
type Comment struct {
	Avatar   string `json:"avatar"`   // 头像
	Name     string `json:"name"`     // 名字
	Score    int    `json:"score"`    // 分数
	Describe string `json:"describe"` // 描述
}
