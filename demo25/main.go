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

type excelOption struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type excelData struct {
	Title       string        `json:"title"`
	Explain     string        `json:"explain"`
	OptionsJson []excelOption `json:"options_json"`
	Answer      string        `json:"answer"`
	KindText    string        `json:"kind_text"`
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
	// 填充表头
	excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[0], 1), "类型")
	excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[1], 1), "题干")
	excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[2], 1), "答案")
	excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[3], 1), "选项A")
	excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[4], 1), "选项B")
	excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[5], 1), "选项C")
	excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[6], 1), "选项D")
	excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[7], 1), "选项E")
	excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[8], 1), "选项F")
	excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[9], 1), "解析")
	// 填充内容
	for i, v := range dataList {
		col := i + 2
		// 类型
		err := excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[0], col), v.KindText)
		if err != nil {
			panic(err)
		}
		// 题干
		err = excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[1], col), v.Title)
		if err != nil {
			panic(err)
		}
		// 答案
		err = excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[2], col), v.Answer)
		if err != nil {
			panic(err)
		}
		// 选项
		optionIndex := 3
		for _, option := range v.OptionsJson { // A B C D,E,F 3,4,5,6,7,8
			err = excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[optionIndex], col), option.Value)
			if err != nil {
				panic(err)
			}
			optionIndex++
		}
		// 解析
		err = excelFile.SetCellValue(defaultSheet, fmt.Sprintf("%s%d", excelKey[9], col), v.Explain)
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
