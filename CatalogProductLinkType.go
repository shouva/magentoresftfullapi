package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductLinkType :
type CatalogProductLinkType struct {
	LinkTypeId uint16 `gorm:"column:link_type_id;primary_key" form:"link_type_id;primary_key" json:"link_type_id;primary_key"`
	Code       string `gorm:"column:code" form:"code" json:"code"`
}

// TableName :
func (*CatalogProductLinkType) TableName() string {
	return "catalog_product_link_type"
}

// handler create
func initRoutersCatalogProductLinkType(r *gin.Engine, catalogproductlinktype string) {
	route := r.Group("catalogproductlinktype")
	route.GET("/", getCatalogProductLinkTypes)
	route.GET("/:id", getCatalogProductLinkType)
	route.POST("/", createCatalogProductLinkType)
	route.PUT("/:id", updateCatalogProductLinkType)
	route.DELETE("/:id", deleteCatalogProductLinkType)
}

func getCatalogProductLinkTypes(c *gin.Context) {
	var catalogProductLinkTypes []CatalogProductLinkType
	if err := g.Find(&catalogProductLinkTypes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinkTypes)
	}
}

func getCatalogProductLinkType(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLinkType CatalogProductLinkType
	if err := g.Where("id = ?", id).First(&catalogProductLinkType).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinkType)
	}
}

func createCatalogProductLinkType(c *gin.Context) {
	var catalogProductLinkType CatalogProductLinkType
	c.BindJSON(&catalogProductLinkType)
	g.Create(&catalogProductLinkType)
	c.JSON(200, catalogProductLinkType)
}

func updateCatalogProductLinkType(c *gin.Context) {
	var catalogProductLinkType CatalogProductLinkType
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductLinkType).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductLinkType)
	g.Save(&catalogProductLinkType)
	c.JSON(200, catalogProductLinkType)
}
func deleteCatalogProductLinkType(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLinkType CatalogProductLinkType
	d := g.Where("id = ?", id).Delete(&catalogProductLinkType)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
