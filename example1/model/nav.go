package model

// Nav 页面导航
type Nav struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	Active   bool   `json:"active"`
	Children []Nav  `json:"children"` // 子菜单
}
