package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//type ClothType string
//
//const(
//	Top ClothType = "top"
//)

const typeTops string = "/en/shop-ladies/tops"

func build_url() *url.URL {
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

type HMresponse struct {
	Command  string   `json:"command"`
	Method   string   `json:"method"`
	Selector string   `json:"selector"`
	Data     string   `json:"data"`
	Args     []string `json:"args"`
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

func get_json(url *url.URL) string {
	resp, err := http.Get(url.String())
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	s := string(body)

	return s
}

func decode_json(data string) (string, error) {
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
func main() {
	requestUrl := build_url()
	rawJson := get_json(requestUrl)
	requestData, err := decode_json(rawJson)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(requestData)
}
