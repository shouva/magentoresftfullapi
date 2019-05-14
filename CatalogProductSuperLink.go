package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductSuperLink :
type CatalogProductSuperLink struct {
	LinkId    uint16 `gorm:"column:link_id;primary_key" form:"link_id;primary_key" json:"link_id;primary_key"`
	ProductId uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	ParentId  uint16 `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
}

// TableName :
func (*CatalogProductSuperLink) TableName() string {
	return "catalog_product_super_link"
}

// handler create
func initRoutersCatalogProductSuperLink(r *gin.Engine, catalogproductsuperlink string) {
	route := r.Group("catalogproductsuperlink")
	route.GET("/", getCatalogProductSuperLinks)
	route.GET("/:id", getCatalogProductSuperLink)
	route.POST("/", createCatalogProductSuperLink)
	route.PUT("/:id", updateCatalogProductSuperLink)
	route.DELETE("/:id", deleteCatalogProductSuperLink)
}

func getCatalogProductSuperLinks(c *gin.Context) {
	var catalogProductSuperLinks []CatalogProductSuperLink
	if err := g.Find(&catalogProductSuperLinks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductSuperLinks)
	}
}

func getCatalogProductSuperLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductSuperLink CatalogProductSuperLink
	if err := g.Where("id = ?", id).First(&catalogProductSuperLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductSuperLink)
	}
}

func createCatalogProductSuperLink(c *gin.Context) {
	var catalogProductSuperLink CatalogProductSuperLink
	c.BindJSON(&catalogProductSuperLink)
	g.Create(&catalogProductSuperLink)
	c.JSON(200, catalogProductSuperLink)
}

func updateCatalogProductSuperLink(c *gin.Context) {
	var catalogProductSuperLink CatalogProductSuperLink
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductSuperLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductSuperLink)
	g.Save(&catalogProductSuperLink)
	c.JSON(200, catalogProductSuperLink)
}
func deleteCatalogProductSuperLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductSuperLink CatalogProductSuperLink
	d := g.Where("id = ?", id).Delete(&catalogProductSuperLink)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
