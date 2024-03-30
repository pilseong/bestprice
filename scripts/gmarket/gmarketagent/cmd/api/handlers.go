package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pilseong/gmarketagent/models"
)

type GmarketBEST struct {
	PageProps struct {
		ServerData struct {
			Home struct {
				Items []models.Item `json:"itemData,omitempty"`
			} `json:"home,omitempty"`
		} `json:"serverData,omitempty"`
	} `json:"pageProps,omitempty"`
}

// func (app *Config) mockData() {
// 	items := []models.Item{{
// 		ID:                  0,
// 		BatchID:             0,
// 		Rank:                0,
// 		GoodsCode:           "3030854245",
// 		GoodsName:           "(15%+10%) 초등 도서 베스트 100 특가",
// 		LinkURL:             "http://item.gmarket.co.kr/Item?goodscode=3030854245",
// 		DiscountPrice:       "15,370",
// 		Price:               "22,000",
// 		DeliveryInfo:        "무료",
// 		ItemPriceType:       "Normal",
// 		ImageURL:            "//gdimg.gmarket.co.kr/3030854245/still",
// 		DiscountRate:        30,
// 		ExpressShippingText: "",
// 		DeliveryText:        "무료배송",
// 		ConsultingPeriod:    "",
// 		IsPriceExpose:       true,
// 		IsStartPrice:        false,
// 		IsFreeShipping:      true,
// 		IsSmileShipping:     false,
// 		IsExpressShipping:   false,
// 		IsCartVisible:       false,
// 		IsBigSmileItem:      false,
// 		ImageChgDt:          1711519780,
// 	}}

// 	app.BulkCreate(items)
// }

func (app *Config) getData() {

	gmarketServiceURL := "https://www.gmarket.co.kr/n/best/_next/data/7cMUsHihI87LgmKG-KVYx/n/best.json"

	// request 생성
	request, err := http.NewRequest("GET",
		gmarketServiceURL,
		nil)

	if err != nil {
		log.Println(time.Now(), "Cannot make request with url: "+err.Error())
		return
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0")
	request.Header.Add("x-nextjs-data", "1")

	// rest client 생성 및 인증 호출
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(time.Now(), "Error occurred while sending requets to url "+err.Error())
		return
	}

	defer response.Body.Close()

	// 오류 일 경우 에러 반환
	if response.StatusCode != http.StatusOK {
		log.Println(time.Now(), "Not successfull request "+response.Status)
		return
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(time.Now(), "Cannot read body from response "+err.Error())
		return
	}

	// 결과 출력
	var result GmarketBEST
	if err := json.Unmarshal(bytes, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// DB 저장
	app.BulkCreate(result.PageProps.ServerData.Home.Items)

}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i any) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
