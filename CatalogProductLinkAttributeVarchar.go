package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductLinkAttributeVarchar :
type CatalogProductLinkAttributeVarchar struct {
	ValueId                uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	ProductLinkAttributeId uint16 `gorm:"column:product_link_attribute_id" form:"product_link_attribute_id" json:"product_link_attribute_id"`
	LinkId                 uint16 `gorm:"column:link_id" form:"link_id" json:"link_id"`
	Value                  string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductLinkAttributeVarchar) TableName() string {
	return "catalog_product_link_attribute_varchar"
}

// handler create
func initRoutersCatalogProductLinkAttributeVarchar(r *gin.Engine, catalogproductlinkattributevarchar string) {
	route := r.Group("catalogproductlinkattributevarchar")
	route.GET("/", getCatalogProductLinkAttributeVarchars)
	route.GET("/:id", getCatalogProductLinkAttributeVarchar)
	route.POST("/", createCatalogProductLinkAttributeVarchar)
	route.PUT("/:id", updateCatalogProductLinkAttributeVarchar)
	route.DELETE("/:id", deleteCatalogProductLinkAttributeVarchar)
}

func getCatalogProductLinkAttributeVarchars(c *gin.Context) {
	var catalogProductLinkAttributeVarchars []CatalogProductLinkAttributeVarchar
	if err := g.Find(&catalogProductLinkAttributeVarchars).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinkAttributeVarchars)
	}
}

func getCatalogProductLinkAttributeVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLinkAttributeVarchar CatalogProductLinkAttributeVarchar
	if err := g.Where("id = ?", id).First(&catalogProductLinkAttributeVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinkAttributeVarchar)
	}
}

func createCatalogProductLinkAttributeVarchar(c *gin.Context) {
	var catalogProductLinkAttributeVarchar CatalogProductLinkAttributeVarchar
	c.BindJSON(&catalogProductLinkAttributeVarchar)
	g.Create(&catalogProductLinkAttributeVarchar)
	c.JSON(200, catalogProductLinkAttributeVarchar)
}

func updateCatalogProductLinkAttributeVarchar(c *gin.Context) {
	var catalogProductLinkAttributeVarchar CatalogProductLinkAttributeVarchar
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductLinkAttributeVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductLinkAttributeVarchar)
	g.Save(&catalogProductLinkAttributeVarchar)
	c.JSON(200, catalogProductLinkAttributeVarchar)
}
func deleteCatalogProductLinkAttributeVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLinkAttributeVarchar CatalogProductLinkAttributeVarchar
	d := g.Where("id = ?", id).Delete(&catalogProductLinkAttributeVarchar)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
