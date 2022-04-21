package job

import (
	"fmt"
	"github.com/lvxin0315/7788demo/example2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 多个库合并到一个
func YonghuiDataToOneDB(mainM model.Goods, srcM ...model.Goods) {
	dsn := "root:root@tcp(127.0.0.1:3306)/cb_data?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	// 结构迁移
	_ = db.AutoMigrate(mainM)

	var allGoodsList []model.Goods
	// 查询所有子表中的数据
	for _, m := range srcM {
		var srcGoodsList []model.Goods

		err = db.Debug().Table(m.TableName()).Find(&srcGoodsList).Error
		if err != nil {
			panic(err)
		}
		for i := range srcGoodsList {
			if yonghuiHasGoodsList(&allGoodsList, &srcGoodsList[i]) {
				continue
			}
			// 添加到总合集中
			allGoodsList = append(allGoodsList, srcGoodsList[i])
		}
	}
	fmt.Println("总计数据：", len(allGoodsList))
	// 保存
	for _, goods := range allGoodsList {
		goods.Model = gorm.Model{}
		err = db.Table(mainM.TableName()).Create(&goods).Error
		if err != nil {
			panic(err)
		}
	}
}

// 判断重复
func yonghuiHasGoodsList(allGoodsList *[]model.Goods, goods *model.Goods) bool {
	for _, m := range *allGoodsList {
		if m.StoreName == goods.StoreName && m.CategoryName == goods.CategoryName {
			return true
		}
	}
	return false
}
