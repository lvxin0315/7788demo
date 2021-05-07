package vue_element_tpl

import "strings"

const TableLinkColumnTpl = `
		<el-table-column label="$title" align="center">
			<template slot-scope="{row}">
				<span class="link-type" @click="handleUpdate(row)">{{ row.$field }}</span>
			</template>
		</el-table-column>
`

type TableLinkColumn struct {
	Title string
	Field string
}

func (el *TableLinkColumn) ToString() string {
	tpl := TableLinkColumnTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	return tpl
}

func (el *TableLinkColumn) IsTableColumn() bool {
	return true
}
