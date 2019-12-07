package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/daos"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/services"
	"log"
	"net/http"
)

func GetTypes(c *gin.Context) {
	s := services.NewTypeService(daos.NewCategoryDAO())

	inputCategory := c.Param("category")
	fmt.Println(inputCategory)
	if productTypes, err := s.Get(inputCategory); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, productTypes)
	}
}
