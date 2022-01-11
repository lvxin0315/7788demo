package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/xuri/excelize/v2"
)

var excelKey = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

type excelData struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

var excelPath = "excel"

// 通过读取json内容，生成excel
func main() {
	// 初始化，创建excel文件夹
	err := os.MkdirAll(excelPath, 0755)
	if err != nil {
		panic(err)
	}
	// 读取当前目录
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fiList, err := ioutil.ReadDir(pwd)
	if err != nil {
		panic(err)
	}
	for _, info := range fiList {
		// 非json文件，过~
		if path.Ext(info.Name()) != ".json" {
			continue
		}
		createExcelFile(info.Name())
	}
}

// 创建文件
func createExcelFile(jsonFilePath string) {
	excelFilename, dataList := readJsonFileInfo(jsonFilePath)
	// 创建excel
	excelFile := excelize.NewFile()
	var defaultSheet = "Sheet1"
	// 创建一个工作表
	index := excelFile.NewSheet(defaultSheet)
	excelFile.SetActiveSheet(index)
	// 填充内容 TODO
	for i, v := range dataList {
		err := excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[0], i+1), v.Name)
		if err != nil {
			panic(err)
		}
		err = excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[1], i+1), v.Title)
		if err != nil {
			panic(err)
		}
	}

	// 生成文件
	if err := excelFile.SaveAs(excelPath + "/" + excelFilename); err != nil {
		fmt.Println(err)
	}
}

// 读取文件信息和内容
func readJsonFileInfo(jsonFilePath string) (excelFilename string, dataList []excelData) {
	content, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &dataList)
	if err != nil {
		panic(err)
	}
	fileNameWithSuffix := path.Base(jsonFilePath)
	excelFilename = strings.ReplaceAll(fileNameWithSuffix, ".json", ".xlsx")
	return
}
