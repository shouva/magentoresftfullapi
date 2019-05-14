package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntityInt :
type CatalogProductEntityInt struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        uint16 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductEntityInt) TableName() string {
	return "catalog_product_entity_int"
}

// handler create
func initRoutersCatalogProductEntityInt(r *gin.Engine, catalogproductentityint string) {
	route := r.Group("catalogproductentityint")
	route.GET("/", getCatalogProductEntityInts)
	route.GET("/:id", getCatalogProductEntityInt)
	route.POST("/", createCatalogProductEntityInt)
	route.PUT("/:id", updateCatalogProductEntityInt)
	route.DELETE("/:id", deleteCatalogProductEntityInt)
}

func getCatalogProductEntityInts(c *gin.Context) {
	var catalogProductEntityInts []CatalogProductEntityInt
	if err := g.Find(&catalogProductEntityInts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityInts)
	}
}

func getCatalogProductEntityInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityInt CatalogProductEntityInt
	if err := g.Where("id = ?", id).First(&catalogProductEntityInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityInt)
	}
}

func createCatalogProductEntityInt(c *gin.Context) {
	var catalogProductEntityInt CatalogProductEntityInt
	c.BindJSON(&catalogProductEntityInt)
	g.Create(&catalogProductEntityInt)
	c.JSON(200, catalogProductEntityInt)
}

func updateCatalogProductEntityInt(c *gin.Context) {
	var catalogProductEntityInt CatalogProductEntityInt
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntityInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntityInt)
	g.Save(&catalogProductEntityInt)
	c.JSON(200, catalogProductEntityInt)
}
func deleteCatalogProductEntityInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityInt CatalogProductEntityInt
	d := g.Where("id = ?", id).Delete(&catalogProductEntityInt)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
