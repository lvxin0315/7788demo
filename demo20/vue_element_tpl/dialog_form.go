package vue_element_tpl

import "strings"

const DialogFormTpl = `
			<el-form-item label="$title" prop="$field">
				<el-input v-model="temp.$field" placeholder="Please input $title" $isTextarea />
			</el-form-item>
`

type DialogForm struct {
	Title      string
	Field      string
	IsTextarea bool
}

func (el *DialogForm) ToString() string {
	tpl := DialogFormTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	if el.IsTextarea {
		tpl = strings.ReplaceAll(tpl, "$isTextarea", `type="textarea"`)
	} else {
		tpl = strings.ReplaceAll(tpl, "$isTextarea", ``)
	}
	return tpl
}

func (el *DialogForm) IsDialogForm() bool {
	return true
}
