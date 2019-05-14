package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductOptionTypeValue :
type CatalogProductOptionTypeValue struct {
	OptionTypeId uint16 `gorm:"column:option_type_id;primary_key" form:"option_type_id;primary_key" json:"option_type_id;primary_key"`
	OptionId     uint16 `gorm:"column:option_id" form:"option_id" json:"option_id"`
	Sku          string `gorm:"column:sku" form:"sku" json:"sku"`
	SortOrder    uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
}

// TableName :
func (*CatalogProductOptionTypeValue) TableName() string {
	return "catalog_product_option_type_value"
}

// handler create
func initRoutersCatalogProductOptionTypeValue(r *gin.Engine, catalogproductoptiontypevalue string) {
	route := r.Group("catalogproductoptiontypevalue")
	route.GET("/", getCatalogProductOptionTypeValues)
	route.GET("/:id", getCatalogProductOptionTypeValue)
	route.POST("/", createCatalogProductOptionTypeValue)
	route.PUT("/:id", updateCatalogProductOptionTypeValue)
	route.DELETE("/:id", deleteCatalogProductOptionTypeValue)
}

func getCatalogProductOptionTypeValues(c *gin.Context) {
	var catalogProductOptionTypeValues []CatalogProductOptionTypeValue
	if err := g.Find(&catalogProductOptionTypeValues).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptionTypeValues)
	}
}

func getCatalogProductOptionTypeValue(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOptionTypeValue CatalogProductOptionTypeValue
	if err := g.Where("id = ?", id).First(&catalogProductOptionTypeValue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptionTypeValue)
	}
}

func createCatalogProductOptionTypeValue(c *gin.Context) {
	var catalogProductOptionTypeValue CatalogProductOptionTypeValue
	c.BindJSON(&catalogProductOptionTypeValue)
	g.Create(&catalogProductOptionTypeValue)
	c.JSON(200, catalogProductOptionTypeValue)
}

func updateCatalogProductOptionTypeValue(c *gin.Context) {
	var catalogProductOptionTypeValue CatalogProductOptionTypeValue
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductOptionTypeValue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductOptionTypeValue)
	g.Save(&catalogProductOptionTypeValue)
	c.JSON(200, catalogProductOptionTypeValue)
}
func deleteCatalogProductOptionTypeValue(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOptionTypeValue CatalogProductOptionTypeValue
	d := g.Where("id = ?", id).Delete(&catalogProductOptionTypeValue)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
