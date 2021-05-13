package vue_element_tpl

import "strings"

const DialogRadioFormTpl = `
			<el-form-item label="$title" prop="$field">
				<el-radio-group v-model="temp.$field">
					<el-radio label="Radio1" />
					<el-radio label="Radio2" />
				</el-radio-group>
			</el-form-item>
`

type DialogRadioForm struct {
	Title string
	Field string
}

func (el *DialogRadioForm) ToString() string {
	tpl := DialogRadioFormTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	return tpl
}

func (el *DialogRadioForm) IsDialogForm() bool {
	return true
}
