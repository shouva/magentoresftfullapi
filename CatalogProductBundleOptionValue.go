package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductBundleOptionValue :
type CatalogProductBundleOptionValue struct {
	ValueId  uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	OptionId uint16 `gorm:"column:option_id" form:"option_id" json:"option_id"`
	StoreId  uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Title    string `gorm:"column:title" form:"title" json:"title"`
}

// TableName :
func (*CatalogProductBundleOptionValue) TableName() string {
	return "catalog_product_bundle_option_value"
}

// handler create
func initRoutersCatalogProductBundleOptionValue(r *gin.Engine, catalogproductbundleoptionvalue string) {
	route := r.Group("catalogproductbundleoptionvalue")
	route.GET("/", getCatalogProductBundleOptionValues)
	route.GET("/:id", getCatalogProductBundleOptionValue)
	route.POST("/", createCatalogProductBundleOptionValue)
	route.PUT("/:id", updateCatalogProductBundleOptionValue)
	route.DELETE("/:id", deleteCatalogProductBundleOptionValue)
}

func getCatalogProductBundleOptionValues(c *gin.Context) {
	var catalogProductBundleOptionValues []CatalogProductBundleOptionValue
	if err := g.Find(&catalogProductBundleOptionValues).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundleOptionValues)
	}
}

func getCatalogProductBundleOptionValue(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundleOptionValue CatalogProductBundleOptionValue
	if err := g.Where("id = ?", id).First(&catalogProductBundleOptionValue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundleOptionValue)
	}
}

func createCatalogProductBundleOptionValue(c *gin.Context) {
	var catalogProductBundleOptionValue CatalogProductBundleOptionValue
	c.BindJSON(&catalogProductBundleOptionValue)
	g.Create(&catalogProductBundleOptionValue)
	c.JSON(200, catalogProductBundleOptionValue)
}

func updateCatalogProductBundleOptionValue(c *gin.Context) {
	var catalogProductBundleOptionValue CatalogProductBundleOptionValue
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductBundleOptionValue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductBundleOptionValue)
	g.Save(&catalogProductBundleOptionValue)
	c.JSON(200, catalogProductBundleOptionValue)
}
func deleteCatalogProductBundleOptionValue(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundleOptionValue CatalogProductBundleOptionValue
	d := g.Where("id = ?", id).Delete(&catalogProductBundleOptionValue)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
