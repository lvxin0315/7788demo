package main

import (
	"github.com/lvxin0315/7788demo/demo20/vue_element_tpl"
	"io/ioutil"
)

func main() {
	elFilterContainer := new(vue_element_tpl.ELFilterContainer)
	elFilterContainer.WhereELList = append(elFilterContainer.WhereELList, &vue_element_tpl.WhereInput{
		Title: "标题",
		Field: "title",
	})
	elFilterContainer.WhereELList = append(elFilterContainer.WhereELList, &vue_element_tpl.WhereSelect{
		Title: "分类",
		Field: "category",
	})
	elFilterContainer.WhereELList = append(elFilterContainer.WhereELList, &vue_element_tpl.WhereSelect{
		Title: "分类1",
		Field: "category1",
	})
	elFilterContainer.WhereELList = append(elFilterContainer.WhereELList, &vue_element_tpl.WhereSelect{
		Title: "分类2",
		Field: "category2",
	})
	elFilterContainer.WhereELList = append(elFilterContainer.WhereELList, &vue_element_tpl.WhereSelect{
		Title: "分类3",
		Field: "category3",
	})
	elFilterContainer.WhereELList = append(elFilterContainer.WhereELList, &vue_element_tpl.WhereSelect{
		Title: "分类4",
		Field: "category4",
	})
	elFilterContainer.WhereELList = append(elFilterContainer.WhereELList, &vue_element_tpl.WhereSelect{
		Title: "分类5",
		Field: "category5",
	})
	elFilterContainer.WhereELList = append(elFilterContainer.WhereELList, &vue_element_tpl.WhereSelect{
		Title: "分类6",
		Field: "category6",
	})
	ioutil.WriteFile("tmp/test_where.vue", []byte(elFilterContainer.ToString()), 0755)
}
