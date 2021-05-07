package main

import (
	"github.com/lvxin0315/7788demo/demo20/vue_element_tpl"
	"io/ioutil"
)

func main() {
	elTable := new(vue_element_tpl.ELTable)
	elTable.TableColumnList = append(elTable.TableColumnList, &vue_element_tpl.TableLinkColumn{
		Title: "标题",
		Field: "title",
	})

	elTable.TableColumnList = append(elTable.TableColumnList, &vue_element_tpl.TableColumn{
		Title: "标题1",
		Field: "title1",
	})

	elTable.TableColumnList = append(elTable.TableColumnList, &vue_element_tpl.TableColumn{
		Title: "标题2",
		Field: "title2",
	})

	elTable.TableColumnList = append(elTable.TableColumnList, &vue_element_tpl.TableLinkColumn{
		Title: "标题link",
		Field: "title_link",
	})

	elTable.TableColumnList = append(elTable.TableColumnList, &vue_element_tpl.TableTimestampColumn{
		Title: "时间",
		Field: "timestamp",
	})

	ioutil.WriteFile("tmp/test_table.vue", []byte(elTable.ToString()), 0755)
}
