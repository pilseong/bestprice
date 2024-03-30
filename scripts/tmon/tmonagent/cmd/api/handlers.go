package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pilseong/tomonagent/models"
)

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

	tmonServiceURL := "https://www.tmon.co.kr/api/home/direct/v1/tabapi/api/best/dealInfoDeals?platformType=PC_WEB&useHeader=AG&categoryNo=0&size=100&bestType=best_100&useCache=CACHE_S60&gender=0&age=0&price=0"

	// request 생성
	request, err := http.NewRequest("GET",
		tmonServiceURL,
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
	var result models.TmonBEST
	if err := json.Unmarshal(bytes, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// log.Println(PrettyPrint(result))

	// // DB 저장
	app.BulkCreate(result.Data)

}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i any) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
