package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorldHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
