package model

type DetailResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id       string `json:"id"`
		Shopid   string `json:"shopid"`
		Mainimgs []struct {
			Imgurl string `json:"imgurl"`
		} `json:"mainimgs"`
		Categorys []interface{} `json:"categorys"`
		Picdetail []string      `json:"picdetail"`
		SkuImages []struct {
			Scalable int    `json:"scalable"`
			ImageUrl string `json:"imageUrl"`
		} `json:"skuImages"`
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Stock    struct {
			Desc      string `json:"desc"`
			Count     int    `json:"count"`
			MinNum    int    `json:"minNum"`
			MinQtyNum int    `json:"minQtyNum"`
		} `json:"stock"`
		Price struct {
			Value      int    `json:"value"`
			Flag       string `json:"flag"`
			Flagdesc   string `json:"flagdesc"`
			Showprice  int    `json:"showprice"`
			CanBuy     int    `json:"canBuy"`
			Spec       string `json:"spec"`
			Superprice int    `json:"superprice"`
			Tariff     int    `json:"tariff"`
		} `json:"price"`
		Spec []struct {
			Pid string `json:"pid"`
		} `json:"spec"`
		Place []struct {
			Prompt string `json:"prompt"`
			Value  string `json:"value"`
		} `json:"place"`
		Action    string `json:"action"`
		Promotion struct {
			Pendingcount int `json:"pendingcount"`
		} `json:"promotion"`
		Seller struct {
			Id     string `json:"id"`
			Title  string `json:"title"`
			Icon   string `json:"icon"`
			Action string `json:"action"`
		} `json:"seller"`
		SkuStatus struct {
			Status            int    `json:"status"`
			StatusDesc        string `json:"statusDesc"`
			ArrivalTimeDesc   string `json:"arrivalTimeDesc"`
			Color             string `json:"color"`
			Sellerservicedesc string `json:"sellerservicedesc"`
		} `json:"skuStatus"`
		ProductType            int           `json:"productType"`
		SmallImg               string        `json:"smallImg"`
		Expiration             int           `json:"expiration"`
		Restricts              []interface{} `json:"restricts"`
		Isdelivery             int           `json:"isdelivery"`
		Cityname               string        `json:"cityname"`
		Cityid                 int           `json:"cityid"`
		SecondSellercategory   string        `json:"secondSellercategory"`
		SecondSellercategoryVo struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"secondSellercategoryVo"`
		Balancerefund          int    `json:"balancerefund"`
		Commercialid           int    `json:"commercialid"`
		Status                 int    `json:"status"`
		Isjoinnewpriceactivity int    `json:"isjoinnewpriceactivity"`
		Sapcategoryid          string `json:"sapcategoryid"`
		Taglist                []struct {
			Type string `json:"type"`
			Text string `json:"text"`
			Sort int    `json:"sort"`
		} `json:"taglist"`
		Arrivalnotice int `json:"arrivalnotice"`
		Batchflag     int `json:"batchflag"`
		Skusaletype   int `json:"skusaletype"`
		Skuservicedes struct {
			Freightdetail []struct {
				Title    string `json:"title"`
				Subtitle string `json:"subtitle"`
				Type     int    `json:"type"`
			} `json:"freightdetail"`
			Servicename string `json:"servicename"`
		} `json:"skuservicedes"`
		Standarddescup []struct {
			Attributes string `json:"attributes"`
			Desc       string `json:"desc"`
		} `json:"standarddescup"`
		Skutype       int `json:"skutype"`
		RankingDetail struct {
			ActionUrl          string `json:"actionUrl"`
			Description        string `json:"description"`
			BackgroundImageUrl string `json:"backgroundImageUrl"`
		} `json:"rankingDetail"`
		FavoriteButton struct {
			State int `json:"state"`
		} `json:"favoriteButton"`
	} `json:"data"`
}
