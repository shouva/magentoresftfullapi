package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryEntityDecimal :
type CatalogCategoryEntityDecimal struct {
	ValueId      uint16  `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16  `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16  `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16  `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16  `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        float64 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogCategoryEntityDecimal) TableName() string {
	return "catalog_category_entity_decimal"
}

// handler create
func initRoutersCatalogCategoryEntityDecimal(r *gin.Engine, catalogcategoryentitydecimal string) {
	route := r.Group("catalogcategoryentitydecimal")
	route.GET("/", getCatalogCategoryEntityDecimals)
	route.GET("/:id", getCatalogCategoryEntityDecimal)
	route.POST("/", createCatalogCategoryEntityDecimal)
	route.PUT("/:id", updateCatalogCategoryEntityDecimal)
	route.DELETE("/:id", deleteCatalogCategoryEntityDecimal)
}

func getCatalogCategoryEntityDecimals(c *gin.Context) {
	var catalogCategoryEntityDecimals []CatalogCategoryEntityDecimal
	if err := g.Find(&catalogCategoryEntityDecimals).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntityDecimals)
	}
}

func getCatalogCategoryEntityDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntityDecimal CatalogCategoryEntityDecimal
	if err := g.Where("id = ?", id).First(&catalogCategoryEntityDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntityDecimal)
	}
}

func createCatalogCategoryEntityDecimal(c *gin.Context) {
	var catalogCategoryEntityDecimal CatalogCategoryEntityDecimal
	c.BindJSON(&catalogCategoryEntityDecimal)
	g.Create(&catalogCategoryEntityDecimal)
	c.JSON(200, catalogCategoryEntityDecimal)
}

func updateCatalogCategoryEntityDecimal(c *gin.Context) {
	var catalogCategoryEntityDecimal CatalogCategoryEntityDecimal
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryEntityDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryEntityDecimal)
	g.Save(&catalogCategoryEntityDecimal)
	c.JSON(200, catalogCategoryEntityDecimal)
}
func deleteCatalogCategoryEntityDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntityDecimal CatalogCategoryEntityDecimal
	d := g.Where("id = ?", id).Delete(&catalogCategoryEntityDecimal)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
