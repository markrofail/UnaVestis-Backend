package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/daos"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/services"
	"log"
	"net/http"
)

func GetCategories(c *gin.Context) {
	s := services.NewCategoryService(daos.NewCategoryDAO())

	if products, err := s.GetAll(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, products)
	}
}
