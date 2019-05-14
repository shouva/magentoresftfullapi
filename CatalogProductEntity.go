package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntity :
type CatalogProductEntity struct {
	EntityId        uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	EntityTypeId    uint16     `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeSetId  uint16     `gorm:"column:attribute_set_id" form:"attribute_set_id" json:"attribute_set_id"`
	TypeId          string     `gorm:"column:type_id" form:"type_id" json:"type_id"`
	Sku             string     `gorm:"column:sku" form:"sku" json:"sku"`
	HasOptions      uint16     `gorm:"column:has_options" form:"has_options" json:"has_options"`
	RequiredOptions uint16     `gorm:"column:required_options" form:"required_options" json:"required_options"`
	CreatedAt       *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
}

// TableName :
func (*CatalogProductEntity) TableName() string {
	return "catalog_product_entity"
}

// handler create
func initRoutersCatalogProductEntity(r *gin.Engine, catalogproductentity string) {
	route := r.Group("catalogproductentity")
	route.GET("/", getCatalogProductEntitys)
	route.GET("/:id", getCatalogProductEntity)
	route.POST("/", createCatalogProductEntity)
	route.PUT("/:id", updateCatalogProductEntity)
	route.DELETE("/:id", deleteCatalogProductEntity)
}

func getCatalogProductEntitys(c *gin.Context) {
	var catalogProductEntitys []CatalogProductEntity
	if err := g.Find(&catalogProductEntitys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntitys)
	}
}

func getCatalogProductEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntity CatalogProductEntity
	if err := g.Where("id = ?", id).First(&catalogProductEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntity)
	}
}

func createCatalogProductEntity(c *gin.Context) {
	var catalogProductEntity CatalogProductEntity
	c.BindJSON(&catalogProductEntity)
	g.Create(&catalogProductEntity)
	c.JSON(200, catalogProductEntity)
}

func updateCatalogProductEntity(c *gin.Context) {
	var catalogProductEntity CatalogProductEntity
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntity)
	g.Save(&catalogProductEntity)
	c.JSON(200, catalogProductEntity)
}
func deleteCatalogProductEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntity CatalogProductEntity
	d := g.Where("id = ?", id).Delete(&catalogProductEntity)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
