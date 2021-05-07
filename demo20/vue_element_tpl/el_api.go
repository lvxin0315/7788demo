package vue_element_tpl

import (
	"github.com/lvxin0315/7788demo/demo20/tools"
	"strings"
)

const ELApiTpl = `
import request from '@/utils/request'

export function fetch$funcNameList(query) {
  return request({
    url: '/$routerName/list',
    method: 'get',
    params: query
  })
}

export function fetch$funcName(id) {
  return request({
    url: '/$routerName/detail',
    method: 'get',
    params: { id }
  })
}

export function create$funcName(data) {
  return request({
    url: '/$routerName/create',
    method: 'post',
    data
  })
}

export function update$funcName(data) {
  return request({
    url: '/$routerName/update',
    method: 'post',
    data
  })
}

export function delete$funcName(id) {
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
	ModuleName string
}

func (el *ELApi) ToString() string {
	// 首字母大写
	tpl := ELApiTpl
	tpl = strings.ReplaceAll(tpl, "$funcName", tools.Capitalize(el.ModuleName))
	tpl = strings.ReplaceAll(tpl, "$routerName", el.ModuleName)
	return tpl
}
