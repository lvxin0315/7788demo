package job

import (
	"fmt"
	"github.com/lvxin0315/7788demo/example2/model"
	"github.com/lvxin0315/7788demo/example2/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"strings"
	"time"
)

var ossClient *utils.OssClient

func YonghuiPicToOss() {
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
	_ = db.AutoMigrate(model.Goods{})

	// 查询所有数据
	var goodsList []model.Goods
	err = db.Model(model.Goods{}).Find(&goodsList).Error
	if err != nil {
		panic(err)
	}

	// oss 客户端
	ossClient = utils.NewOssClient()
	fmt.Println("总计：", len(goodsList))
	for i, goods := range goodsList {
		if i%100 == 0 {
			fmt.Println("正在处理：", i)
		}
		err = yonghuiGoodsImageToOss(&goods)
		if err != nil {
			panic(err)
		}
		// 保存入库
		err = db.Save(goods).Error
		if err != nil {
			panic(err)
		}
	}
}

func yonghuiImageUrl(srcUrl string) (newUrl string, err error) {
	u, err := url.Parse(srcUrl)
	if err != nil {
		return
	}
	newUrl = u.Scheme + "://" + u.Host + u.Path

	return
}

// 图处理
func yonghuiGoodsImageToOss(goods *model.Goods) error {
	// 主图处理
	if goods.Pic != "" {
		mainPicUrl, err := yonghuiImageUrl(goods.Pic)
		if err != nil {
			return err
		}
		ossUrl, err := ossClient.UploadUrlFile(mainPicUrl)
		if err != nil {
			return err
		}
		goods.OssPic = ossUrl
	}

	// 轮播图处理
	if goods.SliderImage != "" {
		sliderImageList := strings.Split(goods.SliderImage, ",")
		var ossSliderImageList []string
		for _, sImageUrl := range sliderImageList {
			ossUrl, err := ossClient.UploadUrlFile(sImageUrl)
			if err != nil {
				return err
			}
			ossSliderImageList = append(ossSliderImageList, ossUrl)
		}
		goods.OssSliderImage = strings.Join(ossSliderImageList, ",")
	}

	return nil
}
