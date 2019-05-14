package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntityGallery :
type CatalogProductEntityGallery struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Position     uint16 `gorm:"column:position" form:"position" json:"position"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductEntityGallery) TableName() string {
	return "catalog_product_entity_gallery"
}

// handler create
func initRoutersCatalogProductEntityGallery(r *gin.Engine, catalogproductentitygallery string) {
	route := r.Group("catalogproductentitygallery")
	route.GET("/", getCatalogProductEntityGallerys)
	route.GET("/:id", getCatalogProductEntityGallery)
	route.POST("/", createCatalogProductEntityGallery)
	route.PUT("/:id", updateCatalogProductEntityGallery)
	route.DELETE("/:id", deleteCatalogProductEntityGallery)
}

func getCatalogProductEntityGallerys(c *gin.Context) {
	var catalogProductEntityGallerys []CatalogProductEntityGallery
	if err := g.Find(&catalogProductEntityGallerys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityGallerys)
	}
}

func getCatalogProductEntityGallery(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityGallery CatalogProductEntityGallery
	if err := g.Where("id = ?", id).First(&catalogProductEntityGallery).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityGallery)
	}
}

func createCatalogProductEntityGallery(c *gin.Context) {
	var catalogProductEntityGallery CatalogProductEntityGallery
	c.BindJSON(&catalogProductEntityGallery)
	g.Create(&catalogProductEntityGallery)
	c.JSON(200, catalogProductEntityGallery)
}

func updateCatalogProductEntityGallery(c *gin.Context) {
	var catalogProductEntityGallery CatalogProductEntityGallery
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntityGallery).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntityGallery)
	g.Save(&catalogProductEntityGallery)
	c.JSON(200, catalogProductEntityGallery)
}
func deleteCatalogProductEntityGallery(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityGallery CatalogProductEntityGallery
	d := g.Where("id = ?", id).Delete(&catalogProductEntityGallery)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
