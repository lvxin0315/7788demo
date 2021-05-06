package vue_element_tpl

/**
 * @Author lvxin0315@163.com
 * @Description vue-element组建
 * @Date 4:37 下午 2021/5/6
 * @Param
 * @return
 **/
type EL interface {
	ToString() string
}

/**
 * @Author lvxin0315@163.com
 * @Description vue-element 列表查询条件组建
 * @Date 4:37 下午 2021/5/6
 * @Param
 * @return
 **/
type ELWhere interface {
	EL
	IsWhere() bool
}

// 默认查询组建的class
const defaultWhereClass = `filter-item`
