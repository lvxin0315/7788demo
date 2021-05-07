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

/**
 * @Author lvxin0315@163.com
 * @Description 列表查询条件部分
 * @Date 10:26 上午 2021/5/7
 * @Param
 * @return
 **/
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
