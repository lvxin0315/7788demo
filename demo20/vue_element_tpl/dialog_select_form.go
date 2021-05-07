package vue_element_tpl

import "strings"

const DialogSelectFormTpl = `
			<el-form-item label="$title" prop="$field">
				<el-select v-model="temp.$field" placeholder="Please select $title" class="filter-item">
					<el-option v-for="item in $fieldOptions" :key="item.key" :label="item.display_name" :value="item.key"/>
				</el-select>
			</el-form-item>
`

type DialogSelectForm struct {
	Title string
	Field string
}

func (el *DialogSelectForm) ToString() string {
	tpl := DialogSelectFormTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	return tpl
}

func (el *DialogSelectForm) IsDialogForm() bool {
	return true
}
