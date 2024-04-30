package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pilseong/elevenamazonagent/models"
)

func (app *Config) getData() error {
	elmazonServiceURL := "https://apis.11st.co.kr/pui/v2/page?pageId=APCBEST&blckSn=7830&pageMode=NEXT&pageNo="

	var save_target []models.Item
	for i := 1; i <= 5; i++ {
		bytes, err := apiCall(fmt.Sprintf("%s%d", elmazonServiceURL, i))

		if err != nil {
			log.Println("Unsuccessful request api call", err)
			return err
		}

		// 결과 출력
		var result models.ElmazonBest
		if err := json.Unmarshal(bytes, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
			return err
		}
		// log.Println("size ", len(result.Data[0].BlockList[0].List))
		// log.Println(PrettyPrint(result.Data[0].BlockList[0].List))

		save_target = append(save_target, result.Data[0].BlockList[0].List...)
	}

	// DB 저장
	// log.Println("size", len(save_target))
	app.BulkCreate(save_target)

	return nil

}

func apiCall(elmazonServiceURL string) ([]byte, error) {
	// request 생성
	request, err := http.NewRequest(http.MethodGet,
		elmazonServiceURL,
		nil)

	if err != nil {
		log.Println(time.Now(), "Cannot make request with url: "+err.Error())
		return nil, err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0")
	request.Header.Add("x-nextjs-data", "1")

	// rest client 생성 및 인증 호출
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(time.Now(), "Error occurred while sending requets to url "+err.Error())
		return nil, err
	}

	defer response.Body.Close()

	// 오류 일 경우 에러 반환
	if response.StatusCode != http.StatusOK {
		log.Println(time.Now(), "Not successfull request "+response.Status)
		return nil, err
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(time.Now(), "Cannot read body from response "+err.Error())
		return nil, err
	}
	return bytes, err
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i any) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
