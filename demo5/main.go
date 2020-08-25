package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"reflect"
)

type Goods struct {
	Title     string `json:"title"`      //商品名称
	BrandName string `json:"brand_name"` //品牌
}

var mapping = `{
  "mappings": {
    "properties": {
        "title": {
          "type": "text"
        },
        "brand_name": {
          "type": "text"
        }
      }
  }
}`

var indexKey = "goods-item"

func main() {
	ctx := context.Background()
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		//由于我的elastic使用的docker单节点安装，所以禁用
		//sniff为true时，客户端会去嗅探整个集群的状态，把集群中其它机器的ip地址加到客户端中
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	//查看version
	version, _ := client.ElasticsearchVersion("http://localhost:9200")
	fmt.Println("version:", version)
	//判断index是不是存在
	exists, err := client.IndexExists(indexKey).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(exists)
	if !exists {
		//创建index
		_, err = client.CreateIndex(indexKey).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
	//TODO 添加数据
	//addGoods(client, ctx)

	//TODO 查询方式
	//查询条件,MatchQuery 计算积分的单字段
	//query := elastic.NewMatchQuery("title", "一次性竹牙签")

	//MultiMatchQuery 计算积分的组合字段
	//query := elastic.NewMultiMatchQuery("唐宗筷 一次性竹牙签", "title", "brand_name")

	//自己写query, brand_name必须是"唐宗筷", title 包含
	query := elastic.NewRawStringQuery(`{
		"bool" : {
			"must" : [
				{
					"match" : { "brand_name" : "唐宗筷" }
				}
			],
			"should" : [
				{
					"match" : { "title" : "唐宗筷" },
					"match" : { "title" : "一次性竹牙签" }
				}
			],
			"minimum_should_match" : 1,
			"boost" : 1.0
		}
	}`)
	//排序规则
	scoreSort := elastic.NewScoreSort()
	scoreSort.Order(false)
	//查询
	searchResult, err := client.Search().Index(indexKey).
		Query(query).
		SortBy(scoreSort).
		Pretty(true).
		From(10).
		Size(10).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	//输出结果看看
	for _, item := range searchResult.Each(reflect.TypeOf(&Goods{})) {
		if t, ok := item.(*Goods); ok {
			fmt.Printf("item by %s: %s\n", t.Title, t.BrandName)
			//fmt.Println(t)
		}
	}
	fmt.Printf("Found a total of %d goods\n", searchResult.TotalHits())
	//看看原结果
	for _, i := range searchResult.Hits.Hits {
		fmt.Println(*i.Score)
	}

	//TODO 为了方便调试，删除index
	//deleteIndex(client, ctx)
}

//利用item.txt批量添加demo数据
func addGoods(client *elastic.Client, ctx context.Context) {
	b, err := ioutil.ReadFile("item.txt")
	if err != nil {
		panic(err)
	}
	bList := bytes.Split(b, []byte("\n"))
	for i, goodsInfoByte := range bList {
		//文本是用逗号分隔的
		gInfo := bytes.Split(goodsInfoByte, []byte(","))
		g := &Goods{
			Title:     string(gInfo[1]),
			BrandName: string(gInfo[0]),
		}
		_, _ = client.Index().Index(indexKey).BodyJson(g).Do(ctx)
		fmt.Println(i, "：", g.Title, g.BrandName, " over")
	}
}

//删除index
func deleteIndex(client *elastic.Client, ctx context.Context) {
	_, _ = client.DeleteIndex(indexKey).Do(ctx)
}
