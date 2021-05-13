package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lvxin0315/7788demo/demo20/model"
	"github.com/lvxin0315/7788demo/demo20/vue_element_tpl"
	"io/ioutil"
)

var elTemplate vue_element_tpl.ELTemplate

// 条件
var filterContainer vue_element_tpl.ELFilterContainer

// form
var dialogContainer vue_element_tpl.ELDialog

// table
var tableContainer vue_element_tpl.ELTable

func init() {
	elTemplate = vue_element_tpl.ELTemplate{}
	// query
	elTemplate.QueryDataList = make(map[string]string)
	// 默认操作
	elTemplate.CtrlContainer = new(vue_element_tpl.ELCtrlContainer)
	elTemplate.FilterContainer = &filterContainer
	elTemplate.DialogFormContainer = &dialogContainer
	elTemplate.TableContainer = &tableContainer
}

var mysqlDBTypeList = []string{
	"int", "decimal", "float", "text", "varchar",
}

func main() {
	// TODO 连库现在没什么用，最好是使用结构体及标签来处理
	gormDB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		"root",
		"root",
		"127.0.0.1",
		3306,
		"gg_study",
		"utf8mb4"))
	if err != nil {
		panic(err)
	}
	// 迁移
	gormDB.AutoMigrate(model.Course{})
	scope := gormDB.NewScope(model.Course{})
	// course 表处理
	elTemplate.ModuleName = "course"
	for _, f := range scope.Fields() {
		addElByType(f)
	}
	// 写js
	writeJsApiFile()
	// 写vue
	writeVueFile()

}

// 类型转数据结构
func _scope2ElData(tagTypeSetting string) string {
	fmt.Println("tagTypeSetting: ", tagTypeSetting)
	for _, dbType := range mysqlDBTypeList {
		if len(tagTypeSetting) < len(dbType) {
			continue
		}
		if tagTypeSetting[:len(dbType)] == dbType {
			return dbType
		}
	}
	return ""
}

// 根据类型添加el结构
func addElByType(scopeField *gorm.Field) {
	//fmt.Println(f.TagSettings)
	tagTypeSetting, ok := scopeField.TagSettingsGet("TYPE")
	if !ok {
		return
	}
	switch _scope2ElData(tagTypeSetting) {
	case "int", "decimal", "float", "varchar":
		// 条件
		filterContainer.WhereELList = append(filterContainer.WhereELList, &vue_element_tpl.WhereInput{
			Title: scopeField.Name,
			Field: scopeField.DBName,
		})
		elTemplate.QueryDataList[scopeField.DBName] = "undefined"
		// form
		dialogContainer.FormList = append(dialogContainer.FormList, &vue_element_tpl.DialogForm{
			Title:      scopeField.Name,
			Field:      scopeField.DBName,
			IsTextarea: false,
		})
		// 表格
		tableContainer.TableColumnList = append(tableContainer.TableColumnList, &vue_element_tpl.TableColumn{
			Title: scopeField.Name,
			Field: scopeField.DBName,
		})
	case "text":
		// form
		dialogContainer.FormList = append(dialogContainer.FormList, &vue_element_tpl.DialogForm{
			Title:      scopeField.Name,
			Field:      scopeField.DBName,
			IsTextarea: true,
		})
	}
}

func writeJsApiFile() {
	elApi := new(vue_element_tpl.ELApi)
	elApi.ModuleName = elTemplate.ModuleName
	ioutil.WriteFile(fmt.Sprintf("tmp/%s.js", elTemplate.ModuleName), []byte(elApi.ToString()), 0755)
}

func writeVueFile() {
	ioutil.WriteFile(fmt.Sprintf("tmp/%s.vue", elTemplate.ModuleName), []byte(elTemplate.ToString()), 0755)
}
