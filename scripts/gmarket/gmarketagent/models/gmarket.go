package models

import "time"

// User is the structure which holds one user from the database.
type Item struct {
	ID                  int64  `json:"id"`
	BatchID             int64  `json:"batchid"`
	Rank                int    `json:"rank"`
	GoodsCode           string `json:"goodsCode,omitempty"`
	GoodsName           string `json:"goodsName,omitempty"`
	LinkURL             string `json:"linkUrl,omitempty"`
	DiscountPrice       string `json:"discountPrice,omitempty"`
	Price               string `json:"price,omitempty"`
	DeliveryInfo        string `json:"deliveryInfo,omitempty"`
	ItemPriceType       string `json:"itemPriceType,omitempty"`
	ImageURL            string `json:"imageUrl,omitempty"`
	DiscountRate        int    `json:"discountRate,omitempty"`
	ExpressShippingText string `json:"expressShippingText,omitempty"`
	DeliveryText        string `json:"deliveryText,omitempty"`
	ConsultingPeriod    string `json:"consultingPeriod,omitempty"`
	IsPriceExpose       bool   `json:"isPriceExpose,omitempty"`
	IsStartPrice        bool   `json:"isStartPrice,omitempty"`
	IsFreeShipping      bool   `json:"isFreeShipping,omitempty"`
	IsSmileShipping     bool   `json:"isSmileShipping,omitempty"`
	IsExpressShipping   bool   `json:"isExpressShipping,omitempty"`
	IsCartVisible       bool   `json:"isCartVisible,omitempty"`
	IsBigSmileItem      bool   `json:"isBigSmileItem,omitempty"`
	ImageChgDt          int    `json:"ImageChgDt,omitempty"`
}

type Batch struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"createdat"`
}

type Latest struct {
	ID      int64 `json:"id"`
	BatchId int64 `json:"batchid"`
}
