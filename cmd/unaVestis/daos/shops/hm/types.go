package daos

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/daos/helpers"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/models"
	"strings"
)

func GetTypesByCategory(productCategory string) ([]*models.Type, error) {
	fmt.Println(productCategory)
	doc := helpers.GetResponse("https://eg.hm.com/en/")

	idMainMenu := ".menu__list.menu--one__list"
	mainList := doc.Find("ul" + idMainMenu)

	idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", numericAlpha[1])

	var typeArr []*models.Type
	mainList.Find("li" + idMenu).Each(func(index int, menu1 *goquery.Selection) {
		idMenuItemText := fmt.Sprintf(".menu__link.menu--%s__link", numericAlpha[1])
		idMenuItemBody := fmt.Sprintf(".menu__list.menu--%s__list", numericAlpha[2])

		title := menu1.Find(idMenuItemText).Text()
		body := menu1.Find(idMenuItemBody)

		menuItem := MenuItem{
			Title: title,
			Body:  body,
		}

		fmt.Println(menuItem.Title)
		if menuItem.Title == strings.ToLower(productCategory) {
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
						typeArr = append(typeArr, &models.Type{title})
					})
				}
			})
		}
	})

	return typeArr, nil
}
