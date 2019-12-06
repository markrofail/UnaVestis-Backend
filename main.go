package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	_ "github.com/gocolly/colly"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const typeTops string = "/en/shop-ladies/tops"

func buildUrl() *url.URL {
	baseurl := "https://eg.hm.com/en/views/ajax"

	u, _ := url.Parse(baseurl)
	//fmt.Println("original:", u)

	q, _ := url.ParseQuery(u.RawQuery)

	paramsMap := make(map[string]string)
	paramsMap["view_name"] = "alshaya_product_list"
	paramsMap["view_display_id"] = "block_1"
	paramsMap["view_args"] = "65951"
	paramsMap["view_path"] = "/en/shop-ladies/tops"
	paramsMap["sort_by"] = "created"
	paramsMap["sort_order"] = "DESC"
	paramsMap["page"] = "1"

	for k, v := range paramsMap {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()
	//fmt.Println("modified:", u)

	return u
}

func build_colly() *colly.Collector {
	c := colly.NewCollector()

	div_id := ".c-products__item"
	c.OnHTML(div_id, func(e *colly.HTMLElement) {
		fmt.Println(e.Index)
	})

	c.OnRequest(func(r *colly.Request) {
	})

	return c
}

func getJson(url *url.URL) string {
	resp, err := http.Get(url.String())
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	s := string(body)

	return s
}

type HMresponse struct {
	Command  string   `json:"command"`
	Method   string   `json:"method"`
	Selector string   `json:"selector"`
	Data     string   `json:"data"`
	Args     []string `json:"args"`
}

func extractData(data string) (string, error) {
	var responses []HMresponse
	if err := json.Unmarshal([]byte(data), &responses); err != nil {
		fmt.Printf("Error whilde decoding %v\n", err)
		return "", err
	}

	for _, element := range responses {
		if element.Command == "insert" {
			return element.Data, nil
		}
	}

	return "", nil
}

type Product struct {
	Name   string   `json:"name"`
	Price  string   `json:"price"`
	Type   string   `json:"type"`
	Color  string   `json:"color"`
	Gender string   `json:"gender"`
	Shop   string   `json:"shop"`
	Image  []string `json:"image_url"`
}

var colors = []string{
	"red", "orange", "yellow", "green", "maroon", "cyan", "brown", "magenta", "tan",
	"blue", "purple", "indigo", "violet", "olive", "white", "gray", "grey", "beige",
	"navy", "aquamarine", "turquoise", "silver", "pink", "black", "lime", "teal",
}

func isColor(word string) bool {
	searchWord := strings.ToLower(word)
	searchWord = strings.TrimSpace(searchWord)
	//fmt.Println(searchWord)
	for _, color := range colors {
		if strings.HasPrefix(searchWord, color) || strings.HasSuffix(searchWord, color) {
			return true
		}
	}
	return false
}

func extractColor(title string) string {
	res1 := strings.Split(title, ".")[0]
	res2 := strings.Split(res1, "-")
	for _, element := range res2 {
		if isColor(element) {
			return element
		}
	}

	return ""
}

func getImages(selection *goquery.Selection) []string {
	var imgArray []string

	images := selection.Find("img")
	for _, img := range images.Nodes {
		for _, attrs := range img.Attr {
			if attrs.Key == "data-src" {
				imgSrc := attrs.Val
				imgSrc = "https://eg.hm.com" + imgSrc
				imgArray = append(imgArray, imgSrc)
			}
		}
	}

	return imgArray
}

func getItems(data string) []Product {
	htmlDoc, _ := html.Parse(strings.NewReader(data))
	doc := goquery.NewDocumentFromNode(htmlDoc)

	var productArr []Product
	doc.Find(".c-products__item").Each(func(i int, s *goquery.Selection) {
		product_name, _ := s.Find("article").Attr("gtm-name")
		product_price, _ := s.Find("article").Attr("gtm-price")
		product_color, _ := s.Find(".product-selected-url").Attr("data--original-url")
		product_color = extractColor(product_color)
		imgArray := getImages(s)

		newProduct := Product{
			Name:   product_name,
			Price:  product_price,
			Type:   "",
			Color:  product_color,
			Gender: "",
			Shop:   "",
			Image:  imgArray,
		}

		productArr = append(productArr, newProduct)
		fmt.Printf("Name: %s;\tColor: %s;\tPrice: %s\n", product_name, product_color, product_price)
		fmt.Println(imgArray)
		fmt.Println()
	})

	return productArr
}

func main() {
	requestUrl := buildUrl()
	rawJson := getJson(requestUrl)
	requestData, err := extractData(rawJson)
	if err != nil {
		log.Println(err)
	}

	products := getItems(requestData)

	jsonResponse, err := json.Marshal(products)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(jsonResponse))
}
