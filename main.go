package main

import (
	"my-app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/v1/markets", controllers.GetMarkets)
	r.GET("/api/v1/markets/:id", controllers.GetMarket)
	r.POST("/api/v1/markets", controllers.CreateMarket)
	r.PUT("/api/v1/markets/:id", controllers.UpdateMarket)
	r.DELETE("/api/v1/markets/:id", controllers.DeleteMarket)

	r.GET("/api/v1/stalls", controllers.GetStalls)
	r.GET("/api/v1/stalls/:id", controllers.GetStall)
	r.POST("/api/v1/stalls", controllers.CreateStall)
	r.PUT("/api/v1/stalls/:id", controllers.UpdateStall)
	r.DELETE("/api/v1/stalls/:id", controllers.DeleteStall)

	// r.GET("/api/v1/stall-prices", controllers.GetStallPrices)
	// r.GET("/api/v1/stall-prices/:id", controllers.GetStallPrice)
	// r.POST("/api/v1/stall-prices", controllers.CreateStallPrice)
	// r.PUT("/api/v1/stall-prices/:id", controllers.UpdateStallPrice)
	// r.DELETE("/api/v1/stall-prices/:id", controllers.DeleteStallPrice)

	r.Run()
}
