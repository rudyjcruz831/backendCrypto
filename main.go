package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/backendCrypto/controllers"
)

// create the routes using GIN web framework
func main() {
	// grabing port from env for running server local or on heroku
	port := os.Getenv("PORT")
	// host := os.Getenv("host")
	// start the GIN route
	// fmt.Println("port", port)
	router := gin.Default()
	// router for cors to be able to access from react
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080", "*"},
		AllowMethods: []string{"POST", "GET", "UPDATE", "DELETE", "PUT"},
	}))

	// routes that get created in the controller folder
	router.GET("/", controllers.GetSomething)
	router.GET("/info", controllers.GetCryptoInfo)
	router.GET("/best", controllers.GetBestPrices)
	// fmt.Println("something:", os.Getenv("PORT"))

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
