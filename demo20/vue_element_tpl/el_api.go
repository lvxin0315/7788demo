package vue_element_tpl

import (
	"github.com/lvxin0315/7788demo/demo20/tools"
	"strings"
)

const ELApiTpl = `
import request from '@/utils/request'

export function fetch$moduleNameList(query) {
  return request({
    url: '/$routerName/list',
    method: 'get',
    params: query
  })
}

export function fetch$moduleName(id) {
  return request({
    url: '/$routerName/detail',
    method: 'get',
    params: { id }
  })
}

export function create$moduleName(data) {
  return request({
    url: '/$routerName/create',
    method: 'post',
    data
  })
}

export function update$moduleName(data) {
  return request({
    url: '/$routerName/update',
    method: 'post',
    data
  })
}

export function delete$moduleName(id) {
  return request({
    url: '/$routerName/delete',
    method: 'post',
    id
  })
}

`

/**
 * @Author lvxin0315@163.com
 * @Description 请求api
 * @Date 11:25 上午 2021/5/7
 * @Param
 * @return
 **/
type ELApi struct {
	ModuleName    string
	camelCaseName string
	upperName     string
}

func (el *ELApi) ToString() string {
	// 首字母大写
	el.upperName = tools.CamelString(tools.Capitalize(el.ModuleName))

	tpl := ELApiTpl
	tpl = strings.ReplaceAll(tpl, "$moduleName", el.upperName)
	tpl = strings.ReplaceAll(tpl, "$routerName", el.ModuleName)
	return tpl
}
