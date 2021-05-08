package main

import (
	"fmt"
	"github.com/lvxin0315/7788demo/demo20/vue_element_tpl"
	"io/ioutil"
)

type elData struct {
	Title string
	Field string
}

type elModule struct {
	Name     string
	dataList []elData
}

func main() {
	m := elModule{
		Name: "goods",
		dataList: []elData{
			{
				Title: "名称",
				Field: "name",
			},
			{
				Title: "标题",
				Field: "title",
			},
			{
				Title: "开始时间",
				Field: "begin_time",
			},
			{
				Title: "描述",
				Field: "des",
			},
		},
	}

	elTemplate := new(vue_element_tpl.ELTemplate)
	elTemplate.ModuleName = m.Name
	// 条件
	filterContainer := new(vue_element_tpl.ELFilterContainer)
	// form
	dialogContainer := new(vue_element_tpl.ELDialog)
	// table
	tableContainer := new(vue_element_tpl.ELTable)

	elTemplate.FilterContainer = filterContainer
	elTemplate.CtrlContainer = new(vue_element_tpl.ELCtrlContainer)
	elTemplate.DialogFormContainer = dialogContainer
	elTemplate.TableContainer = tableContainer

	// 1.生成api js文件
	elApi := new(vue_element_tpl.ELApi)
	elApi.ModuleName = m.Name
	ioutil.WriteFile(fmt.Sprintf("tmp/%s.js", m.Name), []byte(elApi.ToString()), 0755)

	// 2.开始填充模板内容
	filterContainer.WhereELList = append(filterContainer.WhereELList, &vue_element_tpl.WhereInput{
		Title: m.dataList[0].Title,
		Field: m.dataList[0].Field,
	}, &vue_element_tpl.WhereInput{
		Title: m.dataList[1].Title,
		Field: m.dataList[1].Field,
	}, &vue_element_tpl.WhereInput{
		Title: m.dataList[2].Title,
		Field: m.dataList[2].Field,
	}, &vue_element_tpl.WhereInput{
		Title: m.dataList[3].Title,
		Field: m.dataList[3].Field,
	})

	dialogContainer.FormList = append(dialogContainer.FormList, &vue_element_tpl.DialogForm{
		Title:      m.dataList[0].Title,
		Field:      m.dataList[0].Field,
		IsTextarea: false,
	}, &vue_element_tpl.DialogForm{
		Title:      m.dataList[1].Title,
		Field:      m.dataList[1].Field,
		IsTextarea: false,
	}, &vue_element_tpl.DialogTimestampForm{
		Title: m.dataList[2].Title,
		Field: m.dataList[2].Field,
	}, &vue_element_tpl.DialogForm{
		Title:      m.dataList[3].Title,
		Field:      m.dataList[3].Field,
		IsTextarea: true,
	})

	tableContainer.TableColumnList = append(tableContainer.TableColumnList, &vue_element_tpl.TableColumn{
		Title: m.dataList[0].Title,
		Field: m.dataList[0].Field,
	}, &vue_element_tpl.TableLinkColumn{
		Title: m.dataList[1].Title,
		Field: m.dataList[1].Field,
	}, &vue_element_tpl.TableTimestampColumn{
		Title: m.dataList[2].Title,
		Field: m.dataList[2].Field,
	})

	// query
	elTemplate.QueryDataList = map[string]string{
		"name":       "undefined",
		"title":      "undefined",
		"begin_time": "new Date()",
		"des":        "undefined",
	}

	// 保存文件
	ioutil.WriteFile(fmt.Sprintf("tmp/%s.vue", m.Name), []byte(elTemplate.ToString()), 0755)

}
