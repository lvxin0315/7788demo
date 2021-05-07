package vue_element_tpl

const ELCtrlContainerTpl = `
	<div class="ctrl-container">
		<el-button class="ctrl-item" type="primary" icon="el-icon-edit" @click="handleCreate">
			新增
		</el-button>
		<el-button v-waves :loading="downloadLoading" class="ctrl-item" type="primary" icon="el-icon-download" @click="handleDownload">
			导出
		</el-button>
	</div>
`

type ELCtrlContainer struct {
}

func (el *ELCtrlContainer) ToString() string {
	return ELCtrlContainerTpl
}
