package main

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/playwright-community/playwright-go"
)

var option playwright.BrowserTypeLaunchOptions

type Item struct {
	Rank          int    `json:"rank"`
	Title         string `json:"title"`
	LinkInfo      string `json:"linkInfo"`
	ImageLink     string `json:"imageLink"`
	OriginalPrice string `json:"originalPrice"`
	DiscountPrice string `json:"discountPrice"`
	DiscountRate  string `json:"discountRate"`
	ShippingText  string `json:"shippingText"`
}

func main() {

	option.Headless = playwright.Bool(true)

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(option)
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	page.SetViewportSize(1280, 1080)

	// page.On("response", func(response playwright.Response) {
	// 	log.Println(response.URL())
	// 	if strings.Contains(response.URL(), "products") {
	// 		log.Println("inside")
	// 		bytes, err := response.Body()
	// 		if err != nil {
	// 			log.Fatal()
	// 		}

	// 		log.Printf("url: %v body %v", response.URL(), string(bytes))
	// 	}
	// })

	if _, err = page.Goto("https://www.11st.co.kr/browsing/BestSeller.tmall?method=getBestSellerMain&xfrom=main^gnb"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	})

	for i := 0; i < 5; i++ {
		page.Evaluate("window.scrollTo(0, document.body.scrollHeight);")
		// log.Println("scrolling key press", i)
		time.Sleep(100 * time.Millisecond)
	}

	allItems, err := page.Locator("#bestPrdList > div > ul.cfix > li").All()
	if err != nil {
		log.Fatal()
	}

	var items []Item
	log.Println(len(allItems))
	for i := 0; i < 20; i++ {

		rank, err := allItems[i].Locator("div > a > span.best").TextContent()
		if err != nil {
			log.Println("cannot get the rank", err.Error())
			rank = "0"
		}
		// log.Println(rank)
		rankInt, err := strconv.Atoi(rank)
		if err != nil {
			log.Println("cannot get the rank while parsing", err.Error())
			rankInt = 0
		}

		jsonResult, err := allItems[i].Locator("div > div.store > a").GetAttribute("data-log-body")
		if err != nil {
			log.Println("cannot get the productId while parsing", err.Error())
			rankInt = 0
		}

		type ProductId struct {
			ProductNo string `json:"product_no"`
		}
		// var jsonMap map[string]interface{}
		var productId ProductId

		json.Unmarshal([]byte(strings.Replace(jsonResult, "'", "\"", -1)), &productId)
		log.Println(productId.ProductNo)

		title, err := allItems[i].Locator("div > a > div.pname > p").TextContent()
		if err != nil {
			log.Println("cannot get the title", err.Error())
			title = ""
		}
		// log.Println(title)

		link, err := allItems[i].Locator("div > a").First().GetAttribute("href")
		if err != nil {
			log.Fatal("cannot get the link", err.Error())
		}
		// log.Println(link)

		imageLink, err := allItems[i].Locator("div > a > div.img_plot > img").GetAttribute("src")
		if err != nil {
			log.Fatal("cannot get the imageLink", err.Error())
		}
		// log.Println(imageLink)

		originalPrice, err := allItems[i].Locator("div > a > div.pname s.normal_price").TextContent(
			playwright.LocatorTextContentOptions{
				Timeout: playwright.Float(100),
			},
		)

		if err != nil {
			log.Println("cannot get the original price", err.Error())
			originalPrice = ""
		}
		// log.Println(originalPrice)

		discountPrice, err := allItems[i].Locator("div > a > div.pname strong.sale_price").TextContent()
		if err != nil {
			log.Fatal("cannot get the discount price", err.Error())
		}
		// log.Println(discountPrice)

		discountRate, err := allItems[i].Locator("div > a > div.pname span.sale").TextContent(
			playwright.LocatorTextContentOptions{
				Timeout: playwright.Float(100),
			},
		)
		if err != nil {
			log.Println("cannot get the discount rate", err.Error())
			discountRate = ""
		}
		log.Println(discountRate)

		shippingText, err := allItems[i].Locator("div > a > div.pname div.s_flag > em").TextContent(
			playwright.LocatorTextContentOptions{
				Timeout: playwright.Float(100),
			},
		)
		if err != nil {
			log.Println("cannot get the shipping text", err.Error())
			shippingText = ""
		}
		// log.Println(shippingText)

		item := Item{
			Rank:          rankInt,
			Title:         title,
			LinkInfo:      link,
			ImageLink:     imageLink,
			OriginalPrice: originalPrice,
			DiscountPrice: discountPrice + "원",
			DiscountRate:  discountRate,
			ShippingText:  shippingText,
		}

		items = append(items, item)

	}

	time.Sleep(time.Second)
	browser.Close()
}
