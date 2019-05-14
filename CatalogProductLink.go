package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductLink :
type CatalogProductLink struct {
	LinkId          uint16 `gorm:"column:link_id;primary_key" form:"link_id;primary_key" json:"link_id;primary_key"`
	ProductId       uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	LinkedProductId uint16 `gorm:"column:linked_product_id" form:"linked_product_id" json:"linked_product_id"`
	LinkTypeId      uint16 `gorm:"column:link_type_id" form:"link_type_id" json:"link_type_id"`
}

// TableName :
func (*CatalogProductLink) TableName() string {
	return "catalog_product_link"
}

// handler create
func initRoutersCatalogProductLink(r *gin.Engine, catalogproductlink string) {
	route := r.Group("catalogproductlink")
	route.GET("/", getCatalogProductLinks)
	route.GET("/:id", getCatalogProductLink)
	route.POST("/", createCatalogProductLink)
	route.PUT("/:id", updateCatalogProductLink)
	route.DELETE("/:id", deleteCatalogProductLink)
}

func getCatalogProductLinks(c *gin.Context) {
	var catalogProductLinks []CatalogProductLink
	if err := g.Find(&catalogProductLinks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLinks)
	}
}

func getCatalogProductLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLink CatalogProductLink
	if err := g.Where("id = ?", id).First(&catalogProductLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductLink)
	}
}

func createCatalogProductLink(c *gin.Context) {
	var catalogProductLink CatalogProductLink
	c.BindJSON(&catalogProductLink)
	g.Create(&catalogProductLink)
	c.JSON(200, catalogProductLink)
}

func updateCatalogProductLink(c *gin.Context) {
	var catalogProductLink CatalogProductLink
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductLink)
	g.Save(&catalogProductLink)
	c.JSON(200, catalogProductLink)
}
func deleteCatalogProductLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductLink CatalogProductLink
	d := g.Where("id = ?", id).Delete(&catalogProductLink)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
