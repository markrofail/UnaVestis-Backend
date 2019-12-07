package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/apis"
	"github.com/markrofail/fashion_scraping_api/cmd/unaVestis/config"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https:///blob/master/LICENSE

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		//v1.Use(auth())
		v1.GET("/products", apis.GetProducts)
		v1.GET("/categories", apis.GetCategories)
		v1.GET("/types/:category", apis.GetTypes)
		v1.GET("/", apis.HelloWorldHandler)
	}

	_ = r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}

//func auth() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		authHeader := c.GetHeader("Authorization")
//		if len(authHeader) == 0 {
//			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
//			c.Abort()
//		}
//		if authHeader != config.Config.ApiKey {
//			httputil.NewError(c, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to this operation: api_key=%s", authHeader))
//			c.Abort()
//		}
//		c.Next()
//	}
//}
