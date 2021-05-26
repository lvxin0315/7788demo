package service

import (
	"github.com/lvxin0315/7788demo/example1/model"
	"github.com/siddontang/go-log/log"
)

// GetIndexDataService IndexData 首页数据
func GetIndexDataService() *model.IndexData {
	return &IndexData
}

// GetNavService 获取导航
func GetNavService(route string) []model.Nav {
	var navList []model.Nav
	err := _readJsonFile("nav", &navList)
	if err != nil {
		log.Errorln("GetNavService: ", err)
		return nil
	}
	// 根据当前导航设置选中状态
	for i, nav := range navList {
		if nav.Url == route {
			navList[i].Active = true
		}
	}
	return navList
}
