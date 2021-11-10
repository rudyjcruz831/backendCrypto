package app

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/backendCrypto/controllers"
)

// create the routes using GIN web framework
func StartApp() {
	// grabing port from env for running server local or on heroku
	port := os.Getenv("PORT")

	router := gin.Default()
	// router for cors to be able to access from react
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET"},
	}))

	// routes that get created in the controller folder
	router.GET("/", controllers.GetSomething)
	router.GET("/info", controllers.GetCryptoInfo)
	router.GET("/best", controllers.GetBestPrices)
	// if port env is empty then make it default 8080
	if port == "" {
		port = "8081"
	}
	// running the server on localhost with given port
	router.Run(":" + port)
}
