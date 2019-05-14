package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TaxCalculationRule :
type TaxCalculationRule struct {
	TaxCalculationRuleId uint16 `gorm:"column:tax_calculation_rule_id;primary_key" form:"tax_calculation_rule_id;primary_key" json:"tax_calculation_rule_id;primary_key"`
	Code                 string `gorm:"column:code" form:"code" json:"code"`
	Priority             uint16 `gorm:"column:priority" form:"priority" json:"priority"`
	Position             uint16 `gorm:"column:position" form:"position" json:"position"`
	CalculateSubtotal    uint16 `gorm:"column:calculate_subtotal" form:"calculate_subtotal" json:"calculate_subtotal"`
}

// TableName :
func (*TaxCalculationRule) TableName() string {
	return "tax_calculation_rule"
}

// handler create
func initRoutersTaxCalculationRule(r *gin.Engine, taxcalculationrule string) {
	route := r.Group("taxcalculationrule")
	route.GET("/", getTaxCalculationRules)
	route.GET("/:id", getTaxCalculationRule)
	route.POST("/", createTaxCalculationRule)
	route.PUT("/:id", updateTaxCalculationRule)
	route.DELETE("/:id", deleteTaxCalculationRule)
}

func getTaxCalculationRules(c *gin.Context) {
	var taxCalculationRules []TaxCalculationRule
	if err := g.Find(&taxCalculationRules).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxCalculationRules)
	}
}

func getTaxCalculationRule(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxCalculationRule TaxCalculationRule
	if err := g.Where("id = ?", id).First(&taxCalculationRule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxCalculationRule)
	}
}

func createTaxCalculationRule(c *gin.Context) {
	var taxCalculationRule TaxCalculationRule
	c.BindJSON(&taxCalculationRule)
	g.Create(&taxCalculationRule)
	c.JSON(200, taxCalculationRule)
}

func updateTaxCalculationRule(c *gin.Context) {
	var taxCalculationRule TaxCalculationRule
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&taxCalculationRule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&taxCalculationRule)
	g.Save(&taxCalculationRule)
	c.JSON(200, taxCalculationRule)
}
func deleteTaxCalculationRule(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxCalculationRule TaxCalculationRule
	d := g.Where("id = ?", id).Delete(&taxCalculationRule)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
