package daos

import (
	"fmt"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/daos/shops/hm"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/models"
	"log"
)

// ProductDAO persists user data in database
type ProductDAO struct{}

// NewUserDAO creates a new ProductDAO
func NewProductDAO() *ProductDAO {
	return &ProductDAO{}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *ProductDAO) Get(id uint) (*models.Product, error) {
	var product models.Product

	// Query Database here...

	//user = models.User{
	//	Model: models.Model{ID: 1},
	//	FirstName: "Martin",
	//	LastName: "Heinz",
	//	Address: "Not gonna tell you",
	//	Email: "martin7.heinz@gmail.com"}

	// if using Gorm:
	//err := config.Config.DB.Where("id = ?", id).
	//	First(&user).
	//	Error

	return &product, nil
}

func (dao *ProductDAO) GetAll(productCategory string, productType string) (*[]models.Product, error) {
	requestUrl := daos.BuildURL(productCategory, productType)
	testPage := daos.GetPage(productCategory, productType)
	fmt.Println(testPage)

	rawJson := daos.GetJSON(requestUrl)
	requestData, err := daos.ExtractData(rawJson)
	if err != nil {
		log.Println(err)
	}

	products := daos.GetItems(requestData)

	fmt.Println(productCategory, productType)
	return &products, nil
}
