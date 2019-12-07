package services

import "github.com/markrofail/fashion_scraping_api/cmd/unaVestis/models"

type categoryDAO interface {
	GetAll() ([]*models.Category, error)
}

type CategoryService struct {
	dao categoryDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewCategoryService(dao categoryDAO) *CategoryService {
	return &CategoryService{dao}
}

// Get just retrieves user using User DAO, here can be additional logic for processing data retrieved by DAOs
func (s *CategoryService) GetAll() ([]*models.Category, error) {
	return s.dao.GetAll()
}
