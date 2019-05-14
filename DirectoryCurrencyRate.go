package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DirectoryCurrencyRate :
type DirectoryCurrencyRate struct {
	CurrencyFrom string  `gorm:"column:currency_from;primary_key" form:"currency_from;primary_key" json:"currency_from;primary_key"`
	CurrencyTo   string  `gorm:"column:currency_to;primary_key" form:"currency_to;primary_key" json:"currency_to;primary_key"`
	Rate         float64 `gorm:"column:rate" form:"rate" json:"rate"`
}

// TableName :
func (*DirectoryCurrencyRate) TableName() string {
	return "directory_currency_rate"
}

// handler create
func initRoutersDirectoryCurrencyRate(r *gin.Engine, directorycurrencyrate string) {
	route := r.Group("directorycurrencyrate")
	route.GET("/", getDirectoryCurrencyRates)
	route.GET("/:id", getDirectoryCurrencyRate)
	route.POST("/", createDirectoryCurrencyRate)
	route.PUT("/:id", updateDirectoryCurrencyRate)
	route.DELETE("/:id", deleteDirectoryCurrencyRate)
}

func getDirectoryCurrencyRates(c *gin.Context) {
	var directoryCurrencyRates []DirectoryCurrencyRate
	if err := g.Find(&directoryCurrencyRates).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, directoryCurrencyRates)
	}
}

func getDirectoryCurrencyRate(c *gin.Context) {
	id := c.Params.ByName("id")
	var directoryCurrencyRate DirectoryCurrencyRate
	if err := g.Where("id = ?", id).First(&directoryCurrencyRate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, directoryCurrencyRate)
	}
}

func createDirectoryCurrencyRate(c *gin.Context) {
	var directoryCurrencyRate DirectoryCurrencyRate
	c.BindJSON(&directoryCurrencyRate)
	g.Create(&directoryCurrencyRate)
	c.JSON(200, directoryCurrencyRate)
}

func updateDirectoryCurrencyRate(c *gin.Context) {
	var directoryCurrencyRate DirectoryCurrencyRate
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&directoryCurrencyRate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&directoryCurrencyRate)
	g.Save(&directoryCurrencyRate)
	c.JSON(200, directoryCurrencyRate)
}
func deleteDirectoryCurrencyRate(c *gin.Context) {
	id := c.Params.ByName("id")
	var directoryCurrencyRate DirectoryCurrencyRate
	d := g.Where("id = ?", id).Delete(&directoryCurrencyRate)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
