package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TaxCalculation :
type TaxCalculation struct {
	TaxCalculationId     uint16 `gorm:"column:tax_calculation_id;primary_key" form:"tax_calculation_id;primary_key" json:"tax_calculation_id;primary_key"`
	TaxCalculationRateId uint16 `gorm:"column:tax_calculation_rate_id" form:"tax_calculation_rate_id" json:"tax_calculation_rate_id"`
	TaxCalculationRuleId uint16 `gorm:"column:tax_calculation_rule_id" form:"tax_calculation_rule_id" json:"tax_calculation_rule_id"`
	CustomerTaxClassId   uint16 `gorm:"column:customer_tax_class_id" form:"customer_tax_class_id" json:"customer_tax_class_id"`
	ProductTaxClassId    uint16 `gorm:"column:product_tax_class_id" form:"product_tax_class_id" json:"product_tax_class_id"`
}

// TableName :
func (*TaxCalculation) TableName() string {
	return "tax_calculation"
}

// handler create
func initRoutersTaxCalculation(r *gin.Engine, taxcalculation string) {
	route := r.Group("taxcalculation")
	route.GET("/", getTaxCalculations)
	route.GET("/:id", getTaxCalculation)
	route.POST("/", createTaxCalculation)
	route.PUT("/:id", updateTaxCalculation)
	route.DELETE("/:id", deleteTaxCalculation)
}

func getTaxCalculations(c *gin.Context) {
	var taxCalculations []TaxCalculation
	if err := g.Find(&taxCalculations).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxCalculations)
	}
}

func getTaxCalculation(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxCalculation TaxCalculation
	if err := g.Where("id = ?", id).First(&taxCalculation).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxCalculation)
	}
}

func createTaxCalculation(c *gin.Context) {
	var taxCalculation TaxCalculation
	c.BindJSON(&taxCalculation)
	g.Create(&taxCalculation)
	c.JSON(200, taxCalculation)
}

func updateTaxCalculation(c *gin.Context) {
	var taxCalculation TaxCalculation
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&taxCalculation).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&taxCalculation)
	g.Save(&taxCalculation)
	c.JSON(200, taxCalculation)
}
func deleteTaxCalculation(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxCalculation TaxCalculation
	d := g.Where("id = ?", id).Delete(&taxCalculation)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
