package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/backendCrypto/controllers"
)

// create the routes using GIN web framework
func main() {

	// start the GIN route
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

	// the http route where it starts listing
	port := os.Getenv("PORT")
	// if err != nil {
	// 	log.Fatalf(err)
	// }
	router.Run("localhost:" + port)
}
