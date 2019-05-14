package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductLinkAttributeDecimal :
type CatalogProductLinkAttributeDecimal struct {
	ValueId                uint16  `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	ProductLinkAttributeId uint16  `gorm:"column:product_link_attribute_id" form:"product_link_attribute_id" json:"product_link_attribute_id"`
	LinkId                 uint16  `gorm:"column:link_id" form:"link_id" json:"link_id"`
	Value                  float64 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductLinkAttributeDecimal) TableName() string {
	return "catalog_product_link_attribute_decimal"
}

// handler create
func initRoutersCatalogProductLinkAttributeDecimal(r *gin.Engine, catalogproductlinkattributedecimal string) {
	route := r.Group("catalogproductlinkattributedecimal")
	route.GET("/", getCatalogProductLinkAttributeDecimals)
	route.GET("/:id", getCatalogProductLinkAttributeDecimal)
	route.POST("/", createCatalogProductLinkAttributeDecimal)
	route.PUT("/:id", updateCatalogProductLinkAttributeDecimal)
	route.DELETE("/:id", deleteCatalogProductLinkAttributeDecimal)
}

func getCatalogProductLinkAttributeDecimals(c *gin.Context) {
	var catalogProductLinkAttributeDecimals []CatalogProductLinkAttributeDecimal
	if err := g.Find(&catalogProductLinkAttributeDecimals).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinkAttributeDecimals)
	}
}

func getCatalogProductLinkAttributeDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLinkAttributeDecimal CatalogProductLinkAttributeDecimal
	if err := g.Where("id = ?", id).First(&catalogProductLinkAttributeDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinkAttributeDecimal)
	}
}

func createCatalogProductLinkAttributeDecimal(c *gin.Context) {
	var catalogProductLinkAttributeDecimal CatalogProductLinkAttributeDecimal
	c.BindJSON(&catalogProductLinkAttributeDecimal)
	g.Create(&catalogProductLinkAttributeDecimal)
	c.JSON(200, catalogProductLinkAttributeDecimal)
}

func updateCatalogProductLinkAttributeDecimal(c *gin.Context) {
	var catalogProductLinkAttributeDecimal CatalogProductLinkAttributeDecimal
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductLinkAttributeDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductLinkAttributeDecimal)
	g.Save(&catalogProductLinkAttributeDecimal)
	c.JSON(200, catalogProductLinkAttributeDecimal)
}
func deleteCatalogProductLinkAttributeDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLinkAttributeDecimal CatalogProductLinkAttributeDecimal
	d := g.Where("id = ?", id).Delete(&catalogProductLinkAttributeDecimal)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
