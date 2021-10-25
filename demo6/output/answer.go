package output

// AwYexamSubject 科目表
type AwYexamSubject struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null"`
	SubjectName string `gorm:"column:subject_name;type:varchar(50);not null"` // 科目名称
	Status      int    `gorm:"column:status;type:int(11);not null"`           // 1显示 0不显示
	Weigh       int    `gorm:"column:weigh;type:int(11);not null"`            // 排序
	Createtime  int    `gorm:"column:createtime;type:int(11);not null"`       // 创建时间
}

func (t *AwYexamSubject) TableName() string {
	return "aw_yexam_subject"
}

// AwYexamUnit 章节表
type AwYexamUnit struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null"`
	Pid        int    `gorm:"column:pid;type:int(11);not null"`           // 上级章节ID
	SubjectID  int    `gorm:"column:subject_id;type:int(11);not null"`    // 科目ID
	UnitName   string `gorm:"column:unit_name;type:varchar(50);not null"` // 章节名称
	IsLast     int    `gorm:"column:is_last;type:int(11);not null"`       // 是否为终极章节  1是  0否
	Status     int    `gorm:"column:status;type:int(11);not null"`        // 1 显示 0隐藏
	Createtime int    `gorm:"column:createtime;type:int(11);not null"`    // 创建时间
	Sort       int    `gorm:"column:sort;type:int(11);not null"`
}

func (t *AwYexamUnit) TableName() string {
	return "aw_yexam_unit"
}
