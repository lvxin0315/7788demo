package main

import (
	"encoding/json"
	"fmt"
	"github.com/lvxin0315/7788demo/example1/model"
	"io/ioutil"
)

// 测试数据
func main() {
	navData()
	indexData()
}

func _writeJsonFile(jsonName string, data []byte) {
	_ = ioutil.WriteFile(fmt.Sprintf("data/%s.json", jsonName), data, 0755)
}

// 导航数组
func navData() {
	var navList []model.Nav
	navList = append(navList, model.Nav{
		Title:    "首页",
		Url:      "/",
		Active:   false,
		Children: nil,
	}, model.Nav{
		Title:    "关于",
		Url:      "/about",
		Active:   false,
		Children: nil,
	}, model.Nav{
		Title:    "课程",
		Url:      "/courses",
		Active:   false,
		Children: nil,
	}, model.Nav{
		Title:    "联系",
		Url:      "/contact",
		Active:   false,
		Children: nil,
	})
	data, _ := json.Marshal(navList)
	_writeJsonFile("nav", data)
}

// 首页基础数据
func indexData() {
	index := model.IndexData{
		HeadTitle:                "这里是绝佳的实践基地",
		HeadDescribe:             "在我们团队中，学习任何内容都用实践的方式，这样相比传统的方式更有效率、更容易理解透彻",
		CourseIntroduceTitle:     "内容永远在更新，目前累计3000节课程",
		CourseIntroduceDescribe:  "专业团队综合多家IT平台信息，清晰的分辨出最理想的方案",
		SubjectIntroduceTitle:    "学无止境",
		SubjectIntroduceDescribe: "并非苦海无涯",
		SubjectIntroduceLeftCategory: []string{
			"Go",
			"PHP",
			"Java",
			"JavaScript",
			"Python",
		},
		SubjectIntroduceRightCategory: []string{
			"Linux",
			"Vue",
			"Docker",
			"Swarm",
			"K8s",
		},
		TeachingMethodsTitle:    "实践方式",
		TeachingMethodsDescribe: "公司自有产品、合作公司的内部方案、创新大赛。。。",
		TeachingMethodsCards:    []model.TeachingMethodsCard{},
		MemberIntroduceTitle:    "成员介绍",
		MemberIntroduceDescribe: "成员介绍。。。。。。",
	}
	data, _ := json.Marshal(index)
	_writeJsonFile("index", data)
}
