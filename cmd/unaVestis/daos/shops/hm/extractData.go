package daos

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/models"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func BuildURL(productCategory string, itemType string) *url.URL {
	baseURL := "https://eg.hm.com/en/views/ajax"

	u, _ := url.Parse(baseURL)
	q, _ := url.ParseQuery(u.RawQuery)

	paramsMap := make(map[string]string)
	paramsMap["view_name"] = "alshaya_product_list"
	paramsMap["view_display_id"] = "block_1"
	paramsMap["view_args"] = GetViewArgs(productCategory, itemType)
	//paramsMap["view_path"] = fmt.Sprintf("/en/shop-%s/tops", gender)
	paramsMap["sort_by"] = "created"
	paramsMap["sort_order"] = "DESC"
	paramsMap["page"] = "1"

	for k, v := range paramsMap {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()
	return u
}

func GetJSON(url *url.URL) string {
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

func ExtractData(data string) (string, error) {
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

var colors = GetAllColor()

func IsColor(word string) string {
	searchWord := strings.ToLower(word)
	searchWord = strings.TrimSpace(searchWord)
	//fmt.Println(searchWord)
	for _, color := range colors {
		if strings.HasPrefix(searchWord, color) || strings.HasSuffix(searchWord, color) {
			return color
		}
	}
	return ""
}

func ExtractColor(title string) string {
	res1 := strings.Split(title, ".")[0]
	res2 := strings.Split(res1, "-")
	for _, element := range res2 {
		result := IsColor(element)
		if result != "" {
			return result
		}
	}

	return ""
}

func GetImages(selection *goquery.Selection) []string {
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

func GetItems(data string) []models.Product {
	htmlDoc, _ := html.Parse(strings.NewReader(data))
	doc := goquery.NewDocumentFromNode(htmlDoc)

	var productArr []models.Product
	doc.Find(".c-products__item").Each(func(i int, s *goquery.Selection) {
		productName, _ := s.Find("article").Attr("gtm-name")
		productPrice, _ := s.Find("article").Attr("gtm-price")
		productColor, _ := s.Find(".product-selected-url").Attr("data--original-url")
		productColor = ExtractColor(productColor)
		imgArray := GetImages(s)

		newProduct := models.Product{
			Name:  productName,
			Price: productPrice,
			Color: productColor,
			Image: imgArray,
		}

		productArr = append(productArr, newProduct)
		fmt.Printf("Name: %s;\tColor: %s;\tPrice: %s\n", productName, productColor, productPrice)
		fmt.Println(imgArray)
		fmt.Println()
	})

	return productArr
}
