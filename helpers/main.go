package helpers

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetResponse(pageUrl string) *goquery.Document {
	resp, err := http.Get(pageUrl)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}
