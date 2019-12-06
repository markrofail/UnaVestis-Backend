package main

import (
	"fmt"
	_ "github.com/gocolly/colly"
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
	fmt.Println("original:", u)

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
	return u
}

func main() {

	fmt.Println("modified:", u)

	//fmt.Println(encode_url(typeTops))
	//c := colly.NewCollector()
	//
	////// Find and visit all links
	////c.OnHTML("div", func(e *colly.HTMLElement) {
	////	e.Request.Visit(e.Attr("href"))
	////})
	//
	//div_id := ".c-products__item"
	//c.OnHTML(div_id, func(e *colly.HTMLElement) {
	//	fmt.Println(e.Index)
	//})
	//
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL)
	//})
	//
	//url := "https://eg.hm.com/en/shop-ladies/new-arrivals/clothes/"
	//c.Visit(url)
}
