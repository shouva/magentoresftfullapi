package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductLinkAttribute :
type CatalogProductLinkAttribute struct {
	ProductLinkAttributeId   uint16 `gorm:"column:product_link_attribute_id;primary_key" form:"product_link_attribute_id;primary_key" json:"product_link_attribute_id;primary_key"`
	LinkTypeId               uint16 `gorm:"column:link_type_id" form:"link_type_id" json:"link_type_id"`
	ProductLinkAttributeCode string `gorm:"column:product_link_attribute_code" form:"product_link_attribute_code" json:"product_link_attribute_code"`
	DataType                 string `gorm:"column:data_type" form:"data_type" json:"data_type"`
}

// TableName :
func (*CatalogProductLinkAttribute) TableName() string {
	return "catalog_product_link_attribute"
}

// handler create
func initRoutersCatalogProductLinkAttribute(r *gin.Engine, catalogproductlinkattribute string) {
	route := r.Group("catalogproductlinkattribute")
	route.GET("/", getCatalogProductLinkAttributes)
	route.GET("/:id", getCatalogProductLinkAttribute)
	route.POST("/", createCatalogProductLinkAttribute)
	route.PUT("/:id", updateCatalogProductLinkAttribute)
	route.DELETE("/:id", deleteCatalogProductLinkAttribute)
}

func getCatalogProductLinkAttributes(c *gin.Context) {
	var catalogProductLinkAttributes []CatalogProductLinkAttribute
	if err := g.Find(&catalogProductLinkAttributes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinkAttributes)
	}
}

func getCatalogProductLinkAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLinkAttribute CatalogProductLinkAttribute
	if err := g.Where("id = ?", id).First(&catalogProductLinkAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinkAttribute)
	}
}

func createCatalogProductLinkAttribute(c *gin.Context) {
	var catalogProductLinkAttribute CatalogProductLinkAttribute
	c.BindJSON(&catalogProductLinkAttribute)
	g.Create(&catalogProductLinkAttribute)
	c.JSON(200, catalogProductLinkAttribute)
}

func updateCatalogProductLinkAttribute(c *gin.Context) {
	var catalogProductLinkAttribute CatalogProductLinkAttribute
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductLinkAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductLinkAttribute)
	g.Save(&catalogProductLinkAttribute)
	c.JSON(200, catalogProductLinkAttribute)
}
func deleteCatalogProductLinkAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLinkAttribute CatalogProductLinkAttribute
	d := g.Where("id = ?", id).Delete(&catalogProductLinkAttribute)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
