package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TaxCalculationRateTitle :
type TaxCalculationRateTitle struct {
	TaxCalculationRateTitleId uint16 `gorm:"column:tax_calculation_rate_title_id;primary_key" form:"tax_calculation_rate_title_id;primary_key" json:"tax_calculation_rate_title_id;primary_key"`
	TaxCalculationRateId      uint16 `gorm:"column:tax_calculation_rate_id" form:"tax_calculation_rate_id" json:"tax_calculation_rate_id"`
	StoreId                   uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Value                     string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*TaxCalculationRateTitle) TableName() string {
	return "tax_calculation_rate_title"
}

// handler create
func initRoutersTaxCalculationRateTitle(r *gin.Engine, taxcalculationratetitle string) {
	route := r.Group("taxcalculationratetitle")
	route.GET("/", getTaxCalculationRateTitles)
	route.GET("/:id", getTaxCalculationRateTitle)
	route.POST("/", createTaxCalculationRateTitle)
	route.PUT("/:id", updateTaxCalculationRateTitle)
	route.DELETE("/:id", deleteTaxCalculationRateTitle)
}

func getTaxCalculationRateTitles(c *gin.Context) {
	var taxCalculationRateTitles []TaxCalculationRateTitle
	if err := g.Find(&taxCalculationRateTitles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxCalculationRateTitles)
	}
}

func getTaxCalculationRateTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxCalculationRateTitle TaxCalculationRateTitle
	if err := g.Where("id = ?", id).First(&taxCalculationRateTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxCalculationRateTitle)
	}
}

func createTaxCalculationRateTitle(c *gin.Context) {
	var taxCalculationRateTitle TaxCalculationRateTitle
	c.BindJSON(&taxCalculationRateTitle)
	g.Create(&taxCalculationRateTitle)
	c.JSON(200, taxCalculationRateTitle)
}

func updateTaxCalculationRateTitle(c *gin.Context) {
	var taxCalculationRateTitle TaxCalculationRateTitle
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&taxCalculationRateTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&taxCalculationRateTitle)
	g.Save(&taxCalculationRateTitle)
	c.JSON(200, taxCalculationRateTitle)
}
func deleteTaxCalculationRateTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxCalculationRateTitle TaxCalculationRateTitle
	d := g.Where("id = ?", id).Delete(&taxCalculationRateTitle)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
