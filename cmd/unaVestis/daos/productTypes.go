package daos

import (
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/daos/shops/hm"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/models"
)

// ProductDAO persists user data in database
type TypeDAO struct{}

// NewUserDAO creates a new ProductDAO
func NewTypeDAO() *CategoryDAO {
	return &CategoryDAO{}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *CategoryDAO) Get(productCategory string) ([]*models.Type, error) {
	categories, _ := daos.GetTypesByCategory(productCategory)

	return categories, nil
}
