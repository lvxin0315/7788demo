package main

import (
	"encoding/json"
	"fmt"
	"github.com/lvxin0315/7788demo/tt01/config"
	"github.com/lvxin0315/7788demo/tt01/model"
	"github.com/lvxin0315/7788demo/tt01/runner"
	"io/ioutil"
	"os"
)

func main() {
	// 使用环境变量设置代理
	config.SetProxyEnv()

	// TODO

	err, res := readJsonFileToStruct(config.ResponseJsonDir + string(os.PathSeparator) + "demo.json")
	if err != nil {
		panic(err)
	}

	err, total := runner.BatchWork(res)
	if err != nil {
		panic(err)
	}

	fmt.Println("已下载：", total)
}

// 读取指定json文件，并解析成struct
func readJsonFileToStruct(jsonFileName string) (err error, res model.VideoListResponse) {
	bs, err := ioutil.ReadFile(jsonFileName)
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &res)
	if err != nil {
		return
	}
	return
}
