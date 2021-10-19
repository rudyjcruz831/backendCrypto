package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/backendCrypto/controllers"
)

func main() {

	router := gin.Default()
	// router.Get
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080", "*"},
		AllowMethods: []string{"POST", "GET", "UPDATE", "DELETE", "PUT"},
		// AllowHeaders: []string{"Authorization", "Content-Type", "user_id"},
	}))
	router.GET("/info", controllers.GetCryptoInfo)
	router.Run("localhost:8080")

}
