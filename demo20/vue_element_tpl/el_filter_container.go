package vue_element_tpl

import "strings"

const ELFilterContainerTpl = `
<div class="filter-container">
	$whereELHtml

	<el-button v-waves class="filter-search-item" type="primary" icon="el-icon-search" @click="handleFilter">
        查询
    </el-button>
</div>
`

type ELFilterContainer struct {
	WhereELList []ELWhere
	// 存放HTML
	whereELHtml string
}

func (el *ELFilterContainer) ToString() string {
	el.whereELHtml = ""
	for _, whereEL := range el.WhereELList {
		el.whereELHtml += whereEL.ToString()
	}
	tpl := ELFilterContainerTpl
	tpl = strings.ReplaceAll(tpl, "$whereELHtml", el.whereELHtml)
	return tpl
}
