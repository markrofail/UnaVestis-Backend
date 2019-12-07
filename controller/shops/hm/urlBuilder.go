package hm

import (
	"fmt"
	"github.com/markrofail/fashion_scraping_api/helpers"
	"log"
	"regexp"
	"strings"
)

func GetPage(kind string, itemType string) string {
	baseUrl := "https://eg.hm.com/en"

	processedKind := strings.ToLower(kind)

	processedItem := strings.ToLower(itemType)
	processedItem = strings.Replace(processedItem, " & ", "-", -1)
	processedItem = strings.Replace(processedItem, " ", "-", -1)

	builtUrl := fmt.Sprintf("%s/shop-%s/shop-product/%s/", baseUrl, processedKind, processedItem)
	return builtUrl
}

func GetViewArgs(kind string, itemType string) string {
	pageUrl := GetPage(kind, itemType)

	doc := helpers.GetResponse(pageUrl)

	link := doc.Find("link[rel=latest-version]")
	viewArgsLink, _ := link.Attr("href")

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(viewArgsLink, "")
	return processedString
}
