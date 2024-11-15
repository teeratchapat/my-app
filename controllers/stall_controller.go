package controllers

import (
	"my-app/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var stalls = []models.Stall{}
var stallIDCounter = 1

func GetStalls(c *gin.Context) {
	c.JSON(http.StatusOK, stalls)
}

func GetStall(c *gin.Context) {
	id := c.Param("id")
	for _, stall := range stalls {
		if strings.EqualFold(stall.ID, id) {
			c.JSON(http.StatusOK, stall)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Stall not found"})
}

func CreateStall(c *gin.Context) {
	var stall models.Stall
	if err := c.ShouldBindJSON(&stall); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if strings.TrimSpace(stall.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stall name cannot be empty"})
		return
	}
	if len(stall.Name) > 255 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stall name is too long"})
		return
	}

	if strings.TrimSpace(stall.MarketID) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Market ID cannot be empty"})
		return
	}

	for _, s := range stalls {
		if strings.EqualFold(s.Name, stall.Name) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Stall name must be unique"})
			return
		}
	}

	stall.ID = "STL" + strconv.Itoa(stallIDCounter)
	stallIDCounter++
	stalls = append(stalls, stall)
	c.JSON(http.StatusCreated, stall)
}

func UpdateStall(c *gin.Context) {
	id := c.Param("id")
	for i, stall := range stalls {
		if strings.EqualFold(stall.ID, id) {
			var updatedStall models.Stall
			if err := c.ShouldBindJSON(&updatedStall); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if strings.TrimSpace(updatedStall.Name) == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Stall name cannot be empty"})
				return
			}
			if len(updatedStall.Name) > 255 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Stall name is too long"})
				return
			}

			if strings.TrimSpace(updatedStall.MarketID) == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Market ID cannot be empty"})
				return
			}

			for _, s := range stalls {
				if strings.EqualFold(s.Name, updatedStall.Name) && s.ID != id {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Stall name must be unique"})
					return
				}
			}

			stalls[i].Name = updatedStall.Name
			stalls[i].MarketID = updatedStall.MarketID
			c.JSON(http.StatusOK, stalls[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Stall not found"})
}

func DeleteStall(c *gin.Context) {
	id := c.Param("id")
	for i, stall := range stalls {
		if strings.EqualFold(stall.ID, id) {
			stalls = append(stalls[:i], stalls[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Stall deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Stall not found"})
}
