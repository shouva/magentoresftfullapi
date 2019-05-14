package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntityMediaGallery :
type CatalogProductEntityMediaGallery struct {
	ValueId     uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	AttributeId uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId    uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value       string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductEntityMediaGallery) TableName() string {
	return "catalog_product_entity_media_gallery"
}

// handler create
func initRoutersCatalogProductEntityMediaGallery(r *gin.Engine, catalogproductentitymediagallery string) {
	route := r.Group("catalogproductentitymediagallery")
	route.GET("/", getCatalogProductEntityMediaGallerys)
	route.GET("/:id", getCatalogProductEntityMediaGallery)
	route.POST("/", createCatalogProductEntityMediaGallery)
	route.PUT("/:id", updateCatalogProductEntityMediaGallery)
	route.DELETE("/:id", deleteCatalogProductEntityMediaGallery)
}

func getCatalogProductEntityMediaGallerys(c *gin.Context) {
	var catalogProductEntityMediaGallerys []CatalogProductEntityMediaGallery
	if err := g.Find(&catalogProductEntityMediaGallerys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityMediaGallerys)
	}
}

func getCatalogProductEntityMediaGallery(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityMediaGallery CatalogProductEntityMediaGallery
	if err := g.Where("id = ?", id).First(&catalogProductEntityMediaGallery).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityMediaGallery)
	}
}

func createCatalogProductEntityMediaGallery(c *gin.Context) {
	var catalogProductEntityMediaGallery CatalogProductEntityMediaGallery
	c.BindJSON(&catalogProductEntityMediaGallery)
	g.Create(&catalogProductEntityMediaGallery)
	c.JSON(200, catalogProductEntityMediaGallery)
}

func updateCatalogProductEntityMediaGallery(c *gin.Context) {
	var catalogProductEntityMediaGallery CatalogProductEntityMediaGallery
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntityMediaGallery).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntityMediaGallery)
	g.Save(&catalogProductEntityMediaGallery)
	c.JSON(200, catalogProductEntityMediaGallery)
}
func deleteCatalogProductEntityMediaGallery(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityMediaGallery CatalogProductEntityMediaGallery
	d := g.Where("id = ?", id).Delete(&catalogProductEntityMediaGallery)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
