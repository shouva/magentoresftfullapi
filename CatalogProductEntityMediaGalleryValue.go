package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntityMediaGalleryValue :
type CatalogProductEntityMediaGalleryValue struct {
	ValueId  uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	StoreId  uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Label    string `gorm:"column:label" form:"label" json:"label"`
	Position uint16 `gorm:"column:position" form:"position" json:"position"`
	Disabled uint16 `gorm:"column:disabled" form:"disabled" json:"disabled"`
}

// TableName :
func (*CatalogProductEntityMediaGalleryValue) TableName() string {
	return "catalog_product_entity_media_gallery_value"
}

// handler create
func initRoutersCatalogProductEntityMediaGalleryValue(r *gin.Engine, catalogproductentitymediagalleryvalue string) {
	route := r.Group("catalogproductentitymediagalleryvalue")
	route.GET("/", getCatalogProductEntityMediaGalleryValues)
	route.GET("/:id", getCatalogProductEntityMediaGalleryValue)
	route.POST("/", createCatalogProductEntityMediaGalleryValue)
	route.PUT("/:id", updateCatalogProductEntityMediaGalleryValue)
	route.DELETE("/:id", deleteCatalogProductEntityMediaGalleryValue)
}

func getCatalogProductEntityMediaGalleryValues(c *gin.Context) {
	var catalogProductEntityMediaGalleryValues []CatalogProductEntityMediaGalleryValue
	if err := g.Find(&catalogProductEntityMediaGalleryValues).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityMediaGalleryValues)
	}
}

func getCatalogProductEntityMediaGalleryValue(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityMediaGalleryValue CatalogProductEntityMediaGalleryValue
	if err := g.Where("id = ?", id).First(&catalogProductEntityMediaGalleryValue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityMediaGalleryValue)
	}
}

func createCatalogProductEntityMediaGalleryValue(c *gin.Context) {
	var catalogProductEntityMediaGalleryValue CatalogProductEntityMediaGalleryValue
	c.BindJSON(&catalogProductEntityMediaGalleryValue)
	g.Create(&catalogProductEntityMediaGalleryValue)
	c.JSON(200, catalogProductEntityMediaGalleryValue)
}

func updateCatalogProductEntityMediaGalleryValue(c *gin.Context) {
	var catalogProductEntityMediaGalleryValue CatalogProductEntityMediaGalleryValue
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntityMediaGalleryValue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntityMediaGalleryValue)
	g.Save(&catalogProductEntityMediaGalleryValue)
	c.JSON(200, catalogProductEntityMediaGalleryValue)
}
func deleteCatalogProductEntityMediaGalleryValue(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityMediaGalleryValue CatalogProductEntityMediaGalleryValue
	d := g.Where("id = ?", id).Delete(&catalogProductEntityMediaGalleryValue)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
