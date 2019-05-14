package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryEntityDatetime :
type CatalogCategoryEntityDatetime struct {
	ValueId      uint16     `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16     `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16     `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16     `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        *time.Time `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogCategoryEntityDatetime) TableName() string {
	return "catalog_category_entity_datetime"
}

// handler create
func initRoutersCatalogCategoryEntityDatetime(r *gin.Engine, catalogcategoryentitydatetime string) {
	route := r.Group("catalogcategoryentitydatetime")
	route.GET("/", getCatalogCategoryEntityDatetimes)
	route.GET("/:id", getCatalogCategoryEntityDatetime)
	route.POST("/", createCatalogCategoryEntityDatetime)
	route.PUT("/:id", updateCatalogCategoryEntityDatetime)
	route.DELETE("/:id", deleteCatalogCategoryEntityDatetime)
}

func getCatalogCategoryEntityDatetimes(c *gin.Context) {
	var catalogCategoryEntityDatetimes []CatalogCategoryEntityDatetime
	if err := g.Find(&catalogCategoryEntityDatetimes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntityDatetimes)
	}
}

func getCatalogCategoryEntityDatetime(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntityDatetime CatalogCategoryEntityDatetime
	if err := g.Where("id = ?", id).First(&catalogCategoryEntityDatetime).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntityDatetime)
	}
}

func createCatalogCategoryEntityDatetime(c *gin.Context) {
	var catalogCategoryEntityDatetime CatalogCategoryEntityDatetime
	c.BindJSON(&catalogCategoryEntityDatetime)
	g.Create(&catalogCategoryEntityDatetime)
	c.JSON(200, catalogCategoryEntityDatetime)
}

func updateCatalogCategoryEntityDatetime(c *gin.Context) {
	var catalogCategoryEntityDatetime CatalogCategoryEntityDatetime
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryEntityDatetime).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryEntityDatetime)
	g.Save(&catalogCategoryEntityDatetime)
	c.JSON(200, catalogCategoryEntityDatetime)
}
func deleteCatalogCategoryEntityDatetime(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntityDatetime CatalogCategoryEntityDatetime
	d := g.Where("id = ?", id).Delete(&catalogCategoryEntityDatetime)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
