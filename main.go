package main

import (
	"country_holidays/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHolidays(c *gin.Context) {
	year := c.Param("id")
	countryCode := c.Param("countryCode")

	holidays, err := handlers.GetHolidaysHandler(year, countryCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch holidays"})
		return
	}
	c.IndentedJSON(http.StatusOK, holidays)
}

func main() {
	router := gin.Default()

	router.GET("/holidays/:id/:countryCode", getHolidays)

	router.Run(":8080")
}
