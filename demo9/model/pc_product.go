package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//带有回调操作的结构体
type PCProduct struct {
	gorm.Model
	Code  string
	Price uint
}

func (m *PCProduct) TableName() string {
	return "t_pc_product"
}

//BeforeSave
func (m *PCProduct) BeforeSave() (err error) {
	fmt.Println("我是PCProduct.BeforeSave, code: ", m.Code)
	return nil
}

//AfterSave
func (m *PCProduct) AfterSave() (err error) {
	fmt.Println("我是PCProduct.AfterSave, code: ", m.Code)
	return nil
}

//BeforeCreate
func (m *PCProduct) BeforeCreate() (err error) {
	fmt.Println("我是PCProduct.BeforeCreate, code: ", m.Code)
	return nil
}

//AfterCreate
func (m *PCProduct) AfterCreate() (err error) {
	fmt.Println("我是PCProduct.AfterCreate, code: ", m.Code)
	return nil
}
