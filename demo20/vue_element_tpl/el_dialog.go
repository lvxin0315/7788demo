package vue_element_tpl

import "strings"

const ELDialogTpl = `
		<el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
			<el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
				
				$formELHtml

			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button @click="dialogFormVisible = false">
				取消
				</el-button>
				<el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">
				确定
				</el-button>
			</div>
		</el-dialog>
`

/**
 * @Author lvxin0315@163.com
 * @Description 新增&编辑的弹框
 * @Date 11:59 上午 2021/5/7
 * @Param
 * @return
 **/
type ELDialog struct {
	FormList   []ELDialogForm
	formELHtml string
}

func (el ELDialog) ToString() string {
	el.formELHtml = ""
	for _, formEL := range el.FormList {
		el.formELHtml += formEL.ToString()
	}
	tpl := ELDialogTpl
	tpl = strings.ReplaceAll(tpl, "$formELHtml", el.formELHtml)
	return tpl
}
