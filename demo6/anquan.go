package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lvxin0315/7788demo/demo6/output"
	"io/ioutil"
	"time"
)

type Category struct {
	Name     string     `json:"name"`
	Children []Category `json:"children"`
}

// 读取文件
func readAnquanPage(pageName string) []byte {
	content, err := ioutil.ReadFile(fmt.Sprintf("./anquan/%s.html", pageName))
	if err != nil {
		panic(err)
	}
	return content
}

// 解析html
func anquanHtmlParse(pageContent []byte) []Category {
	var categoryList []Category
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(pageContent))
	if err != nil {
		panic(err)
	}
	trSelection := doc.Find("tr")
	category := Category{}
	trSelection.Each(func(i int, selection *goquery.Selection) {
		category.Name = selection.Find(".f2 a").Text()
		//fmt.Println(category.Name)
		selection.Find(".f3 a").Each(func(i int, selection *goquery.Selection) {
			category.Children = append(category.Children, Category{
				Name:     selection.Text(),
				Children: nil,
			})
			//fmt.Println(selection.Text())
		})
		categoryList = append(categoryList, category)
	})
	return categoryList
}

func main() {
	var allCategoryList []Category
	category1 := Category{
		Name:     "安检",
		Children: nil,
	}
	page1Content := readAnquanPage("p1")
	category1.Children = anquanHtmlParse(page1Content)

	category2 := Category{
		Name:     "质监",
		Children: nil,
	}
	page2Content := readAnquanPage("p2")
	category2.Children = anquanHtmlParse(page2Content)

	category3 := Category{
		Name:     "职业技能鉴定",
		Children: nil,
	}
	page3Content := readAnquanPage("p3")
	category3.Children = anquanHtmlParse(page3Content)

	category4 := Category{
		Name:     "建筑类从业人员",
		Children: nil,
	}
	page4Content := readAnquanPage("p4")
	category4.Children = anquanHtmlParse(page4Content)

	allCategoryList = append(allCategoryList, category1, category2, category3, category4)

	// 输出json
	jsonByte, err := json.Marshal(allCategoryList)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("category.json", jsonByte, 0755)

	// 入库
	db, err := gorm.Open("mysql", "answer:answer@(39.101.170.235:3306)/answer?charset=utf8mb4&parseTime=True&loc=Local")
	tx := db.Begin()
	for i, categoryList := range allCategoryList {
		// 科目
		subject := output.AwYexamSubject{
			SubjectName: categoryList.Name,
			Status:      1,
			Weigh:       i,
			Createtime:  int(time.Now().Unix()),
		}
		err = tx.Save(&subject).Error
		if err != nil {
			panic(err)
		}
		// 一级章节
		for i2, category := range categoryList.Children {
			fUnit := output.AwYexamUnit{
				SubjectID:  subject.ID,
				UnitName:   category.Name,
				IsLast:     0,
				Status:     1,
				Createtime: int(time.Now().Unix()),
				Sort:       i2,
			}
			err = tx.Save(&fUnit).Error
			if err != nil {
				panic(err)
			}
			// 二级章节
			for i3, child := range category.Children {
				sUnit := output.AwYexamUnit{
					Pid:        fUnit.ID,
					SubjectID:  subject.ID,
					UnitName:   child.Name,
					IsLast:     0,
					Status:     1,
					Createtime: int(time.Now().Unix()),
					Sort:       i3,
				}
				err = tx.Save(&sUnit).Error
				if err != nil {
					panic(err)
				}
				// 三级章节-固定
				tUnit1 := output.AwYexamUnit{
					Pid:        sUnit.ID,
					SubjectID:  subject.ID,
					UnitName:   "新训",
					IsLast:     1,
					Status:     1,
					Createtime: int(time.Now().Unix()),
					Sort:       1,
				}
				err = tx.Save(&tUnit1).Error
				if err != nil {
					panic(err)
				}
				tUnit2 := output.AwYexamUnit{
					Pid:        sUnit.ID,
					SubjectID:  subject.ID,
					UnitName:   "复训",
					IsLast:     1,
					Status:     1,
					Createtime: int(time.Now().Unix()),
					Sort:       2,
				}
				err = tx.Save(&tUnit2).Error
				if err != nil {
					panic(err)
				}
				tUnit3 := output.AwYexamUnit{
					Pid:        sUnit.ID,
					SubjectID:  subject.ID,
					UnitName:   "换证",
					IsLast:     1,
					Status:     1,
					Createtime: int(time.Now().Unix()),
					Sort:       3,
				}
				err = tx.Save(&tUnit3).Error
				if err != nil {
					panic(err)
				}
			}
		}
	}

	tx.Commit()
}
