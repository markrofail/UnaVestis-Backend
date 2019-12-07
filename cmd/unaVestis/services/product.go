package services

import "github.com/markrofail/fashion_scraping_api/cmd/unaVestis/models"

type productDAO interface {
	Get(id uint) (*models.Product, error)
	GetAll(ind string, productType string) (*[]models.Product, error)
}

type ProductService struct {
	dao productDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewProductService(dao productDAO) *ProductService {
	return &ProductService{dao}
}

// Get just retrieves user using User DAO, here can be additional logic for processing data retrieved by DAOs
func (s *ProductService) GetAll(kind string, productType string) (*[]models.Product, error) {
	return s.dao.GetAll(kind, productType)
}
