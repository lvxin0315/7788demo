package main

import (
	"github.com/lvxin0315/7788demo/demo20/vue_element_tpl"
	"io/ioutil"
)

func main() {
	elDialog := new(vue_element_tpl.ELDialog)
	elDialog.FormList = append(elDialog.FormList, &vue_element_tpl.DialogForm{
		Title:      "标题",
		Field:      "title",
		IsTextarea: false,
	})
	elDialog.FormList = append(elDialog.FormList, &vue_element_tpl.DialogForm{
		Title:      "备注",
		Field:      "remark",
		IsTextarea: true,
	})
	elDialog.FormList = append(elDialog.FormList, &vue_element_tpl.DialogTimestampForm{
		Title: "时间",
		Field: "timestamp",
	})
	elDialog.FormList = append(elDialog.FormList, &vue_element_tpl.DialogSelectForm{
		Title: "类型",
		Field: "calendarType",
	})
	ioutil.WriteFile("tmp/test_dialog.vue", []byte(elDialog.ToString()), 0755)
}
