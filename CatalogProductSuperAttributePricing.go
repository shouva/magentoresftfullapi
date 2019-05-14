package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductSuperAttributePricing :
type CatalogProductSuperAttributePricing struct {
	ValueId                 uint16  `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	ProductSuperAttributeId uint16  `gorm:"column:product_super_attribute_id" form:"product_super_attribute_id" json:"product_super_attribute_id"`
	ValueIndex              string  `gorm:"column:value_index" form:"value_index" json:"value_index"`
	IsPercent               uint16  `gorm:"column:is_percent" form:"is_percent" json:"is_percent"`
	PricingValue            float64 `gorm:"column:pricing_value" form:"pricing_value" json:"pricing_value"`
	WebsiteId               uint16  `gorm:"column:website_id" form:"website_id" json:"website_id"`
}

// TableName :
func (*CatalogProductSuperAttributePricing) TableName() string {
	return "catalog_product_super_attribute_pricing"
}

// handler create
func initRoutersCatalogProductSuperAttributePricing(r *gin.Engine, catalogproductsuperattributepricing string) {
	route := r.Group("catalogproductsuperattributepricing")
	route.GET("/", getCatalogProductSuperAttributePricings)
	route.GET("/:id", getCatalogProductSuperAttributePricing)
	route.POST("/", createCatalogProductSuperAttributePricing)
	route.PUT("/:id", updateCatalogProductSuperAttributePricing)
	route.DELETE("/:id", deleteCatalogProductSuperAttributePricing)
}

func getCatalogProductSuperAttributePricings(c *gin.Context) {
	var catalogProductSuperAttributePricings []CatalogProductSuperAttributePricing
	if err := g.Find(&catalogProductSuperAttributePricings).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductSuperAttributePricings)
	}
}

func getCatalogProductSuperAttributePricing(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductSuperAttributePricing CatalogProductSuperAttributePricing
	if err := g.Where("id = ?", id).First(&catalogProductSuperAttributePricing).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductSuperAttributePricing)
	}
}

func createCatalogProductSuperAttributePricing(c *gin.Context) {
	var catalogProductSuperAttributePricing CatalogProductSuperAttributePricing
	c.BindJSON(&catalogProductSuperAttributePricing)
	g.Create(&catalogProductSuperAttributePricing)
	c.JSON(200, catalogProductSuperAttributePricing)
}

func updateCatalogProductSuperAttributePricing(c *gin.Context) {
	var catalogProductSuperAttributePricing CatalogProductSuperAttributePricing
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductSuperAttributePricing).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductSuperAttributePricing)
	g.Save(&catalogProductSuperAttributePricing)
	c.JSON(200, catalogProductSuperAttributePricing)
}
func deleteCatalogProductSuperAttributePricing(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductSuperAttributePricing CatalogProductSuperAttributePricing
	d := g.Where("id = ?", id).Delete(&catalogProductSuperAttributePricing)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
