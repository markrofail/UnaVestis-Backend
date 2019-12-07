package daos

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/daos/helpers"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/models"
)

func GetAllCategories() ([]*models.Category, error) {
	doc := helpers.GetResponse("https://eg.hm.com/en/")

	idMainMenu := fmt.Sprintf(".menu__list.menu--%s__list", numericAlpha[1])
	mainList := doc.Find("ul" + idMainMenu)

	idMenu := fmt.Sprintf(".menu__list-item.menu--%s__list-item", numericAlpha[1])
	menuItems := mainList.Find("li" + idMenu)

	var productCategories []*models.Category
	menuItems.Each(func(index int, menu1 *goquery.Selection) {
		idMenuItemText := fmt.Sprintf(".menu__link.menu--%s__link", numericAlpha[1])
		title := menu1.Find(idMenuItemText).Text()

		category := models.Category{
			Name: title,
		}

		productCategories = append(productCategories, &category)
	})
	return productCategories, nil
}
