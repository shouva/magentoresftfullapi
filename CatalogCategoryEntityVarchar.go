package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryEntityVarchar :
type CatalogCategoryEntityVarchar struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogCategoryEntityVarchar) TableName() string {
	return "catalog_category_entity_varchar"
}

// handler create
func initRoutersCatalogCategoryEntityVarchar(r *gin.Engine, catalogcategoryentityvarchar string) {
	route := r.Group("catalogcategoryentityvarchar")
	route.GET("/", getCatalogCategoryEntityVarchars)
	route.GET("/:id", getCatalogCategoryEntityVarchar)
	route.POST("/", createCatalogCategoryEntityVarchar)
	route.PUT("/:id", updateCatalogCategoryEntityVarchar)
	route.DELETE("/:id", deleteCatalogCategoryEntityVarchar)
}

func getCatalogCategoryEntityVarchars(c *gin.Context) {
	var catalogCategoryEntityVarchars []CatalogCategoryEntityVarchar
	if err := g.Find(&catalogCategoryEntityVarchars).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntityVarchars)
	}
}

func getCatalogCategoryEntityVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntityVarchar CatalogCategoryEntityVarchar
	if err := g.Where("id = ?", id).First(&catalogCategoryEntityVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntityVarchar)
	}
}

func createCatalogCategoryEntityVarchar(c *gin.Context) {
	var catalogCategoryEntityVarchar CatalogCategoryEntityVarchar
	c.BindJSON(&catalogCategoryEntityVarchar)
	g.Create(&catalogCategoryEntityVarchar)
	c.JSON(200, catalogCategoryEntityVarchar)
}

func updateCatalogCategoryEntityVarchar(c *gin.Context) {
	var catalogCategoryEntityVarchar CatalogCategoryEntityVarchar
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryEntityVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryEntityVarchar)
	g.Save(&catalogCategoryEntityVarchar)
	c.JSON(200, catalogCategoryEntityVarchar)
}
func deleteCatalogCategoryEntityVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntityVarchar CatalogCategoryEntityVarchar
	d := g.Where("id = ?", id).Delete(&catalogCategoryEntityVarchar)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
