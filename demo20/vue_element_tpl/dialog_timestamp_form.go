package vue_element_tpl

import "strings"

const DialogTimestampFormTpl = `
			<el-form-item label="$title" prop="$field">
				<el-date-picker v-model="temp.$field" placeholder="Please input $title" type="datetime" />
			</el-form-item>
`

type DialogTimestampForm struct {
	Title string
	Field string
}

func (el *DialogTimestampForm) ToString() string {
	tpl := DialogTimestampFormTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	return tpl
}

func (el *DialogTimestampForm) IsDialogForm() bool {
	return true
}
