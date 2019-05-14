package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntityVarchar :
type CatalogProductEntityVarchar struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductEntityVarchar) TableName() string {
	return "catalog_product_entity_varchar"
}

// handler create
func initRoutersCatalogProductEntityVarchar(r *gin.Engine, catalogproductentityvarchar string) {
	route := r.Group("catalogproductentityvarchar")
	route.GET("/", getCatalogProductEntityVarchars)
	route.GET("/:id", getCatalogProductEntityVarchar)
	route.POST("/", createCatalogProductEntityVarchar)
	route.PUT("/:id", updateCatalogProductEntityVarchar)
	route.DELETE("/:id", deleteCatalogProductEntityVarchar)
}

func getCatalogProductEntityVarchars(c *gin.Context) {
	var catalogProductEntityVarchars []CatalogProductEntityVarchar
	if err := g.Find(&catalogProductEntityVarchars).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityVarchars)
	}
}

func getCatalogProductEntityVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityVarchar CatalogProductEntityVarchar
	if err := g.Where("id = ?", id).First(&catalogProductEntityVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityVarchar)
	}
}

func createCatalogProductEntityVarchar(c *gin.Context) {
	var catalogProductEntityVarchar CatalogProductEntityVarchar
	c.BindJSON(&catalogProductEntityVarchar)
	g.Create(&catalogProductEntityVarchar)
	c.JSON(200, catalogProductEntityVarchar)
}

func updateCatalogProductEntityVarchar(c *gin.Context) {
	var catalogProductEntityVarchar CatalogProductEntityVarchar
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntityVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntityVarchar)
	g.Save(&catalogProductEntityVarchar)
	c.JSON(200, catalogProductEntityVarchar)
}
func deleteCatalogProductEntityVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityVarchar CatalogProductEntityVarchar
	d := g.Where("id = ?", id).Delete(&catalogProductEntityVarchar)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
