package eleventst

import "time"

type Item struct {
	ID            int64  `json:"id"`
	BatchID       int64  `json:"batchId"`
	ProductId     string `json:"productId"`
	Rank          int    `json:"rank"`
	Title         string `json:"title"`
	LinkInfo      string `json:"linkInfo"`
	ImageLink     string `json:"imageLink"`
	OriginalPrice string `json:"originalPrice"`
	DiscountPrice string `json:"discountPrice"`
	DiscountRate  string `json:"discountRate"`
	ShippingText  string `json:"shippingText"`
}

type Batch struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"createdat"`
}

type Latest struct {
	ID      int64 `json:"id"`
	BatchId int64 `json:"batchid"`
}
