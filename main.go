package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/backendCrypto/controllers"
)

// create the routes using GIN web framework
func main() {
	// grabing port from env for running server local or on heroku
	port := os.Getenv("PORT")
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

	PrintMemUsage()
	router.Run("localhost:" + port)
	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
