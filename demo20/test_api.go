package main

import (
	"github.com/lvxin0315/7788demo/demo20/vue_element_tpl"
	"io/ioutil"
)

func main() {
	elApi := new(vue_element_tpl.ELApi)
	elApi.ModuleName = "demo_ta"
	ioutil.WriteFile("tmp/test_api.vue", []byte(elApi.ToString()), 0755)
}
