package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryEntityInt :
type CatalogCategoryEntityInt struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        uint16 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogCategoryEntityInt) TableName() string {
	return "catalog_category_entity_int"
}

// handler create
func initRoutersCatalogCategoryEntityInt(r *gin.Engine, catalogcategoryentityint string) {
	route := r.Group("catalogcategoryentityint")
	route.GET("/", getCatalogCategoryEntityInts)
	route.GET("/:id", getCatalogCategoryEntityInt)
	route.POST("/", createCatalogCategoryEntityInt)
	route.PUT("/:id", updateCatalogCategoryEntityInt)
	route.DELETE("/:id", deleteCatalogCategoryEntityInt)
}

func getCatalogCategoryEntityInts(c *gin.Context) {
	var catalogCategoryEntityInts []CatalogCategoryEntityInt
	if err := g.Find(&catalogCategoryEntityInts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntityInts)
	}
}

func getCatalogCategoryEntityInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntityInt CatalogCategoryEntityInt
	if err := g.Where("id = ?", id).First(&catalogCategoryEntityInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntityInt)
	}
}

func createCatalogCategoryEntityInt(c *gin.Context) {
	var catalogCategoryEntityInt CatalogCategoryEntityInt
	c.BindJSON(&catalogCategoryEntityInt)
	g.Create(&catalogCategoryEntityInt)
	c.JSON(200, catalogCategoryEntityInt)
}

func updateCatalogCategoryEntityInt(c *gin.Context) {
	var catalogCategoryEntityInt CatalogCategoryEntityInt
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryEntityInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryEntityInt)
	g.Save(&catalogCategoryEntityInt)
	c.JSON(200, catalogCategoryEntityInt)
}
func deleteCatalogCategoryEntityInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntityInt CatalogCategoryEntityInt
	d := g.Where("id = ?", id).Delete(&catalogCategoryEntityInt)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
