package model

import "github.com/jinzhu/gorm"

// Course 课程
type Course struct {
	gorm.Model
	Title         string  `gorm:"column:title;type:varchar(100);not null"`              // 标题
	Cover         string  `gorm:"column:cover;type:varchar(100);not null"`              // 封面
	Summary       string  `gorm:"column:summary;type:varchar(255);not null"`            // 简介
	Keywords      string  `gorm:"column:keywords;type:varchar(100);not null"`           // 关键字
	Details       string  `gorm:"column:details;type:text;not null"`                    // 详情
	CategoryID    int     `gorm:"column:category_id;type:int(11) unsigned;not null"`    // 分类编号
	TeacherID     uint32  `gorm:"column:teacher_id;type:int(11) unsigned;not null"`     // 讲师编号
	OriginPrice   float64 `gorm:"column:origin_price;type:decimal(10,2);not null"`      // 原始价格
	MarketPrice   float64 `gorm:"column:market_price;type:decimal(10,2);not null"`      // 优惠价格
	VipPrice      float64 `gorm:"column:vip_price;type:decimal(10,2);not null"`         // 会员价格
	StudyExpiry   int     `gorm:"column:study_expiry;type:int(11) unsigned;not null"`   // 学习期限
	RefundExpiry  int     `gorm:"column:refund_expiry;type:int(11) unsigned;not null"`  // 退款期限
	Rating        float32 `gorm:"column:rating;type:float;not null"`                    // 用户评分
	Score         float32 `gorm:"column:score;type:float;not null"`                     // 综合得分
	Level         int     `gorm:"column:level;type:int(11) unsigned;not null"`          // 难度
	Attrs         string  `gorm:"column:attrs;type:varchar(1000);not null"`             // 扩展属性
	Featured      int     `gorm:"column:featured;type:int(11) unsigned;not null"`       // 推荐标识
	Published     uint32  `gorm:"column:published;type:int(11) unsigned;not null"`      // 发布标识
	Deleted       int     `gorm:"column:deleted;type:tinyint(1) unsigned;not null"`     // 删除标识
	ResourceCount int     `gorm:"column:resource_count;type:int(11) unsigned;not null"` // 资源数
	UserCount     int     `gorm:"column:user_count;type:int(11) unsigned;not null"`     // 学员数
	LessonCount   int     `gorm:"column:lesson_count;type:int(11) unsigned;not null"`   // 课时数
	PackageCount  int     `gorm:"column:package_count;type:int(11) unsigned;not null"`  // 套餐数
	ReviewCount   int     `gorm:"column:review_count;type:int(11) unsigned;not null"`   // 评价数
	ConsultCount  int     `gorm:"column:consult_count;type:int(11) unsigned;not null"`  // 咨询数
	FavoriteCount int     `gorm:"column:favorite_count;type:int(11) unsigned;not null"` // 收藏数
}
