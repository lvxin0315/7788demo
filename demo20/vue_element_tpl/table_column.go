package vue_element_tpl

import "strings"

const TableColumnTpl = `
		<el-table-column label="$title" align="center">
			<template slot-scope="{row}">
				<span>{{ row.$field }}</span>
			</template>
		</el-table-column>
`

type TableColumn struct {
	Title string
	Field string
}

func (el *TableColumn) ToString() string {
	tpl := TableColumnTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	return tpl
}

func (el *TableColumn) IsTableColumn() bool {
	return true
}
