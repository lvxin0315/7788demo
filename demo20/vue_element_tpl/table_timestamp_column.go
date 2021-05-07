package vue_element_tpl

import "strings"

const TableTimestampColumnTpl = `
		<el-table-column label="$title" align="center">
			<template slot-scope="{row}">
				<span>{{ row.$field | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
			</template>
		</el-table-column>
`

type TableTimestampColumn struct {
	Title string
	Field string
}

func (el *TableTimestampColumn) ToString() string {
	tpl := TableTimestampColumnTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	return tpl
}

func (el *TableTimestampColumn) IsTableColumn() bool {
	return true
}
