package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductBundleOption :
type CatalogProductBundleOption struct {
	OptionId uint16 `gorm:"column:option_id;primary_key" form:"option_id;primary_key" json:"option_id;primary_key"`
	ParentId uint16 `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	Required uint16 `gorm:"column:required" form:"required" json:"required"`
	Position uint16 `gorm:"column:position" form:"position" json:"position"`
	Type     string `gorm:"column:type" form:"type" json:"type"`
}

// TableName :
func (*CatalogProductBundleOption) TableName() string {
	return "catalog_product_bundle_option"
}

// handler create
func initRoutersCatalogProductBundleOption(r *gin.Engine, catalogproductbundleoption string) {
	route := r.Group("catalogproductbundleoption")
	route.GET("/", getCatalogProductBundleOptions)
	route.GET("/:id", getCatalogProductBundleOption)
	route.POST("/", createCatalogProductBundleOption)
	route.PUT("/:id", updateCatalogProductBundleOption)
	route.DELETE("/:id", deleteCatalogProductBundleOption)
}

func getCatalogProductBundleOptions(c *gin.Context) {
	var catalogProductBundleOptions []CatalogProductBundleOption
	if err := g.Find(&catalogProductBundleOptions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundleOptions)
	}
}

func getCatalogProductBundleOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundleOption CatalogProductBundleOption
	if err := g.Where("id = ?", id).First(&catalogProductBundleOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundleOption)
	}
}

func createCatalogProductBundleOption(c *gin.Context) {
	var catalogProductBundleOption CatalogProductBundleOption
	c.BindJSON(&catalogProductBundleOption)
	g.Create(&catalogProductBundleOption)
	c.JSON(200, catalogProductBundleOption)
}

func updateCatalogProductBundleOption(c *gin.Context) {
	var catalogProductBundleOption CatalogProductBundleOption
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductBundleOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductBundleOption)
	g.Save(&catalogProductBundleOption)
	c.JSON(200, catalogProductBundleOption)
}
func deleteCatalogProductBundleOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundleOption CatalogProductBundleOption
	d := g.Where("id = ?", id).Delete(&catalogProductBundleOption)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
