package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryEntityText :
type CatalogCategoryEntityText struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogCategoryEntityText) TableName() string {
	return "catalog_category_entity_text"
}

// handler create
func initRoutersCatalogCategoryEntityText(r *gin.Engine, catalogcategoryentitytext string) {
	route := r.Group("catalogcategoryentitytext")
	route.GET("/", getCatalogCategoryEntityTexts)
	route.GET("/:id", getCatalogCategoryEntityText)
	route.POST("/", createCatalogCategoryEntityText)
	route.PUT("/:id", updateCatalogCategoryEntityText)
	route.DELETE("/:id", deleteCatalogCategoryEntityText)
}

func getCatalogCategoryEntityTexts(c *gin.Context) {
	var catalogCategoryEntityTexts []CatalogCategoryEntityText
	if err := g.Find(&catalogCategoryEntityTexts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntityTexts)
	}
}

func getCatalogCategoryEntityText(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntityText CatalogCategoryEntityText
	if err := g.Where("id = ?", id).First(&catalogCategoryEntityText).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntityText)
	}
}

func createCatalogCategoryEntityText(c *gin.Context) {
	var catalogCategoryEntityText CatalogCategoryEntityText
	c.BindJSON(&catalogCategoryEntityText)
	g.Create(&catalogCategoryEntityText)
	c.JSON(200, catalogCategoryEntityText)
}

func updateCatalogCategoryEntityText(c *gin.Context) {
	var catalogCategoryEntityText CatalogCategoryEntityText
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryEntityText).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryEntityText)
	g.Save(&catalogCategoryEntityText)
	c.JSON(200, catalogCategoryEntityText)
}
func deleteCatalogCategoryEntityText(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntityText CatalogCategoryEntityText
	d := g.Where("id = ?", id).Delete(&catalogCategoryEntityText)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
