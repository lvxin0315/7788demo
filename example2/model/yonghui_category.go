package model

type AllCategoryResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Cityid     int                           `json:"cityid"`
		Sellerid   int                           `json:"sellerid"`
		Sellername string                        `json:"sellername"`
		Shopid     string                        `json:"shopid"`
		Categorys  []AllCategoryResponseCategory `json:"categorys"`
		Isdelivery int                           `json:"isdelivery"`
		Cityname   string                        `json:"cityname"`
	} `json:"data"`
	Now int64 `json:"now"`
}

type AllCategoryResponseCategory struct {
	Categoryid             string                           `json:"categoryid"`
	Categoryname           string                           `json:"categoryname"`
	Imgurl                 string                           `json:"imgurl"`
	Subcategory            []AllCategoryResponseSubCategory `json:"subcategory"`
	SaleCategoryTemplateId int                              `json:"saleCategoryTemplateId,omitempty"`
}

type AllCategoryResponseSubCategory struct {
	Categoryid   string `json:"categoryid"`
	Categoryname string `json:"categoryname"`
	Subcategory  []struct {
		Categoryid             string        `json:"categoryid"`
		Categoryname           string        `json:"categoryname"`
		Subcategory            []interface{} `json:"subcategory"`
		SaleCategoryTemplateId int           `json:"saleCategoryTemplateId,omitempty"`
	} `json:"subcategory"`
	PreSaleFlag            int `json:"preSaleFlag"`
	SaleCategoryTemplateId int `json:"saleCategoryTemplateId,omitempty"`
}
