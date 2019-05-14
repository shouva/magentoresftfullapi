package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductWebsite :
type CatalogProductWebsite struct {
	ProductId uint16 `gorm:"column:product_id;primary_key" form:"product_id;primary_key" json:"product_id;primary_key"`
	WebsiteId uint16 `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
}

// TableName :
func (*CatalogProductWebsite) TableName() string {
	return "catalog_product_website"
}

// handler create
func initRoutersCatalogProductWebsite(r *gin.Engine, catalogproductwebsite string) {
	route := r.Group("catalogproductwebsite")
	route.GET("/", getCatalogProductWebsites)
	route.GET("/:id", getCatalogProductWebsite)
	route.POST("/", createCatalogProductWebsite)
	route.PUT("/:id", updateCatalogProductWebsite)
	route.DELETE("/:id", deleteCatalogProductWebsite)
}

func getCatalogProductWebsites(c *gin.Context) {
	var catalogProductWebsites []CatalogProductWebsite
	if err := g.Find(&catalogProductWebsites).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductWebsites)
	}
}

func getCatalogProductWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductWebsite CatalogProductWebsite
	if err := g.Where("id = ?", id).First(&catalogProductWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductWebsite)
	}
}

func createCatalogProductWebsite(c *gin.Context) {
	var catalogProductWebsite CatalogProductWebsite
	c.BindJSON(&catalogProductWebsite)
	g.Create(&catalogProductWebsite)
	c.JSON(200, catalogProductWebsite)
}

func updateCatalogProductWebsite(c *gin.Context) {
	var catalogProductWebsite CatalogProductWebsite
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductWebsite)
	g.Save(&catalogProductWebsite)
	c.JSON(200, catalogProductWebsite)
}
func deleteCatalogProductWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductWebsite CatalogProductWebsite
	d := g.Where("id = ?", id).Delete(&catalogProductWebsite)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
