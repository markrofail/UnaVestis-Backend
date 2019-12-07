package services

import "github.com/markrofail/fashion_scraping_api/cmd/unaVestis/models"

type typeDAO interface {
	Get(category string) ([]*models.Type, error)
}

type TypeService struct {
	dao typeDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewTypeService(dao typeDAO) *TypeService {
	return &TypeService{dao}
}

// Get just retrieves user using User DAO, here can be additional logic for processing data retrieved by DAOs
func (s *TypeService) Get(category string) ([]*models.Type, error) {
	return s.dao.Get(category)
}
