package vue_element_tpl

import "strings"

type WhereSelect struct {
	Title string
	Field string
}

const whereSelectTpl = `
	<el-select v-model="listQuery.$field" placeholder="$title" clearable class="$class">
		<el-option v-for="item in $fieldOptions" :key="item" :label="item" :value="item"/>
	</el-select>
`

func (el *WhereSelect) ToString() string {
	tpl := whereSelectTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	tpl = strings.ReplaceAll(tpl, "$class", defaultWhereClass)
	return tpl
}

func (el *WhereSelect) IsWhere() bool {
	return true
}
