package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntityDecimal :
type CatalogProductEntityDecimal struct {
	ValueId      uint16  `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16  `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16  `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16  `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16  `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        float64 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductEntityDecimal) TableName() string {
	return "catalog_product_entity_decimal"
}

// handler create
func initRoutersCatalogProductEntityDecimal(r *gin.Engine, catalogproductentitydecimal string) {
	route := r.Group("catalogproductentitydecimal")
	route.GET("/", getCatalogProductEntityDecimals)
	route.GET("/:id", getCatalogProductEntityDecimal)
	route.POST("/", createCatalogProductEntityDecimal)
	route.PUT("/:id", updateCatalogProductEntityDecimal)
	route.DELETE("/:id", deleteCatalogProductEntityDecimal)
}

func getCatalogProductEntityDecimals(c *gin.Context) {
	var catalogProductEntityDecimals []CatalogProductEntityDecimal
	if err := g.Find(&catalogProductEntityDecimals).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityDecimals)
	}
}

func getCatalogProductEntityDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityDecimal CatalogProductEntityDecimal
	if err := g.Where("id = ?", id).First(&catalogProductEntityDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityDecimal)
	}
}

func createCatalogProductEntityDecimal(c *gin.Context) {
	var catalogProductEntityDecimal CatalogProductEntityDecimal
	c.BindJSON(&catalogProductEntityDecimal)
	g.Create(&catalogProductEntityDecimal)
	c.JSON(200, catalogProductEntityDecimal)
}

func updateCatalogProductEntityDecimal(c *gin.Context) {
	var catalogProductEntityDecimal CatalogProductEntityDecimal
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntityDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntityDecimal)
	g.Save(&catalogProductEntityDecimal)
	c.JSON(200, catalogProductEntityDecimal)
}
func deleteCatalogProductEntityDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityDecimal CatalogProductEntityDecimal
	d := g.Where("id = ?", id).Delete(&catalogProductEntityDecimal)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
