package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexWebsite :
type CatalogProductIndexWebsite struct {
	WebsiteId   uint16     `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	WebsiteDate *time.Time `gorm:"column:website_date" form:"website_date" json:"website_date"`
	Rate        float32    `gorm:"column:rate" form:"rate" json:"rate"`
}

// TableName :
func (*CatalogProductIndexWebsite) TableName() string {
	return "catalog_product_index_website"
}

// handler create
func initRoutersCatalogProductIndexWebsite(r *gin.Engine, catalogproductindexwebsite string) {
	route := r.Group("catalogproductindexwebsite")
	route.GET("/", getCatalogProductIndexWebsites)
	route.GET("/:id", getCatalogProductIndexWebsite)
	route.POST("/", createCatalogProductIndexWebsite)
	route.PUT("/:id", updateCatalogProductIndexWebsite)
	route.DELETE("/:id", deleteCatalogProductIndexWebsite)
}

func getCatalogProductIndexWebsites(c *gin.Context) {
	var catalogProductIndexWebsites []CatalogProductIndexWebsite
	if err := g.Find(&catalogProductIndexWebsites).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexWebsites)
	}
}

func getCatalogProductIndexWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexWebsite CatalogProductIndexWebsite
	if err := g.Where("id = ?", id).First(&catalogProductIndexWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexWebsite)
	}
}

func createCatalogProductIndexWebsite(c *gin.Context) {
	var catalogProductIndexWebsite CatalogProductIndexWebsite
	c.BindJSON(&catalogProductIndexWebsite)
	g.Create(&catalogProductIndexWebsite)
	c.JSON(200, catalogProductIndexWebsite)
}

func updateCatalogProductIndexWebsite(c *gin.Context) {
	var catalogProductIndexWebsite CatalogProductIndexWebsite
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexWebsite)
	g.Save(&catalogProductIndexWebsite)
	c.JSON(200, catalogProductIndexWebsite)
}
func deleteCatalogProductIndexWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexWebsite CatalogProductIndexWebsite
	d := g.Where("id = ?", id).Delete(&catalogProductIndexWebsite)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
