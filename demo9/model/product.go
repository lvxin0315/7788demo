package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

//通过结构体定义table名称
func (m *Product) TableName() string {
	return "t_product"
}
