package vue_element_tpl

import (
	"fmt"
	"github.com/lvxin0315/7788demo/demo20/tools"
	"strings"
)

const ELTemplateTpl = `
<template>
  <div class="app-container">

    $filterContainer
    
    $ctrlContainer

    $tableContainer

    $dialogFormContainer

  </div>
</template>

<script>
import { fetch$funcNameList, fetch$funcName, create$funcName, update$funcName, delete$funcName } from '@/api/$moduleName'
import waves from '@/directive/waves'
import {parseTime} from '@/utils'
import Pagination from '@/components/Pagination'

export default {
  name: 'ComplexTable',
  components: {Pagination},
  directives: {waves},
  filters: {
   
  },
  data() {
    return {
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20,
        $queryHtml
        sort: '+id'
      },

      $optionsHtml

      sortOptions: [{label: 'ID Ascending', key: '+id'}, {label: 'ID Descending', key: '-id'}],

      showReviewer: false,
      temp: {
        id: undefined,
        $queryHtml
      },

      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      dialogPvVisible: false,
 
      rules: {

      },
      downloadLoading: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      fetch$funcNameList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total

        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
   
    sortChange(data) {
      const {prop, order} = data
      if (prop === 'id') {
        this.sortByID(order)
      }
    },
    sortByID(order) {
      if (order === 'ascending') {
        this.listQuery.sort = '+id'
      } else {
        this.listQuery.sort = '-id'
      }
      this.handleFilter()
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        $queryHtml
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          create$funcName(this.temp).then(() => {
            this.list.unshift(this.temp)
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row) 
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          update$funcName(tempData).then(() => {
            const index = this.list.findIndex(v => v.id === this.temp.id)
            this.list.splice(index, 1, this.temp)
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleDelete(row, index) {
      this.$notify({
        title: '成功',
        message: '删除成功',
        type: 'success',
        duration: 2000
      })
      this.list.splice(index, 1)
    },
    getSortClass: function (key) {
      const sort = this.listQuery.sort
      return sort === '+${key}' ? 'ascending' : 'descending'
    }
  }
}
</script>

`

/**
 * @Author lvxin0315@163.com
 * @Description 整套模板
 * @Date 2:23 下午 2021/5/7
 * @Param
 * @return
 **/
type ELTemplate struct {
	ModuleName string
	// 条件字段 字段名:默认值
	QueryDataList       map[string]string
	FilterContainer     EL
	CtrlContainer       EL
	TableContainer      EL
	DialogFormContainer EL
	// api 相关
	ELApiContainer ELApi

	// options 变量
	OptionList []string
}

func (el *ELTemplate) ToString() string {
	tpl := ELTemplateTpl
	tpl = strings.ReplaceAll(tpl, "$filterContainer", el.FilterContainer.ToString())
	tpl = strings.ReplaceAll(tpl, "$ctrlContainer", el.CtrlContainer.ToString())
	tpl = strings.ReplaceAll(tpl, "$tableContainer", el.TableContainer.ToString())
	tpl = strings.ReplaceAll(tpl, "$dialogFormContainer", el.DialogFormContainer.ToString())

	//处理api相关
	tpl = strings.ReplaceAll(tpl, "$funcName", tools.Capitalize(el.ModuleName))
	tpl = strings.ReplaceAll(tpl, "$moduleName", el.ModuleName)
	tpl = strings.ReplaceAll(tpl, "$queryHtml", el._toQueryHtml())
	tpl = strings.ReplaceAll(tpl, "$optionsHtml", el._toQueryHtml())

	return tpl
}

/**
 * @Author lvxin0315@163.com
 * @Description 条件处理
 * @Date 3:49 下午 2021/5/7
 * @Param
 * @return
 **/
func (el *ELTemplate) _toQueryHtml() string {
	queryHtml := ""
	for k, v := range el.QueryDataList {
		queryHtml += fmt.Sprintf(`%s: %s,
`, k, v)
	}
	return queryHtml
}

/**
 * @Author lvxin0315@163.com
 * @Description option变量
 * @Date 5:37 下午 2021/5/7
 * @Param
 * @return
 **/
func (el *ELTemplate) _toOptionsHtml() string {
	optionsHtml := ""
	for _, o := range el.OptionList {
		optionsHtml += fmt.Sprintf(`%sOptions: [1, 2, 3],
`, o)
	}
	return optionsHtml
}
