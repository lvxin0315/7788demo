package vue_element_tpl

import "strings"

const DialogSwitchFormTpl = `
			<el-form-item label="$title" prop="$field">
        		<el-switch v-model="temp.$field" />
      		</el-form-item>
`

type DialogSwitchForm struct {
	Title string
	Field string
}

func (el *DialogSwitchForm) ToString() string {
	tpl := DialogSwitchFormTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	return tpl
}

func (el *DialogSwitchForm) IsDialogForm() bool {
	return true
}
