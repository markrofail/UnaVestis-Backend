package daos

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func GetAllColor() []string {
	resp, err := http.Get("https://eg.hm.com/en/shop-ladies/new-arrivals/clothes/")
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	idMainMenu := ".item-list__swatch_list"
	colorList := doc.Find("ul" + idMainMenu)

	var colors []string
	colorList.Find("li.facet-item").Each(func(i int, selection *goquery.Selection) {
		color := selection.Find("span.facet-item__value").Clone().Children().Remove().End().Text()
		colors = append(colors, strings.TrimSpace(color))
	})
	return colors
}
