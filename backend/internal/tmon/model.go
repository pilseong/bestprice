package tmon

import "time"

type Item struct {
	ID                int64  `json:"id"`
	BatchID           int64  `json:"batchid"`
	Rank              int    `json:"rank"`
	StartDate         string `json:"startDate"`
	EndDate           string `json:"endDate"`
	Title             string `json:"title"`
	DealNo            int64  `json:"dealNo"`
	DeliveryFeePolicy string `json:"deliveryFeePolicy"`
	DeliveryFee       int    `json:"deliveryFee"`
	CategoryName      string `json:"categoryName"`
	URL               string `json:"url"`
	Pc3ColImageURL    string `json:"pc3ColImageUrl"`
	OriginalPrice     int    `json:"originalPrice"`
	Price             int    `json:"price"`
	DiscountRate      int    `json:"discountRate"`
}

type Batch struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"createdat"`
}

type Latest struct {
	ID      int64 `json:"id"`
	BatchId int64 `json:"batchid"`
}
