package view

import (
	"UnitConverter/pkg/handler"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowForm(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func ConvertUnit(c *gin.Context) {
	valueStr := c.PostForm("value")
	fromUnit := c.PostForm("from_unit")
	toUnit := c.PostForm("to_unit")

	fmt.Println("Received values - Value:", valueStr, "From:", fromUnit, "To:", toUnit)

	if valueStr == "" || fromUnit == "" || toUnit == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing form fields"})
		return
	}

	fmt.Println("herer1")
	convertedValue, err := handler.ConvertUnits(valueStr, fromUnit, toUnit)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"Value":    valueStr, // Keep input value
			"FromUnit": fromUnit, // Keep selected from-unit
			"ToUnit":   toUnit,
			"error":    err.Error()})
		return
	}

	fmt.Println("herer2")

	fmt.Println(valueStr, "--", convertedValue)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Value":    valueStr, // Keep input value
		"FromUnit": fromUnit, // Keep selected from-unit
		"ToUnit":   toUnit,   // Keep selected to-unit
		"Result":   convertedValue,
	})
}
