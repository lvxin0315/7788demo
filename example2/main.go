package main

import (
	"github.com/lvxin0315/7788demo/example2/job"
	"github.com/lvxin0315/7788demo/example2/param"
)

type yonghuiStore struct {
	StoreId  string
	SellerId string
}

var yonghuiStoreList = []yonghuiStore{
	{"9301", "7"},
	{"9744", "7"},
	{"9355", "7"},
	{"WTT075", "22"},
}

func main() {
	// 多参数处理
	//for _, store := range yonghuiStoreList {
	//	param.YonghuiShopid = store.StoreId
	//	param.YonghuiSellerid = store.SellerId
	//	job.YonghuiData()
	//}

	// 处理远程图片
	for _, store := range yonghuiStoreList {
		param.YonghuiShopid = store.StoreId
		param.YonghuiSellerid = store.SellerId
		job.YonghuiPicToOss()
	}

	// TODO
	//ossClient := utils.NewOssClient()
	//ossUrl, err := ossClient.UploadUrlFile("http://image.yonghuivip.com/product/B-274840/1585377118592a3f96f110ed8dfa05bfeb03ffdb8b238f5eb7588.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ossUrl)
}
