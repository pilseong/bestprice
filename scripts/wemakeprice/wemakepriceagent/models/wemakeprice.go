package models

import "time"

type WemakepriceBEST struct {
	Data struct {
		Deals []Item `json:"deals"`
	} `json:"data"`
}

type Item struct {
	LinkInfo       int     `json:"linkInfo"`
	LinkType       string  `json:"linkType"`
	DispNm         string  `json:"dispNm"`
	MediumImgURL   string  `json:"mediumImgUrl"`
	OriginPrice    int     `json:"originPrice"`
	SalePrice      int     `json:"salePrice"`
	DiscountPrice  int     `json:"discountPrice"`
	DiscountType   string  `json:"discountType"`
	DiscountText   string  `json:"discountText"`
	DiscountRate   int     `json:"discountRate"`
	SalesCount     int     `json:"salesCount"`
	ShipText       string  `json:"shipText"`
	AutoLabel1     string  `json:"autoLabel1"`
	ReviewCount    int     `json:"reviewCount"`
	ReviewAvgPoint float64 `json:"reviewAvgPoint"`
}

type Batch struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"createdat"`
}

type Latest struct {
	ID      int64 `json:"id"`
	BatchId int64 `json:"batchid"`
}
