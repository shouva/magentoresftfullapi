package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductSuperAttributeLabel :
type CatalogProductSuperAttributeLabel struct {
	ValueId                 uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	ProductSuperAttributeId uint16 `gorm:"column:product_super_attribute_id" form:"product_super_attribute_id" json:"product_super_attribute_id"`
	StoreId                 uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	UseDefault              uint16 `gorm:"column:use_default" form:"use_default" json:"use_default"`
	Value                   string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductSuperAttributeLabel) TableName() string {
	return "catalog_product_super_attribute_label"
}

// handler create
func initRoutersCatalogProductSuperAttributeLabel(r *gin.Engine, catalogproductsuperattributelabel string) {
	route := r.Group("catalogproductsuperattributelabel")
	route.GET("/", getCatalogProductSuperAttributeLabels)
	route.GET("/:id", getCatalogProductSuperAttributeLabel)
	route.POST("/", createCatalogProductSuperAttributeLabel)
	route.PUT("/:id", updateCatalogProductSuperAttributeLabel)
	route.DELETE("/:id", deleteCatalogProductSuperAttributeLabel)
}

func getCatalogProductSuperAttributeLabels(c *gin.Context) {
	var catalogProductSuperAttributeLabels []CatalogProductSuperAttributeLabel
	if err := g.Find(&catalogProductSuperAttributeLabels).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductSuperAttributeLabels)
	}
}

func getCatalogProductSuperAttributeLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductSuperAttributeLabel CatalogProductSuperAttributeLabel
	if err := g.Where("id = ?", id).First(&catalogProductSuperAttributeLabel).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductSuperAttributeLabel)
	}
}

func createCatalogProductSuperAttributeLabel(c *gin.Context) {
	var catalogProductSuperAttributeLabel CatalogProductSuperAttributeLabel
	c.BindJSON(&catalogProductSuperAttributeLabel)
	g.Create(&catalogProductSuperAttributeLabel)
	c.JSON(200, catalogProductSuperAttributeLabel)
}

func updateCatalogProductSuperAttributeLabel(c *gin.Context) {
	var catalogProductSuperAttributeLabel CatalogProductSuperAttributeLabel
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductSuperAttributeLabel).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductSuperAttributeLabel)
	g.Save(&catalogProductSuperAttributeLabel)
	c.JSON(200, catalogProductSuperAttributeLabel)
}
func deleteCatalogProductSuperAttributeLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductSuperAttributeLabel CatalogProductSuperAttributeLabel
	d := g.Where("id = ?", id).Delete(&catalogProductSuperAttributeLabel)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
