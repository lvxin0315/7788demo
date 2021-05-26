package model

type IndexData struct {
	HeadTitle    string `json:"head_title"`    // 首页主标题
	HeadDescribe string `json:"head_describe"` // 首页主描述

	CourseIntroduceTitle    string `json:"course_introduce_title"`    // 课程介绍标题
	CourseIntroduceDescribe string `json:"course_introduce_describe"` // 课程介绍描述

	SubjectIntroduceTitle         string   `json:"subject_introduce_title"`          // 学科内容介绍标题
	SubjectIntroduceDescribe      string   `json:"subject_introduce_describe"`       // 学科内容介绍描述
	SubjectIntroduceLeftCategory  []string `json:"subject_introduce_left_category"`  // 左侧学科内容介绍分类
	SubjectIntroduceRightCategory []string `json:"subject_introduce_right_category"` // 右侧学科内容介绍分类

	TeachingMethodsTitle    string                `json:"teaching_methods_title"`    // 授课方式标题
	TeachingMethodsDescribe string                `json:"teaching_methods_describe"` // 授课方式描述
	TeachingMethodsCards    []TeachingMethodsCard `json:"teaching_methods_cards"`    // 授课方式卡片内容

	MemberIntroduceTitle    string `json:"member_introduce_title"`    // 成员介绍标题
	MemberIntroduceDescribe string `json:"member_introduce_describe"` // 成员介绍描述
}

type TeachingMethodsCard struct {
	TeachingMethodsCardTitle    string `json:"teaching_methods_card_title"`
	TeachingMethodsCardDescribe string `json:"teaching_methods_card_describe"`
}
