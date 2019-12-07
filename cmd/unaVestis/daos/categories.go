package daos

import (
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/daos/shops/hm"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/models"
)

// ProductDAO persists user data in database
type CategoryDAO struct{}

// NewUserDAO creates a new ProductDAO
func NewCategoryDAO() *CategoryDAO {
	return &CategoryDAO{}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *CategoryDAO) GetAll() ([]*models.Category, error) {
	categories, _ := daos.GetAllCategories()

	return categories, nil
}
