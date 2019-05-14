package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntityText :
type CatalogProductEntityText struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductEntityText) TableName() string {
	return "catalog_product_entity_text"
}

// handler create
func initRoutersCatalogProductEntityText(r *gin.Engine, catalogproductentitytext string) {
	route := r.Group("catalogproductentitytext")
	route.GET("/", getCatalogProductEntityTexts)
	route.GET("/:id", getCatalogProductEntityText)
	route.POST("/", createCatalogProductEntityText)
	route.PUT("/:id", updateCatalogProductEntityText)
	route.DELETE("/:id", deleteCatalogProductEntityText)
}

func getCatalogProductEntityTexts(c *gin.Context) {
	var catalogProductEntityTexts []CatalogProductEntityText
	if err := g.Find(&catalogProductEntityTexts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityTexts)
	}
}

func getCatalogProductEntityText(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityText CatalogProductEntityText
	if err := g.Where("id = ?", id).First(&catalogProductEntityText).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityText)
	}
}

func createCatalogProductEntityText(c *gin.Context) {
	var catalogProductEntityText CatalogProductEntityText
	c.BindJSON(&catalogProductEntityText)
	g.Create(&catalogProductEntityText)
	c.JSON(200, catalogProductEntityText)
}

func updateCatalogProductEntityText(c *gin.Context) {
	var catalogProductEntityText CatalogProductEntityText
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntityText).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntityText)
	g.Save(&catalogProductEntityText)
	c.JSON(200, catalogProductEntityText)
}
func deleteCatalogProductEntityText(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityText CatalogProductEntityText
	d := g.Where("id = ?", id).Delete(&catalogProductEntityText)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
