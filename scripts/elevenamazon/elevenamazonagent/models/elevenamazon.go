package models

type ElmazonBest struct {
	Data []struct {
		BlockList []struct {
			List []Item `json:"list"`
		} `json:"blockList"`
	} `json:"data"`
}

type Item struct {
	DeliveryInfos [][]struct {
		Text string `json:"text,omitempty"`
	} `json:"deliveryInfos"`
	PrdNo         string `json:"prdNo"`
	PrdNm         string `json:"prdNm"`
	ImageURL      string `json:"imageUrl"`
	FinalDscPrice string `json:"finalDscPrice"`
	LinkURL       string `json:"linkUrl"`
	LogData       struct {
		DataBody struct {
			ContentType       string `json:"content_type"`
			LastDiscountPrice string `json:"last_discount_price"`
			RankNo            string `json:"rank_no"`
		} `json:"dataBody"`
	} `json:"logData"`
}
