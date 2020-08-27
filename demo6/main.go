package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/fedesog/webdriver"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"strings"
	"time"
)

var sogouUrl = "https://weixin.sogou.com"
var sogouWxUrl = sogouUrl + "/weixin?query=%s&_sug_type_=&s_from=input&_sug_=n&type=2&page=%d&ie=utf8"
var keyword = "elasticsearch"
var maxPage = 10
var gormDB *gorm.DB

//chromedriver 文件地址
var driverPath = "/Users/lvxin/IdeaProjects/SeleniumDemo/src/test/resources/chromedriver"

//存放内容的结构
type wxContent struct {
	gorm.Model  `json:"-"`
	Title       string `json:"title"`
	ContentText string `json:"content_text" gorm:"type:longtext"`
	ContentHtml string `json:"content_html" gorm:"type:longtext"`
}

func (m *wxContent) TableName() string {
	return "wx_data"
}

//webdriver
var chromeDriver *webdriver.ChromeDriver

//存放所有获取内容
var wxContentList []*wxContent

//解析flag参数
func initFlag() {
	flag.StringVar(&keyword, "keyword", "elasticsearch", "关键词")
	flag.Parse()
}

func main() {
	initFlag()
	//腾讯反爬虫还是挺nb的，我们使用webdriver来获取内容
	chromeDriver = webdriver.NewChromeDriver(driverPath)
	err := chromeDriver.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		fmt.Println(err)
		return
	}
	searchPageUrl(session)
	//把内容保存到文件里
	if len(wxContentList) > 0 {
		//saveWxContentListForJsonFile()
		saveWxContentListForMysql()
		saveWxContentListForElasticsearch()
	}
	session.Delete()
	chromeDriver.Stop()
}

//解析搜索页面链接，并打开页面获取里面内容，再进行处理
func searchPageUrl(session *webdriver.Session) {
	for i := 1; i <= maxPage; i++ {
		//浏览器打开列表页
		err := session.Url(fmt.Sprintf(sogouWxUrl, keyword, i))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(3 * time.Second)
		eleList, err := session.FindElements(webdriver.CSS_Selector, ".txt-box >h3 >a")
		if err != nil {
			fmt.Println(err)
			return
		}
		//记录当前页面的连接
		var urlList []string
		for _, ele := range eleList {
			u, err := ele.GetAttribute("href")
			if err != nil {
				continue
			}
			urlList = append(urlList, u)
		}
		//分别访问详情页面，并解析内容
		for _, u := range urlList {
			//打开页面
			session.Url(u)
			time.Sleep(3 * time.Second)
			h, err := session.Source()
			if err != nil {
				continue
			}
			//解析页面内容
			wxC := contentPageQuery([]byte(h))
			if wxC != nil {
				wxContentList = append(wxContentList, wxC)
			}
		}
	}
	return
}

//解析详情页面
func contentPageQuery(content []byte) *wxContent {
	//fmt.Println(string(content))
	//ioutil.WriteFile("1.html", content, 0755)
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(content))
	if err != nil {
		fmt.Println("goquery:", err)
		return nil
	}
	//title, Text()会去掉html标签内容
	title := doc.Find("#activity-name").Text()
	//去掉空格
	title = strings.TrimSpace(title)
	//内容
	contentText := doc.Find(".rich_media_content").Text()
	contentHtml, _ := doc.Find(".rich_media_content").Html()

	//fmt.Println(title)
	//fmt.Println(contentHtml)

	if title == "" || contentHtml == "" {
		return nil
	}
	return &wxContent{
		Title:       title,
		ContentText: contentText,
		ContentHtml: contentHtml,
	}
}

//把内容转成json并保存到文件中
func saveWxContentListForJsonFile() {
	//keyword做文件名
	b, err := json.Marshal(wxContentList)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s.json", keyword), b, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//初始化数据库
func initMysqlDB(conn string) {
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(5)
	gormDB = db
}

//保存内容到mysql
func saveWxContentListForMysql() {
	//初始化连接mysql
	initMysqlDB(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"root",
		"127.0.0.1",
		3306,
		"wx_data",
	))
	//初始化建表，表存在就不用了
	db := gormDB.New()
	if !db.HasTable(wxContent{}) {
		db.CreateTable(wxContent{})
	}
	//保存到db
	for _, item := range wxContentList {
		err := db.Save(item).Error
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

//保存内容到elasticsearch
func saveWxContentListForElasticsearch() {
	var indexKey = "wx-item"
	var mapping = `{
  "mappings": {
    "properties": {
        "title": {
          "type": "text"
        },
        "content_text": {
          "type": "text"
        },
		"content_html": {
          "type": "text"
        }
      }
  }
}`
	//es操作可以参考demo5
	ctx := context.Background()
	client, err := elastic.NewClient(
		elastic.SetURL("http://172.16.0.203:9200"),
		elastic.SetSniff(false))
	if err != nil {
		fmt.Println("NewClient: ", err)
		return
	}
	//判断index是不是存在
	exists, err := client.IndexExists(indexKey).Do(ctx)
	if err != nil {
		fmt.Println("IndexExists: ", err)
		return
	}
	if !exists {
		//创建index
		_, err = client.CreateIndex(indexKey).BodyString(mapping).Do(ctx)
		if err != nil {
			fmt.Println("CreateIndex: ", err)
			return
		}
	}

	for _, wxC := range wxContentList {
		_, err = client.Index().Index(indexKey).BodyJson(wxC).Do(ctx)
		if err != nil {
			fmt.Println("Do: ", err)
			return
		}
	}
}
