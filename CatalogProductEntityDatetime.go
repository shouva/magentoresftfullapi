package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntityDatetime :
type CatalogProductEntityDatetime struct {
	ValueId      uint16     `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16     `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16     `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16     `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        *time.Time `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductEntityDatetime) TableName() string {
	return "catalog_product_entity_datetime"
}

// handler create
func initRoutersCatalogProductEntityDatetime(r *gin.Engine, catalogproductentitydatetime string) {
	route := r.Group("catalogproductentitydatetime")
	route.GET("/", getCatalogProductEntityDatetimes)
	route.GET("/:id", getCatalogProductEntityDatetime)
	route.POST("/", createCatalogProductEntityDatetime)
	route.PUT("/:id", updateCatalogProductEntityDatetime)
	route.DELETE("/:id", deleteCatalogProductEntityDatetime)
}

func getCatalogProductEntityDatetimes(c *gin.Context) {
	var catalogProductEntityDatetimes []CatalogProductEntityDatetime
	if err := g.Find(&catalogProductEntityDatetimes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityDatetimes)
	}
}

func getCatalogProductEntityDatetime(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityDatetime CatalogProductEntityDatetime
	if err := g.Where("id = ?", id).First(&catalogProductEntityDatetime).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityDatetime)
	}
}

func createCatalogProductEntityDatetime(c *gin.Context) {
	var catalogProductEntityDatetime CatalogProductEntityDatetime
	c.BindJSON(&catalogProductEntityDatetime)
	g.Create(&catalogProductEntityDatetime)
	c.JSON(200, catalogProductEntityDatetime)
}

func updateCatalogProductEntityDatetime(c *gin.Context) {
	var catalogProductEntityDatetime CatalogProductEntityDatetime
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntityDatetime).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntityDatetime)
	g.Save(&catalogProductEntityDatetime)
	c.JSON(200, catalogProductEntityDatetime)
}
func deleteCatalogProductEntityDatetime(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityDatetime CatalogProductEntityDatetime
	d := g.Where("id = ?", id).Delete(&catalogProductEntityDatetime)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
