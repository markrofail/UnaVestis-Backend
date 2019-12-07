package hm

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/markrofail/fashion_scraping_api/helpers"
)

var numericAlpha = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

type MenuItem struct {
	Title string
	Body  *goquery.Selection
}

func GetItemTypes(kind string) []string {
	doc := helpers.GetResponse("https://eg.hm.com/en/")

	idMainMenu := ".menu__list.menu--one__list"
	mainList := doc.Find("ul" + idMainMenu)

	idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", numericAlpha[1])

	var typeArr []string
	mainList.Find("li" + idMenu).Each(func(index int, menu1 *goquery.Selection) {
		idMenuItemText := fmt.Sprintf(".menu__link.menu--%s__link", numericAlpha[1])
		idMenuItemBody := fmt.Sprintf(".menu__list.menu--%s__list", numericAlpha[2])

		title := menu1.Find(idMenuItemText).Text()
		body := menu1.Find(idMenuItemBody)

		menuItem := MenuItem{
			Title: title,
			Body:  body,
		}

		if menuItem.Title == kind {
			idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", numericAlpha[2])

			menuItem.Body.Find("li" + idMenu).Each(func(index int, menu2 *goquery.Selection) {
				idMenuItemText := fmt.Sprintf(".menu__link.menu-%s__link", numericAlpha[2])
				idMenuItemBody := fmt.Sprintf(".menu__list.menu--%s__list", numericAlpha[3])

				title := menu2.Find(idMenuItemText).Text()
				body := menu2.Find(idMenuItemBody)

				menuItem := MenuItem{
					Title: title,
					Body:  body,
				}

				if menuItem.Title == "Shop By Product" {
					idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", numericAlpha[3])

					menuItem.Body.Find("li" + idMenu).Each(func(i int, menuItem *goquery.Selection) {
						idMenuItemText := fmt.Sprintf(".menu__link.menu--%s__link", numericAlpha[3])

						title := menuItem.Find(idMenuItemText).Text()
						typeArr = append(typeArr, title)
					})
				}
			})
		}
	})

	return typeArr
}

func GetAllKinds() []string {
	doc := helpers.GetResponse("https://eg.hm.com/en/")

	idMainMenu := ".menu__list.menu--one__list"
	mainList := doc.Find("ul" + idMainMenu)

	var kinds []string

	idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", numericAlpha[1])
	mainList.Find("li" + idMenu).Each(func(index int, menu1 *goquery.Selection) {
		idMenuItemText := fmt.Sprintf(".menu__link.menu--%s__link", numericAlpha[1])

		title := menu1.Find(idMenuItemText).Text()
		kinds = append(kinds, title)
	})

	return kinds
}

func GetAllTitles() {
	doc := helpers.GetResponse("https://eg.hm.com/en/")

	idMainMenu := ".menu__list.menu--one__list"
	mainList := doc.Find("ul" + idMainMenu)

	nextMenuKeywords := []string{"Men", "Ladies", "Kids"}
	getMenu1Items(mainList, nextMenuKeywords)
}

func getMenu1Items(item *goquery.Selection, keywords []string) {
	idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", numericAlpha[1])

	item.Find("li" + idMenu).Each(func(index int, menu1 *goquery.Selection) {
		idMenuItemText := fmt.Sprintf(".menu__link.menu--%s__link", numericAlpha[1])
		idMenuItemBody := fmt.Sprintf(".menu__list.menu--%s__list", numericAlpha[2])

		title := menu1.Find(idMenuItemText).Text()
		body := menu1.Find(idMenuItemBody)

		menuItem := MenuItem{
			Title: title,
			Body:  body,
		}

		fmt.Println(menuItem.Title)
		nextMenuKeywords := []string{"Shop By Product"}
		if len(keywords) > 0 {
			for _, keyword := range keywords {
				if menuItem.Title == keyword {
					getMenu2Items(menuItem.Body, nextMenuKeywords)
				}
			}
		} else {
			getMenu2Items(menuItem.Body, nextMenuKeywords)
		}
	})
}

func getMenu2Items(item *goquery.Selection, keywords []string) {
	idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", numericAlpha[2])

	item.Find("li" + idMenu).Each(func(index int, menu2 *goquery.Selection) {
		idMenuItemText := fmt.Sprintf(".menu__link.menu-%s__link", numericAlpha[2])
		idMenuItemBody := fmt.Sprintf(".menu__list.menu--%s__list", numericAlpha[3])

		title := menu2.Find(idMenuItemText).Text()
		body := menu2.Find(idMenuItemBody)

		menuItem := MenuItem{
			Title: title,
			Body:  body,
		}

		fmt.Println("> " + menuItem.Title)
		if len(keywords) > 0 {
			for _, keyword := range keywords {
				if menuItem.Title == keyword {
					getMenu3Items(menuItem.Body)
				}
			}
		} else {
			getMenu3Items(menuItem.Body)
		}
	})
}

func getMenu3Items(item *goquery.Selection) {
	idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", numericAlpha[3])

	item.Find("li" + idMenu).Each(func(i int, menuItem *goquery.Selection) {
		idMenuItemText := fmt.Sprintf(".menu__link.menu--%s__link", numericAlpha[3])

		title := menuItem.Find(idMenuItemText).Text()
		fmt.Println(" > " + title)
	})
}
