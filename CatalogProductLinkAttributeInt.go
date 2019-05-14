package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductLinkAttributeInt :
type CatalogProductLinkAttributeInt struct {
	ValueId                uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	ProductLinkAttributeId uint16 `gorm:"column:product_link_attribute_id" form:"product_link_attribute_id" json:"product_link_attribute_id"`
	LinkId                 uint16 `gorm:"column:link_id" form:"link_id" json:"link_id"`
	Value                  uint16 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductLinkAttributeInt) TableName() string {
	return "catalog_product_link_attribute_int"
}

// handler create
func initRoutersCatalogProductLinkAttributeInt(r *gin.Engine, catalogproductlinkattributeint string) {
	route := r.Group("catalogproductlinkattributeint")
	route.GET("/", getCatalogProductLinkAttributeInts)
	route.GET("/:id", getCatalogProductLinkAttributeInt)
	route.POST("/", createCatalogProductLinkAttributeInt)
	route.PUT("/:id", updateCatalogProductLinkAttributeInt)
	route.DELETE("/:id", deleteCatalogProductLinkAttributeInt)
}

func getCatalogProductLinkAttributeInts(c *gin.Context) {
	var catalogProductLinkAttributeInts []CatalogProductLinkAttributeInt
	if err := g.Find(&catalogProductLinkAttributeInts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinkAttributeInts)
	}
}

func getCatalogProductLinkAttributeInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLinkAttributeInt CatalogProductLinkAttributeInt
	if err := g.Where("id = ?", id).First(&catalogProductLinkAttributeInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinkAttributeInt)
	}
}

func createCatalogProductLinkAttributeInt(c *gin.Context) {
	var catalogProductLinkAttributeInt CatalogProductLinkAttributeInt
	c.BindJSON(&catalogProductLinkAttributeInt)
	g.Create(&catalogProductLinkAttributeInt)
	c.JSON(200, catalogProductLinkAttributeInt)
}

func updateCatalogProductLinkAttributeInt(c *gin.Context) {
	var catalogProductLinkAttributeInt CatalogProductLinkAttributeInt
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductLinkAttributeInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductLinkAttributeInt)
	g.Save(&catalogProductLinkAttributeInt)
	c.JSON(200, catalogProductLinkAttributeInt)
}
func deleteCatalogProductLinkAttributeInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLinkAttributeInt CatalogProductLinkAttributeInt
	d := g.Where("id = ?", id).Delete(&catalogProductLinkAttributeInt)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
