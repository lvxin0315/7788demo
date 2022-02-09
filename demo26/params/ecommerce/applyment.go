package ecommerce

// Ecommerce 二级商户进件
type Ecommerce struct {
	OutRequestNo         string               `json:"out_request_no" mapstructure:"out_request_no"`                 // 业务申请编号
	OrganizationType     string               `json:"organization_type" mapstructure:"organization_type"`           // 主体类型
	BusinessLicenseInfo  BusinessLicenseInfo  `json:"business_license_info" mapstructure:"business_license_info"`   // 营业执照/登记证书信息
	OrganizationCertInfo OrganizationCertInfo `json:"organization_cert_info" mapstructure:"organization_cert_info"` // 组织机构代码证信息
	IdDocType            string               `json:"id_doc_type" mapstructure:"id_doc_type"`                       // 经营者/法人证件类型
	IdCardInfo           IdCardInfo           `json:"id_card_info" mapstructure:"id_card_info"`                     // 经营者/法人身份证信息
	IdDocInfo            IdDocInfo            `json:"id_doc_info" mapstructure:"id_doc_info"`                       // 经营者/法人其他类型证件信息
	NeedAccountInfo      bool                 `json:"need_account_info" mapstructure:"need_account_info"`           // 是否填写结算银行账户
	AccountInfo          AccountInfo          `json:"account_info" mapstructure:"account_info"`                     // 结算银行账户
	ContactInfo          ContactInfo          `json:"contact_info" mapstructure:"contact_info"`                     // 超级管理员信息
	SalesSceneInfo       SalesSceneInfo       `json:"sales_scene_info" mapstructure:"sales_scene_info"`             // 店铺信息
	SettlementInfo       SettlementInfo       `json:"settlement_info" mapstructure:"settlement_info"`               // 结算规则
	MerchantShortname    string               `json:"merchant_shortname" mapstructure:"merchant_shortname"`         // 商户简称
	Qualifications       string               `json:"qualifications" mapstructure:"qualifications"`                 // 特殊资质
	BusinessAdditionPics string               `json:"business_addition_pics" mapstructure:"business_addition_pics"` // 补充材料
	BusinessAdditionDesc string               `json:"business_addition_desc" mapstructure:"business_addition_desc"` // 补充说明
}

// BusinessLicenseInfo 营业执照/登记证书信息
type BusinessLicenseInfo struct {
	BusinessLicenseCopy   string `json:"business_license_copy" mapstructure:"business_license_copy"`     // 证件扫描件
	BusinessLicenseNumber string `json:"business_license_number" mapstructure:"business_license_number"` // 证件注册号
	MerchantName          string `json:"merchant_name" mapstructure:"merchant_name"`                     // 商户名称
	LegalPerson           string `json:"legal_person" mapstructure:"legal_person"`                       // 经营者/法定代表人姓名
	CompanyAddress        string `json:"company_address" mapstructure:"company_address"`                 // 注册地址
	BusinessTime          string `json:"business_time" mapstructure:"business_time"`                     // 营业期限
}

// OrganizationCertInfo 组织机构代码证信息
type OrganizationCertInfo struct {
	OrganizationCopy   string `json:"organization_copy" mapstructure:"organization_copy"`     // 组织机构代码证照片
	OrganizationNumber string `json:"organization_number" mapstructure:"organization_number"` // 组织机构代码
	OrganizationTime   string `json:"organization_time" mapstructure:"organization_time"`     // 组织机构代码有效期限
}

// IdCardInfo 经营者/法人身份证信息
type IdCardInfo struct {
	IdCardCopy      string     `json:"id_card_copy" mapstructure:"id_card_copy"`             // 身份证人像面照片
	IdCardNational  string     `json:"id_card_national" mapstructure:"id_card_national"`     // 身份证国徽面照片
	IdCardName      CipherText `json:"id_card_name" mapstructure:"id_card_name"`             // 身份证姓名
	IdCardNumber    CipherText `json:"id_card_number" mapstructure:"id_card_number"`         // 身份证号码
	IdCardValidTime string     `json:"id_card_valid_time" mapstructure:"id_card_valid_time"` // 身份证有效期限
}

// IdDocInfo 经营者/法人其他类型证件信息
type IdDocInfo struct {
	IdDocName    CipherText `json:"id_doc_name" mapstructure:"id_doc_name"`       // 证件姓名
	IdDocNumber  CipherText `json:"id_doc_number" mapstructure:"id_doc_number"`   // 证件号码
	IdDocCopy    string     `json:"id_doc_copy" mapstructure:"id_doc_copy"`       // 证件照片
	DocPeriodEnd string     `json:"doc_period_end" mapstructure:"doc_period_end"` // 证件结束日期
}

// AccountInfo 结算银行账户
type AccountInfo struct {
	BankAccountType string     `json:"bank_account_type" mapstructure:"bank_account_type"` // 74-对公账户、75-对私账户
	AccountBank     string     `json:"account_bank" mapstructure:"account_bank"`           // 开户银行
	AccountName     CipherText `json:"account_name" mapstructure:"account_name"`           // 开户名称
	BankAddressCode string     `json:"bank_address_code" mapstructure:"bank_address_code"` // 开户银行省市编码
	BankBranchId    string     `json:"bank_branch_id" mapstructure:"bank_branch_id"`       // 开户银行联行号
	BankName        string     `json:"bank_name" mapstructure:"bank_name"`                 // 开户银行全称 （含支行）
	AccountNumber   CipherText `json:"account_number" mapstructure:"account_number"`       // 银行账号
}

// ContactInfo 超级管理员信息
type ContactInfo struct {
	ContactType         string     `json:"contact_type" mapstructure:"contact_type"`                     // 超级管理员类型
	ContactName         CipherText `json:"contact_name" mapstructure:"contact_name"`                     // 超级管理员姓名
	ContactIdCardNumber CipherText `json:"contact_id_card_number" mapstructure:"contact_id_card_number"` // 超级管理员身份证件号码
	MobilePhone         CipherText `json:"mobile_phone" mapstructure:"mobile_phone"`                     // 超级管理员手机
	ContactEmail        CipherText `json:"contact_email" mapstructure:"contact_email"`                   // 超级管理员邮箱
}

// SalesSceneInfo 店铺信息
type SalesSceneInfo struct {
	StoreName           string `json:"store_name" mapstructure:"store_name"`                         // 店铺名称
	StoreUrl            string `json:"store_url" mapstructure:"store_url"`                           // 店铺链接
	StoreQrCode         string `json:"store_qr_code" mapstructure:"store_qr_code"`                   // 店铺二维码
	MiniProgramSubAppid string `json:"mini_program_sub_appid" mapstructure:"mini_program_sub_appid"` // 小程序AppID
}

// SettlementInfo 结算规则
type SettlementInfo struct {
	SettlementId      int    `json:"settlement_id" mapstructure:"settlement_id" `           // 结算规则ID
	QualificationType string `json:"qualification_type" mapstructure:"qualification_type" ` // 所属行业
}
