package service

import (
	"encoding/json"
	"fmt"
	"github.com/lvxin0315/7788demo/example1/model"
	"io/ioutil"
)

// IndexData 首页数据
var IndexData model.IndexData

func init() {
	err := _readJsonFile("index", &IndexData)
	if err != nil {
		panic(err)
	}
}

// 读取json文件
func _readJsonFile(jsonName string, data interface{}) error {
	jsonBytes, err := ioutil.ReadFile(fmt.Sprintf("data/%s.json", jsonName))
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, data)
	if err != nil {
		return err
	}
	return nil
}
