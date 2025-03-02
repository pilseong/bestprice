package models

import "time"

// User is the structure which holds one user from the database.
type TmonBEST struct {
	Data []Item `json:"data"`
}

type Item struct {
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	Title        string `json:"title"`
	DealNo       int64  `json:"dealNo"`
	DeliveryInfo struct {
		DeliveryFeePolicy string `json:"deliveryFeePolicy"`
		DeliveryFee       int    `json:"deliveryFee"`
	} `json:"deliveryInfo,omitempty"`
	CategoryInfo struct {
		CategoryName string `json:"categoryName"`
	} `json:"categoryInfo"`
	LaunchInfo struct {
		URL string `json:"url"`
	} `json:"launchInfo"`
	ImageInfo struct {
		Pc3ColImageURL string `json:"pc3ColImageUrl"`
	} `json:"imageInfo"`
	DiscountPrice struct {
		OriginalPrice int `json:"originalPrice"`
		Price         int `json:"price"`
		DiscountRate  int `json:"discountRate"`
	} `json:"discountPrice"`
}

// type Item struct {
// 	StatusType             string `json:"statusType"`
// 	StartDate              string `json:"startDate"`
// 	EndDate                string `json:"endDate"`
// 	TicketStartDate        string `json:"ticketStartDate"`
// 	TicketEndDate          string `json:"ticketEndDate"`
// 	TitleDesc              string `json:"titleDesc"`
// 	DealListPromotionTitle string `json:"dealListPromotionTitle"`
// 	Title                  string `json:"title"`
// 	TitleName              string `json:"titleName"`
// 	TitleImage             string `json:"titleImage"`
// 	SubTitleImage          string `json:"subTitleImage"`
// 	DealType               string `json:"dealType"`
// 	SortType               string `json:"sortType"`
// 	DeliveryInfo           struct {
// 		DeliveryFeePolicy string `json:"deliveryFeePolicy"`
// 		DeliveryFee       int    `json:"deliveryFee"`
// 	} `json:"deliveryInfo,omitempty"`
// 	CategoryInfo struct {
// 		CategoryName string `json:"categoryName"`
// 	} `json:"categoryInfo"`
// 	ContentInfo struct {
// 		EmergencyCustomer bool   `json:"emergencyCustomer"`
// 		MartPromotion1    string `json:"martPromotion1"`
// 		MartPromotion2    string `json:"martPromotion2"`
// 		DealListUseDate   struct {
// 			StartDate string `json:"startDate"`
// 			EndDate   string `json:"endDate"`
// 			Type      string `json:"type"`
// 			IsView    bool   `json:"isView"`
// 		} `json:"dealListUseDate"`
// 		AddChargeMsg           string `json:"addChargeMsg"`
// 		HideDeliveryBannerInfo bool   `json:"hideDeliveryBannerInfo"`
// 		ContentImages          struct {
// 			ProductMainImage string   `json:"productMainImage"`
// 			FrontImages      []string `json:"frontImages"`
// 			MarketingImage   string   `json:"marketingImage"`
// 			IntroImages      []any    `json:"introImages"`
// 		} `json:"contentImages"`
// 	} `json:"contentInfo"`
// 	ImageInfo struct {
// 		MobileImageURL     string `json:"mobileImageUrl"`
// 		Mobile3ColImageURL string `json:"mobile3ColImageUrl"`
// 		Pc3ColImageURL     string `json:"pc3ColImageUrl"`
// 		AroundImageURL     string `json:"aroundImageUrl"`
// 	} `json:"imageInfo"`
// 	SubmitInfo struct {
// 	} `json:"submitInfo,omitempty"`
// 	UserPoint struct {
// 		ReviewCount   int     `json:"reviewCount"`
// 		ReviewAverage float64 `json:"reviewAverage"`
// 		Satisfaction  float64 `json:"satisfaction"`
// 	} `json:"userPoint"`
// 	DiscountPrice struct {
// 		OriginalPrice                 int    `json:"originalPrice"`
// 		TmonPrice                     int    `json:"tmonPrice"`
// 		CouponPrice                   int    `json:"couponPrice"`
// 		Price                         int    `json:"price"`
// 		DiscountAmount                int    `json:"discountAmount"`
// 		DiscountRate                  int    `json:"discountRate"`
// 		TmonDiscountRate              int    `json:"tmonDiscountRate"`
// 		PolicyDiscountRate            int    `json:"policyDiscountRate"`
// 		DealDiscountRate              int    `json:"dealDiscountRate"`
// 		Description                   string `json:"description"`
// 		DiscountLabel                 string `json:"discountLabel"`
// 		AdditionalDiscountOptionExist bool   `json:"additionalDiscountOptionExist"`
// 	} `json:"discountPrice"`
// 	SearchInfo struct {
// 		TravelCityType []any  `json:"travelCityType"`
// 		ProductType    []any  `json:"productType"`
// 		SearchKeyword  string `json:"searchKeyword"`
// 	} `json:"searchInfo"`
// }

type Batch struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"createdat"`
}

type Latest struct {
	ID      int64 `json:"id"`
	BatchId int64 `json:"batchid"`
}
