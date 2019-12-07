package hm

import (
	"fmt"
	"github.com/markrofail/fashion_scraping_api/helpers"
	"log"
	"regexp"
	"strings"
)

func GetPage(kind string, itemType string) string {
	baseURL := "https://eg.hm.com/en"

	processedKind := strings.ToLower(kind)

	processedItem := strings.ToLower(itemType)
	processedItem = strings.Replace(processedItem, " & ", "-", -1)
	processedItem = strings.Replace(processedItem, " ", "-", -1)

	builtURL := fmt.Sprintf("%s/shop-%s/shop-product/%s/", baseURL, processedKind, processedItem)
	return builtURL
}

func GetViewArgs(kind string, itemType string) string {
	pageURL := GetPage(kind, itemType)

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
