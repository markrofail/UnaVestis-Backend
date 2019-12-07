package daos

import (
	"fmt"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/daos/helpers"
	"log"
	"regexp"
	"strings"
)

func GetPage(productCategory string, itemType string) string {
	baseURL := "https://eg.hm.com/en"

	processedCategory := strings.ToLower(productCategory)

	processedItem := strings.ToLower(itemType)
	processedItem = strings.Replace(processedItem, " & ", "-", -1)
	processedItem = strings.Replace(processedItem, " ", "-", -1)

	builtURL := fmt.Sprintf("%s/shop-%s/shop-product/%s/", baseURL, processedCategory, processedItem)
	return builtURL
}

func GetViewArgs(productCategory string, itemType string) string {
	pageURL := GetPage(productCategory, itemType)

	doc := helpers.GetResponse(pageURL)

	link := doc.Find("link[rel=latest-version]")
	viewArgsLink, _ := link.Attr("href")

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(viewArgsLink, "")
	return processedString
}
