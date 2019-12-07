package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/daos"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/services"
	"log"
	"net/http"
)

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /products/?kind={}&type={} [get]
// @Security ApiKeyAuth
func GetProducts(c *gin.Context) {
	s := services.NewProductService(daos.NewProductDAO())

	inputKind := c.Query("kind")
	inputType := c.Query("type")

	if products, err := s.GetAll(inputKind, inputType); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, products)
	}
}
