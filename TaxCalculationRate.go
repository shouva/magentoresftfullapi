package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TaxCalculationRate :
type TaxCalculationRate struct {
	TaxCalculationRateId uint16  `gorm:"column:tax_calculation_rate_id;primary_key" form:"tax_calculation_rate_id;primary_key" json:"tax_calculation_rate_id;primary_key"`
	TaxCountryId         string  `gorm:"column:tax_country_id" form:"tax_country_id" json:"tax_country_id"`
	TaxRegionId          uint16  `gorm:"column:tax_region_id" form:"tax_region_id" json:"tax_region_id"`
	TaxPostcode          string  `gorm:"column:tax_postcode" form:"tax_postcode" json:"tax_postcode"`
	Code                 string  `gorm:"column:code" form:"code" json:"code"`
	Rate                 float64 `gorm:"column:rate" form:"rate" json:"rate"`
	ZipIsRange           uint16  `gorm:"column:zip_is_range" form:"zip_is_range" json:"zip_is_range"`
	ZipFrom              uint16  `gorm:"column:zip_from" form:"zip_from" json:"zip_from"`
	ZipTo                uint16  `gorm:"column:zip_to" form:"zip_to" json:"zip_to"`
}

// TableName :
func (*TaxCalculationRate) TableName() string {
	return "tax_calculation_rate"
}

// handler create
func initRoutersTaxCalculationRate(r *gin.Engine, taxcalculationrate string) {
	route := r.Group("taxcalculationrate")
	route.GET("/", getTaxCalculationRates)
	route.GET("/:id", getTaxCalculationRate)
	route.POST("/", createTaxCalculationRate)
	route.PUT("/:id", updateTaxCalculationRate)
	route.DELETE("/:id", deleteTaxCalculationRate)
}

func getTaxCalculationRates(c *gin.Context) {
	var taxCalculationRates []TaxCalculationRate
	if err := g.Find(&taxCalculationRates).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxCalculationRates)
	}
}

func getTaxCalculationRate(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxCalculationRate TaxCalculationRate
	if err := g.Where("id = ?", id).First(&taxCalculationRate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxCalculationRate)
	}
}

func createTaxCalculationRate(c *gin.Context) {
	var taxCalculationRate TaxCalculationRate
	c.BindJSON(&taxCalculationRate)
	g.Create(&taxCalculationRate)
	c.JSON(200, taxCalculationRate)
}

func updateTaxCalculationRate(c *gin.Context) {
	var taxCalculationRate TaxCalculationRate
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&taxCalculationRate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&taxCalculationRate)
	g.Save(&taxCalculationRate)
	c.JSON(200, taxCalculationRate)
}
func deleteTaxCalculationRate(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxCalculationRate TaxCalculationRate
	d := g.Where("id = ?", id).Delete(&taxCalculationRate)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
