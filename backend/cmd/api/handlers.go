package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type CommonItem struct {
	ShoppingMall string `json:"shoppingMall"`
	ID           int64  `json:"id"`
	Rank         int    `json:"rank"`
	Title        string `json:"title"`
	LinkURL      string `json:"linkUrl"`
	ImageURL     string `json:"imageUrl"`
	BeforePrice  int    `json:"beforePrice"`
	AfterPrice   int    `json:"afterPrice"`
	DiscountRate int    `json:"discountRate"`
	ShippingInfo string `json:"shippingInfo"`
}

// Home displays the status of the api, as JSON.
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllGmarketItems(w http.ResponseWriter, r *http.Request) {
	// 값이 주어지지 않았을 경우 기본값 설정
	pageNo, err := strconv.Atoi(chi.URLParam(r, "pageNo"))
	if err != nil {
		pageNo = 1
	}
	pageSize, err := strconv.Atoi(chi.URLParam(r, "pageSize"))
	if err != nil {
		pageSize = 200
	}

	log.Printf("page size %d, page no %d", pageSize, pageNo)

	items, err := app.getGmarketItems(pageNo, pageSize)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, items)
}

func (app *application) AllTmonItems(w http.ResponseWriter, r *http.Request) {
	// 값이 주어지지 않았을 경우 기본값 설정
	pageNo, err := strconv.Atoi(chi.URLParam(r, "pageNo"))
	if err != nil {
		pageNo = 1
	}
	pageSize, err := strconv.Atoi(chi.URLParam(r, "pageSize"))
	if err != nil {
		pageSize = 200
	}

	items, err := app.getTmonItems(pageNo, pageSize)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, items)
}

func (app *application) AllWeMakePriceItems(w http.ResponseWriter, r *http.Request) {
	// 값이 주어지지 않았을 경우 기본값 설정
	pageNo, err := strconv.Atoi(chi.URLParam(r, "pageNo"))
	if err != nil {
		pageNo = 1
	}
	pageSize, err := strconv.Atoi(chi.URLParam(r, "pageSize"))
	if err != nil {
		pageSize = 200
	}

	items, err := app.getWeMakePriceItems(pageNo, pageSize)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, items)
}

func (app *application) All11StItems(w http.ResponseWriter, r *http.Request) {

	// 값이 주어지지 않았을 경우 기본값 설정
	pageNo, err := strconv.Atoi(chi.URLParam(r, "pageNo"))
	if err != nil {
		pageNo = 1
	}
	pageSize, err := strconv.Atoi(chi.URLParam(r, "pageSize"))
	if err != nil {
		pageSize = 200
	}

	log.Printf("page size %d, page no %d", pageSize, pageNo)

	items, err := app.get11stItem(pageNo, pageSize)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, items)
}

func (app *application) FrontItems(w http.ResponseWriter, r *http.Request) {

	var commonItems []CommonItem

	gmarketItems, err := app.getGmarketItems(1, 50)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		app.errorJSON(w, err)
		return
	}

	weMakePriceItems, err := app.getWeMakePriceItems(1, 50)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		app.errorJSON(w, err)
		return
	}

	eleventItems, err := app.get11stItem(1, 50)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		app.errorJSON(w, err)
		return
	}

	tmonItems, err := app.getTmonItems(1, 50)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		app.errorJSON(w, err)
		return
	}

	for i := 0; i < len(gmarketItems); i++ {
		commonItems = append(commonItems, gmarketItems[i])
		commonItems = append(commonItems, weMakePriceItems[i])
		commonItems = append(commonItems, eleventItems[i])
		commonItems = append(commonItems, tmonItems[i])
	}

	_ = app.writeJSON(w, http.StatusOK, commonItems)
}

func (app *application) getGmarketItems(pageNo, pageSize int) ([]CommonItem, error) {
	items, err := app.GmarketDB.AllItems(pageNo, pageSize)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		return nil, err
	}

	var commonItems []CommonItem
	for _, item := range items {

		bPrice, err := strconv.Atoi(strings.Replace(item.Price, ",", "", -1))
		if err != nil {
			log.Panic("cannot convert before price string to number")
		}
		aPrice, err := strconv.Atoi(strings.Replace(item.DiscountPrice, ",", "", -1))
		if err != nil {
			log.Panic("cannot convert after price string to number")
		}

		commonItems = append(commonItems, CommonItem{
			ShoppingMall: "지마켓",
			ID:           item.ID,
			Rank:         item.Rank,
			Title:        item.GoodsName,
			LinkURL:      strings.Replace(item.LinkURL, "http", "https", 1),
			ImageURL:     fmt.Sprintf("https:%s/300?ver=%s", item.ImageURL, item.GoodsCode),
			BeforePrice:  bPrice,
			AfterPrice:   aPrice,
			DiscountRate: item.DiscountRate,
			ShippingInfo: "배송비 " + strings.Replace(item.DeliveryInfo, "<br/>", "", -1),
		})
	}

	return commonItems, nil
}

func (app *application) get11stItem(pageNo, pageSize int) ([]CommonItem, error) {
	items, err := app.ElevenStDB.AllItems(pageNo, pageSize)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		return nil, err
	}

	var commonItems []CommonItem
	for _, item := range items {

		bPriceInt := 0
		if item.OriginalPrice != "" {
			bPrice := strings.Replace(item.OriginalPrice, "원", "", -1)
			bPrice = strings.Replace(bPrice, ",", "", -1)
			bPriceInt, err = strconv.Atoi(bPrice)
			if err != nil {
				log.Panic("cannot convert before price string to number")
			}
		}

		aPriceInt := 0
		if item.DiscountPrice != "" {
			aPrice := strings.Replace(item.DiscountPrice, "원", "", -1)
			aPrice = strings.Replace(aPrice, ",", "", -1)
			aPriceInt, err = strconv.Atoi(aPrice)
			if err != nil {
				log.Panic("cannot convert after price string to number")
			}
		}

		var discountRateInt = 0
		if item.DiscountRate != "" {
			discountRateStr := strings.Replace(item.DiscountRate, "%", "", -1)
			discountRateInt, err = strconv.Atoi(discountRateStr)
			if err != nil {
				log.Panic("cannot convert after discount rate to number")
			}
		}

		var shippingText = ""
		if item.ShippingText == "" {
			shippingText += "배송비 유료"
		} else {
			shippingText += "배송비 무료"
		}

		commonItems = append(commonItems, CommonItem{
			ShoppingMall: "11번가",
			ID:           item.ID,
			Rank:         item.Rank,
			Title:        item.Title,
			LinkURL:      item.LinkInfo,
			ImageURL:     item.ImageLink,
			BeforePrice:  bPriceInt,
			AfterPrice:   aPriceInt,
			DiscountRate: discountRateInt,
			ShippingInfo: shippingText,
		})
	}

	return commonItems, nil
}

func (app *application) getTmonItems(pageNo, pageSize int) ([]CommonItem, error) {
	items, err := app.TmonDB.AllItems(pageNo, pageSize)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		return nil, err
	}

	var commonItems []CommonItem
	for _, item := range items {

		deliveryStr := ""
		if item.DeliveryFeePolicy != "" {
			deliveryStr += "배송비"
		}

		switch item.DeliveryFeePolicy {
		case "FREE":
			deliveryStr += " 무료"

		case "CONDITION":
			deliveryStr += " 조건부 무료"

		case "":
			deliveryStr += "디지털 쿠폰"
		}

		if item.DeliveryFeePolicy != "FREE" && item.DeliveryFeePolicy != "" {
			deliveryStr += fmt.Sprintf(" %d", item.DeliveryFee)
		}

		commonItems = append(commonItems, CommonItem{
			ShoppingMall: "티몬",
			ID:           item.ID,
			Rank:         item.Rank,
			Title:        item.Title,
			LinkURL:      item.URL,
			ImageURL:     item.Pc3ColImageURL,
			BeforePrice:  item.OriginalPrice,
			AfterPrice:   item.Price,
			DiscountRate: item.DiscountRate,
			ShippingInfo: deliveryStr,
		})
	}

	return commonItems, nil
}

func (app *application) getWeMakePriceItems(pageNo, pageSize int) ([]CommonItem, error) {
	items, err := app.WeMakePriceDB.AllItems(pageNo, pageSize)
	if err != nil {
		log.Println("Error occurred while reading database", err.Error())
		return nil, err
	}

	var commonItems []CommonItem
	for _, item := range items {

		host := ""

		switch strings.ToLower(item.LinkType) {
		case "ticket":
			host += "ticket"
		case "tour":
			host += "tour.wd"
		default:
			host += "front"
		}

		param := ""
		switch strings.ToLower(item.LinkType) {
		case "ticket", "prod":
			param += "product/"
		case "tour":
			param += "wmp/deal.html?dealId="
		default:
			param += strings.ToLower(item.LinkType) + "/"
		}

		shippingText := ""
		if item.ShipText != "" {
			shippingText += item.ShipText
		} else {
			shippingText = "배송비 유료"
		}

		var bPrice int
		if item.DiscountPrice != 0 {
			bPrice = item.SalePrice
		} else {
			bPrice = item.OriginPrice
		}

		var aPrice int
		if item.DiscountPrice != 0 {
			aPrice = item.DiscountPrice
		} else {
			aPrice = item.SalePrice
		}

		commonItems = append(commonItems, CommonItem{
			ShoppingMall: "위메프",
			ID:           item.ID,
			Rank:         item.Rank,
			Title:        item.DispNm,
			LinkURL:      fmt.Sprintf("https://%s.wemakeprice.com/%s%d", host, param, item.LinkInfo),
			ImageURL:     item.MediumImgURL,
			BeforePrice:  bPrice,
			AfterPrice:   aPrice,
			DiscountRate: item.DiscountRate,
			ShippingInfo: shippingText,
		})
	}

	return commonItems, nil
}
