package vue_element_tpl

import "strings"

const DialogCheckBoxFormTpl = `
			<el-form-item label="$title" prop="$field">
				<el-checkbox-group v-model="temp.$field">
					<el-checkbox label="Option1" name="type" />
					<el-checkbox label="Option2" name="type" />
					<el-checkbox label="Option3" name="type" />
					<el-checkbox label="Option4" name="type" />
				</el-checkbox-group>
			</el-form-item>
`

type DialogCheckBoxForm struct {
	Title string
	Field string
}

func (el *DialogCheckBoxForm) ToString() string {
	tpl := DialogCheckBoxFormTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	return tpl
}

func (el *DialogCheckBoxForm) IsDialogForm() bool {
	return true
}
