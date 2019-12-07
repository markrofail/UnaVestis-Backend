package hm

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetAllTitles() {
	resp, err := http.Get("https://eg.hm.com/en/")
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	idMainMenu := ".menu__list.menu--one__list"
	mainList := doc.Find("ul" + idMainMenu)

	getMenu1Items(mainList)
}

func getMenu1Items(item *goquery.Selection) {
	idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", "one")

	item.Find("li" + idMenu).Each(func(index int, menu1 *goquery.Selection) {
		idMenuItemText := fmt.Sprintf(".menu__link.menu--%s__link", "one")
		idMenuItemBody := fmt.Sprintf(".menu__list.menu--%s__list", "two")

		title := menu1.Find(idMenuItemText).Text()
		fmt.Println(title)

		body := menu1.Find(idMenuItemBody)
		getMenu2Items(body, "Shop By Product")
	})
}

func getMenu2Items(item *goquery.Selection, keyword string) {
	idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", "two")

	item.Find("li" + idMenu).Each(func(index int, menu2 *goquery.Selection) {
		idMenuItemText := fmt.Sprintf(".menu__link.menu-%s__link", "two")
		idMenuItemBody := fmt.Sprintf(".menu__list.menu--%s__list", "three")

		title := menu2.Find(idMenuItemText).Text()
		fmt.Println("> " + title)

		body := menu2.Find(idMenuItemBody)
		if title == keyword{
			getMenu3Items(body)
		}
	})
}

func getMenu3Items(item *goquery.Selection) {
	idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", "three")

	item.Find("li" + idMenu).Each(func(i int, menuItem *goquery.Selection) {
		idMenuItemText := fmt.Sprintf(".menu__link.menu--%s__link", "three")

		title := menuItem.Find(idMenuItemText).Text()
		fmt.Println("> > " + title)
	})
}
