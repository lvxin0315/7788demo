package demo

type DataDemo struct {
	// public
	Name                 string // 姓名
	Mobile               string // 手机号
	Gender               int    // 性别，0-未知，1-男，2-女
	Province, City, Area string // 用户所在省市区
	// private
	code string // 内部编码
}

// InitCode 自动生成code
func (data *DataDemo) InitCode() string {
	data.code = "abc"
	return data.code
}
