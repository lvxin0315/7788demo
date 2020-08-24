package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"strings"
)

//公司名称
var companyNameList = []string{
	"大铁棒子科技公司",
	"大铁棒子科技",
	"大铁棒子",
}

//产品品牌，一般都是非正规词语
var brandList = []string{
	"大鹅",
	"大娘",
	"阿拉拉",
	"京大西",
}

//demo的产品名称
var goodsNameList = []string{
	"大鹅四开门省电静音大冰柜 经典款",
	"大娘 詹姆斯使节12代 篮球鞋",
	"大娘 Jordan1 mid AJ1男鞋女鞋乔1黑红黑绿脚趾芝加哥湖人禁穿鸳鸯篮球鞋 鸳鸯拼接554724-124/554725-124 40.5",
	"阿拉拉 天选 15.6英寸游戏笔记本电脑(新锐龙 7nm 8核 R7-4800H 16G 512GSSD GTX1650Ti 4G 144Hz)元气蓝",
	"京大西同款 阿拉拉 天选 15.6英寸游戏笔记本电脑(新锐龙 7nm 8核 R7-4800H 16G 512GSSD RTX2060 6G 144Hz)元气蓝",
}

func main() {
	x := gojieba.NewJieba()
	defer x.Free()
	//添加特殊词
	for _, w := range companyNameList {
		x.AddWord(w)
	}
	for _, w := range brandList {
		x.AddWord(w)
	}
	//搜索引擎模式
	for _, gn := range goodsNameList {
		words := x.CutForSearch(gn, true)
		fmt.Println(strings.Join(words, "/"))
	}
}
