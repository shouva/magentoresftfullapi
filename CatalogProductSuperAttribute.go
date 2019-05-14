package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductSuperAttribute :
type CatalogProductSuperAttribute struct {
	ProductSuperAttributeId uint16 `gorm:"column:product_super_attribute_id;primary_key" form:"product_super_attribute_id;primary_key" json:"product_super_attribute_id;primary_key"`
	ProductId               uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	AttributeId             uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	Position                uint16 `gorm:"column:position" form:"position" json:"position"`
}

// TableName :
func (*CatalogProductSuperAttribute) TableName() string {
	return "catalog_product_super_attribute"
}

// handler create
func initRoutersCatalogProductSuperAttribute(r *gin.Engine, catalogproductsuperattribute string) {
	route := r.Group("catalogproductsuperattribute")
	route.GET("/", getCatalogProductSuperAttributes)
	route.GET("/:id", getCatalogProductSuperAttribute)
	route.POST("/", createCatalogProductSuperAttribute)
	route.PUT("/:id", updateCatalogProductSuperAttribute)
	route.DELETE("/:id", deleteCatalogProductSuperAttribute)
}

func getCatalogProductSuperAttributes(c *gin.Context) {
	var catalogProductSuperAttributes []CatalogProductSuperAttribute
	if err := g.Find(&catalogProductSuperAttributes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductSuperAttributes)
	}
}

func getCatalogProductSuperAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductSuperAttribute CatalogProductSuperAttribute
	if err := g.Where("id = ?", id).First(&catalogProductSuperAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductSuperAttribute)
	}
}

func createCatalogProductSuperAttribute(c *gin.Context) {
	var catalogProductSuperAttribute CatalogProductSuperAttribute
	c.BindJSON(&catalogProductSuperAttribute)
	g.Create(&catalogProductSuperAttribute)
	c.JSON(200, catalogProductSuperAttribute)
}

func updateCatalogProductSuperAttribute(c *gin.Context) {
	var catalogProductSuperAttribute CatalogProductSuperAttribute
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductSuperAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductSuperAttribute)
	g.Save(&catalogProductSuperAttribute)
	c.JSON(200, catalogProductSuperAttribute)
}
func deleteCatalogProductSuperAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductSuperAttribute CatalogProductSuperAttribute
	d := g.Where("id = ?", id).Delete(&catalogProductSuperAttribute)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
