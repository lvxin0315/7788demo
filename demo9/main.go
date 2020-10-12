package main

//gorm操作

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lvxin0315/7788demo/demo9/model"
	"math/rand"
	"time"
)

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

	//2.连接池
	db.DB().SetMaxOpenConns(50) //最大连接数
	db.DB().SetMaxIdleConns(10) //最大空闲数
	//db.DB().SetConnMaxLifetime(time.Minute)      //最大连接数释放时间
	//db.DB().SetConnMaxIdleTime(10 * time.Second) //最大空闲连接释放时间

	//打开日志
	db.LogMode(true)

	migrate(db)

	addTestData(db)

	testSelect(db)

	//2.并行写入数据控，测试连接池效果
	//addBatchTestData(db)

	//3.回调操作效果
	addPcTestData(db)
	//防止程序退出
	select {}
}

//迁移
func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Product{}, &model.PCProduct{})
}

//添加test数据
func addTestData(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		err := db.Save(&model.Product{
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
	var productList []*model.Product
	err := db.Model(model.Product{}).Limit(20).Find(&productList).Error
	checkError(err)
	for _, item := range productList {
		//看看结果
		fmt.Println(item.Code)
	}
}

//批量添加test数据
func addBatchTestData(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30000; i++ {
		go func(i int) {
			err := db.New().Save(&model.Product{
				Code:  fmt.Sprintf("code-%d-%d", time.Now().Unix(), i),
				Price: uint(rand.Intn(100000)),
			}).Error
			if err != nil {
				fmt.Println(err)
			}
		}(i)
	}
}

//带有回调的curd
func addPcTestData(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		err := db.Save(&model.PCProduct{
			Code:  fmt.Sprintf("code-%d-%d", time.Now().Unix(), i),
			Price: uint(rand.Intn(100000)),
		}).Error
		if err != nil {
			fmt.Println(err)
		}
	}
}
