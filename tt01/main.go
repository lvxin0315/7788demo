package main

import (
	"encoding/json"
	"fmt"
	"github.com/lvxin0315/7788demo/tt01/config"
	"github.com/lvxin0315/7788demo/tt01/model"
	"github.com/lvxin0315/7788demo/tt01/runner"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	// 使用环境变量设置代理
	config.SetProxyEnv()
	// 读取所有json文件
	jsonFileList := readJsonDirFile()
	if len(jsonFileList) == 0 {
		panic("无json文件")
	}
	// 遍历下载
	for _, jsonFile := range jsonFileList {
		err, res := readJsonFileToStruct(jsonFile)
		if err != nil {
			fmt.Println("readJsonFileToStruct error:", err)
			continue
		}
		err, total := runner.BatchWork(res)
		if err != nil {
			fmt.Println("BatchWork error:", err)
			continue
		}
		fmt.Println("已下载：", total)
		// json文件重命名
		jsonFileRename(jsonFile)
	}
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

// 读取json目录中的json文件
func readJsonDirFile() (jsonFileList []string) {
	fileInfoList, err := ioutil.ReadDir(config.ResponseJsonDir)
	if err != nil {
		panic(err)
	}
	for _, info := range fileInfoList {
		if strings.ToLower(path.Ext(info.Name())) == ".json" {
			jsonFileList = append(jsonFileList, config.ResponseJsonDir+string(os.PathSeparator)+info.Name())
		}
	}
	return
}

// json文件重命名，防止后续再次被读
func jsonFileRename(jsonFile string) {
	_ = os.Rename(jsonFile, jsonFile+".over")
}
