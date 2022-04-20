package job

import (
	"encoding/json"
	"fmt"
	"github.com/lvxin0315/7788demo/example2/model"
	"github.com/lvxin0315/7788demo/example2/param"
	"github.com/lvxin0315/7788demo/example2/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	yonghuiSearchUrl      = `https://activity.yonghuivip.com/api/web/category/category/sku/700`
	yonghuiDetailUrl      = `https://activity.yonghuivip.com/api/web/product/sku/detail/780`
	yonghuiAllCategoryUrl = `https://activity.yonghuivip.com/api/web/category/category/700`
)

func YonghuiData() {
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

	// 获取所有分类
	allCategoryData, err := getAllCategory()
	if err != nil {
		panic(err)
	}

	var goodsModelList []model.Goods

	for _, category := range allCategoryData.Data.Categorys {
		// 通过二级分类的id获取商品信息
		dataList, err := yonghuiGetGoodsByCategory(category)
		if err != nil {
			panic(err)
		}
		// 留档保存
		goodsModelList = append(goodsModelList, dataList...)
		// 数据库
		fmt.Println(category.Categoryname, " 开始入库")
		for _, goods := range dataList {
			err = db.Create(&goods).Error
			if err != nil {
				fmt.Println("入库error：", err)
			}
		}
	}

	// 留档
	goodsModelListBytes, _ := json.Marshal(goodsModelList)
	_ = ioutil.WriteFile("yonghui.json", goodsModelListBytes, 0755)

	fmt.Println("总计处理：goodsModelList.len: ", len(goodsModelList))

	//categoryData, err := getYonghuiGoodsList()
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, sku := range categoryData.Data.Skus {
	//	fmt.Println(sku.SkuCode)
	//}

	//detailData, err := yonghuiGoodsDetail("R-1421543")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(detailData.Data.Title)

}

// 获取所有的分类
func getAllCategory() (data model.AllCategoryResponse, err error) {
	// ?platform=wechatminiprogram&v=8.4.5.33&cityid=75&sellerid=7&shopid=9355&lat=41.75868810456091&lng=123.33596419851551
	params := make(map[string]interface{})
	params["platform"] = "wechatminiprogram"
	params["v"] = "8.4.5.33"
	params["cityid"] = "75"
	params["sellerid"] = param.YonghuiSellerid
	params["shopid"] = param.YonghuiShopid
	params["lat"] = param.YonghuiLat
	params["lng"] = param.YonghuiLng
	params["abdata"] = url.QueryEscape(`{"category_tree_data_abt":"10"}`)

	res, err := utils.HttpHandle(yonghuiAllCategoryUrl+utils.MapToParamsString(params), http.MethodGet, nil, nil, nil)

	//fmt.Println(string(res))

	if err != nil {
		return
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return
	}

	return
}

// 获取商品列表
// 传入分类id
func getYonghuiGoodsList(categoryId string, page string) (data model.CategoryResponse, err error) {
	params := make(map[string]interface{})
	params["platform"] = "wechatminiprogram"
	params["v"] = "8.4.5.33"
	params["cityid"] = "75"
	params["page"] = page
	params["pagecount"] = "20"
	params["sellerid"] = param.YonghuiSellerid
	params["shopid"] = param.YonghuiShopid
	params["lat"] = param.YonghuiLat
	params["lng"] = param.YonghuiLng
	params["categoryid"] = categoryId

	res, err := utils.HttpHandle(yonghuiSearchUrl+utils.MapToParamsString(params), http.MethodGet, nil, nil, nil)

	//fmt.Println(string(res))

	if err != nil {
		return
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return
	}

	return
}

// 获取商品详情
// 使用sku code
func yonghuiGoodsDetail(skuCode string) (data model.DetailResponse, err error) {
	params := make(map[string]interface{})
	params["platform"] = "wechatminiprogram"
	params["sdk_version"] = "2.13.2"
	params["code"] = skuCode
	params["shopid"] = param.YonghuiShopid

	res, err := utils.HttpHandle(yonghuiDetailUrl+utils.MapToParamsString(params), http.MethodGet, nil, nil, nil)

	//fmt.Println(string(res))

	if err != nil {
		return
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return
	}

	return
}

// 通过分类获取商品列表
func yonghuiGetGoodsByCategory(category model.AllCategoryResponseCategory) (dataList []model.Goods, err error) {
	fmt.Println("开始处理一级分类：", category.Categoryname)
	for _, subCategory := range category.Subcategory {
		fmt.Println("开始处理二级分类：", subCategory.Categoryname)
		goodsModelList, err := yonghuiGetGoodsBySubCategoryId(subCategory.Categoryid)
		if err != nil {
			return nil, err
		}
		// 将一级分类名称添加到数据中
		for _, goods := range goodsModelList {
			goods.CategoryName = category.Categoryname + " - " + goods.CategoryName
			// 添加返回值中
			dataList = append(dataList, goods)
		}
	}

	return
}

// 通过二级分类，并使用分页参数获取所有商品
func yonghuiGetGoodsBySubCategoryId(subCategoryId string) (dataList []model.Goods, err error) {
	// 通过page = 0 获取页数情况
	categoryResult, err := getYonghuiGoodsList(subCategoryId, "0")
	if err != nil {
		return
	}
	// 处理第一页数据
	var yonghuiDataList []model.YonghuiSku
	for _, sku := range categoryResult.Data.Skus {
		yonghuiDataList = append(yonghuiDataList, sku)
	}
	// 从第二页开始继续处理
	if categoryResult.Data.Totalpage > 1 {
		for i := 2; i < categoryResult.Data.Totalpage; i++ {
			otherCategoryResult, err := getYonghuiGoodsList(subCategoryId, fmt.Sprintf("%d", i))
			if err != nil {
				return nil, err
			}
			for _, sku := range otherCategoryResult.Data.Skus {
				yonghuiDataList = append(yonghuiDataList, sku)
			}
		}
	}
	fmt.Println("商品数：", len(yonghuiDataList))
	// 处理所有sku数据
	for _, item := range yonghuiDataList {
		goodsModel, err := yonghuiGoodsDetailForModelGoods(item.SkuCode)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, goodsModel)
	}
	return
}

// 使用商品详情，返回model.good格式
func yonghuiGoodsDetailForModelGoods(skuCode string) (m model.Goods, err error) {
	fmt.Println("处理商品，code：", skuCode)
	data, err := yonghuiGoodsDetail(skuCode)
	if err != nil {
		return
	}

	m.CategoryName = data.Data.SecondSellercategoryVo.Name
	m.StoreName = data.Data.Title
	m.Keyword = ""
	m.UnitName = ""
	m.StoreInfo = data.Data.Subtitle
	// 主图处理
	var mainImageList []string
	for _, mainimg := range data.Data.Mainimgs {
		mainImageList = append(mainImageList, mainimg.Imgurl)
	}
	m.SliderImage = strings.Join(mainImageList, ",")
	m.SpecType = 0
	// 第一章轮播作为主图
	if len(mainImageList) > 0 {
		m.Pic = mainImageList[0]
	}
	m.Price = float64(data.Data.Price.Value / 100)
	return
}
