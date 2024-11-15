package controllers

import (
	"my-app/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var markets = []models.Market{}
var marketIDCounter = 1

func GetMarkets(c *gin.Context) {
	c.JSON(http.StatusOK, markets)
}

func GetMarket(c *gin.Context) {
	id := c.Param("id")
	for _, market := range markets {
		if strings.EqualFold(market.ID, id) {
			c.JSON(http.StatusOK, market)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Market not found"})
}

func CreateMarket(c *gin.Context) {
	var market models.Market
	if err := c.ShouldBindJSON(&market); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if market.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Market name cannot be empty"})
		return
	}

	if len(market.Name) > 255 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Market name is too long"})
		return
	}

	for _, m := range markets {
		if strings.EqualFold(m.Name, market.Name) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Market name must be unique"})
			return
		}
	}

	market.ID = "MKT" + strconv.Itoa(marketIDCounter)
	marketIDCounter++
	markets = append(markets, market)
	c.JSON(http.StatusCreated, market)
}

func UpdateMarket(c *gin.Context) {
	id := c.Param("id")
	for i, market := range markets {
		if strings.EqualFold(market.ID, id) {
			var updatedMarket models.Market
			if err := c.ShouldBindJSON(&updatedMarket); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if strings.EqualFold(market.Name, updatedMarket.Name) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot update,Name is the same as the old name"})
				return
			}

			for _, m := range markets {
				if strings.EqualFold(m.Name, updatedMarket.Name) && !strings.EqualFold(m.ID, id) {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Market name must be unique"})
					return
				}
			}

			markets[i].Name = updatedMarket.Name
			c.JSON(http.StatusOK, markets[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Market not found"})
}

func DeleteMarket(c *gin.Context) {
	id := c.Param("id")
	for i, market := range markets {
		if strings.EqualFold(market.ID, id) {
			markets = append(markets[:i], markets[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Market deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Market not found"})
}
