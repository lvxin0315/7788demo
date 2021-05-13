package vue_element_tpl

import "strings"

const DialogTinymceFormTpl = `
			<el-form-item label="$title" prop="$field">
				<Tinymce ref="editor" v-model="temp.$field" :height="400" />
			</el-form-item>
`

type DialogTinymceForm struct {
	Title string
	Field string
}

func (el *DialogTinymceForm) ToString() string {
	tpl := DialogTinymceFormTpl
	tpl = strings.ReplaceAll(tpl, "$field", el.Field)
	tpl = strings.ReplaceAll(tpl, "$title", el.Title)
	return tpl
}

func (el *DialogTinymceForm) IsDialogForm() bool {
	return true
}
