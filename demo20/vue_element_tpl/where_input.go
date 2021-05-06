package vue_element_tpl

import "strings"

type WhereInput struct {
	Title string
	Field string
}

const whereInputTpl = `
	<el-input v-model="listQuery.$field" placeholder="$title" class="$class" @keyup.enter.native="handleFilter"/>
`

func (el *WhereInput) ToString() string {
	tpl := whereInputTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	tpl = strings.ReplaceAll(tpl, "$class", defaultWhereClass)
	return tpl
}

func (el *WhereInput) IsWhere() bool {
	return true
}
