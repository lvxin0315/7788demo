package main

//gorm操作

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

//通过结构体定义table名称
func (m *Product) TableName() string {
	return "t_product"
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//连接初始化
	db, err := gorm.Open("mysql", "root:root@/demo9?charset=utf8&parseTime=True&loc=Local")
	checkError(err)
	defer db.Close()

	//打开日志
	db.LogMode(true)

	migrate(db)

	addTestData(db)

	testSelect(db)
}

//迁移
func migrate(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

//添加test数据
func addTestData(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		err := db.Save(&Product{
			Code:  fmt.Sprintf("code-%d-%d", time.Now().Unix(), i),
			Price: uint(rand.Intn(100000)),
		}).Error
		if err != nil {
			fmt.Println(err)
		}
	}
}

//正常的查询
func testSelect(db *gorm.DB) {
	var productList []*Product
	err := db.Table("products").Limit(20).Find(&productList).Error
	checkError(err)
	for _, item := range productList {
		//看看结果
		fmt.Println(item.Code)
	}
}
