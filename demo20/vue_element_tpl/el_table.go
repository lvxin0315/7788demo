package vue_element_tpl

import "strings"

const (
	ColumnDefault = iota
	ColumnTimestamp
	ColumnLink
)

const ELTableTpl = `
	<el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      @sort-change="sortChange"
    >

		<el-table-column :label="$t('table.id')" prop="id" sortable="custom" align="center" :class-name="getSortClass('id')">
			<template slot-scope="{row}">
				<span>{{ row.id }}</span>
			</template>
		</el-table-column>

		$tableColumnELHtml

		<el-table-column :label="$t('table.actions')" align="center" width="230" class-name="small-padding fixed-width">
			<template slot-scope="{row,$index}">
				<el-button type="primary" size="mini" @click="handleUpdate(row)">
				{{ $t('table.edit') }}
				</el-button>
				<el-button v-if="row.status!='deleted'" size="mini" type="danger" @click="handleDelete(row,$index)">
				{{ $t('table.delete') }}
				</el-button>
			</template>
		</el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList"/>
`

/**
 * @Author lvxin0315@163.com
 * @Description 列表部分
 * @Date 10:27 上午 2021/5/7
 * @Param
 * @return
 **/
type ELTable struct {
	TableColumnList   []ELTableColumn
	tableColumnELHtml string
}

func (el *ELTable) ToString() string {
	el.tableColumnELHtml = ""
	for _, whereEL := range el.TableColumnList {
		el.tableColumnELHtml += whereEL.ToString()
	}
	tpl := ELTableTpl
	tpl = strings.ReplaceAll(tpl, "$tableColumnELHtml", el.tableColumnELHtml)
	return tpl
}
