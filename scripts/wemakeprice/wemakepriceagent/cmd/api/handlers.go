package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pilseong/wemakepriceagent/models"
)

func (app *Config) getData() {

	wemakepriceServiceURL := "https://front.wemakeprice.com/api/listing/v1/best/detail.json?groupId&categoryId&listType&page=1&perPage=200&maxCount&cacheId&os=pc&version&domain=listing-api.wemakeprice.com&path=%2Fv1%2Fbest%2Fdetail&debug=false"

	// request 생성
	request, err := http.NewRequest("GET",
		wemakepriceServiceURL,
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
	var result models.WemakepriceBEST
	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	// log.Println(PrettyPrint(result))

	// // DB 저장
	app.BulkCreate(result.Data.Deals)

}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i any) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
