package model

import (
	"github.com/lvxin0315/7788demo/example2/param"
	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	CategoryName string  `json:"category_name"`                     // 分类名称
	StoreName    string  `json:"store_name"`                        // 商品名称
	Keyword      string  `json:"keyword"`                           // 关键词
	UnitName     string  `json:"unit_name"`                         // 单位
	StoreInfo    string  `json:"store_info"`                        // 商品简介
	SliderImage  string  `json:"slider_image" gorm:"type:text"`     // 轮播图
	SpecType     int     `json:"spec_type"`                         // 规格类型
	Pic          string  `json:"pic"`                               // 主图
	Price        float64 `json:"price" gorm:"type:decimal(8,2)"`    // 价格
	Cost         float64 `json:"cost" gorm:"type:decimal(8,2)"`     // 成本
	OtPrice      float64 `json:"ot_price" gorm:"type:decimal(8,2)"` // 原价
	Stock        int     `json:"stock"`                             // 库存

	OssSliderImage string `json:"oss_slider_image" gorm:"type:text"` // 轮播图(转存oss后地址)
	OssPic         string `json:"oss_pic"`                           // 主图(转存oss后地址)
}

func (m *Goods) TableName() string {
	return "pl_goods_" + param.YonghuiShopid
}
