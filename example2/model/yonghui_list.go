package model

type CategoryResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TID      string `json:"TID"`
		Activity struct {
		} `json:"activity"`
		Artificialrecommend struct {
		} `json:"artificialrecommend"`
		AsId       string        `json:"asId"`
		BrandNames []string      `json:"brandNames"`
		Categories []interface{} `json:"categories"`
		Coupon     struct {
		} `json:"coupon"`
		Index                 int           `json:"index"`
		Nostockskus           []interface{} `json:"nostockskus"`
		OutOfStockGoodsCounts int           `json:"outOfStockGoodsCounts"`
		Page                  int           `json:"page"`
		Pagecount             int           `json:"pagecount"`
		Productrecommend      []interface{} `json:"productrecommend"`
		RecommendChoose       int           `json:"recommendChoose"`
		RequestId             string        `json:"requestId"`
		Searchresulttype      int           `json:"searchresulttype"`
		Skus                  []YonghuiSku  `json:"skus"`
		Total                 int           `json:"total"`
		Totalpage             int           `json:"totalpage"`
		TraceId               string        `json:"traceId"`
	} `json:"data"`
}

type YonghuiSku struct {
	Action string `json:"action"`
	Batch  struct {
		BatchFlag int `json:"batchFlag"`
	} `json:"batch"`
	Cover struct {
		ImageUrl string `json:"imageUrl"`
	} `json:"cover"`
	InStock    int  `json:"inStock"`
	IsOnSale   bool `json:"isOnSale"`
	Preprocess int  `json:"preprocess"`
	Price      struct {
		Price       string `json:"price"`
		MarketPrice string `json:"marketPrice,omitempty"`
	} `json:"price"`
	RecSlogan   string `json:"recSlogan,omitempty"`
	SkuCode     string `json:"skuCode"`
	SkuSaleType int    `json:"skuSaleType"`
	SkuType     int    `json:"skuType"`
	Spu         struct {
		IsSpu int `json:"isSpu"`
	} `json:"spu"`
	SubTitle string `json:"subTitle,omitempty"`
	Tag      struct {
		RibbonTags []struct {
			Text     string `json:"text"`
			Type     int    `json:"type"`
			ImageUrl string `json:"imageUrl"`
		} `json:"ribbonTags,omitempty"`
		CommonTags []struct {
			Text string `json:"text"`
			Type int    `json:"type"`
		} `json:"commonTags,omitempty"`
	} `json:"tag,omitempty"`
	Title    string `json:"title"`
	Tracking struct {
		PriceInCent       int `json:"priceInCent"`
		MarketPriceInCent int `json:"marketPriceInCent"`
	} `json:"tracking"`
	Type       int `json:"type"`
	CartAction struct {
		ActionType int    `json:"actionType"`
		ActionText string `json:"actionText"`
	} `json:"cartAction,omitempty"`
}
