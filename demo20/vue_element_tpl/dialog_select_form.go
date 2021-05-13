package vue_element_tpl

import (
	"fmt"
	"strings"
)

const DialogSelectFormTpl = `
			<el-form-item label="$title" prop="$field">
				<el-select v-model="temp.$field" placeholder="Please select $title" class="filter-item" $multiple >
					<el-option v-for="item in $fieldOptions" :key="item" :label="item" :value="item"/>
				</el-select>
			</el-form-item>

`

type DialogSelectForm struct {
	Title      string
	Field      string
	IsMultiple bool
	// TODO 应该使用 interface{}, 之后应该单独封装
	FieldOptions []string
}

func (el *DialogSelectForm) ToString() string {
	tpl := DialogSelectFormTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	// 是否多选
	if el.IsMultiple {
		tpl = strings.ReplaceAll(tpl, "$multiple", "multiple")
	} else {
		tpl = strings.ReplaceAll(tpl, "$multiple", "")
	}
	// 选项
	tpl = strings.ReplaceAll(tpl, "$fieldOptions", fmt.Sprintf("['%s']", strings.Join(el.FieldOptions, "','")))
	return tpl
}

func (el *DialogSelectForm) IsDialogForm() bool {
	return true
}
