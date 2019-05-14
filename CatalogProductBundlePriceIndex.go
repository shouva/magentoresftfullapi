package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductBundlePriceIndex :
type CatalogProductBundlePriceIndex struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
}

// TableName :
func (*CatalogProductBundlePriceIndex) TableName() string {
	return "catalog_product_bundle_price_index"
}

// handler create
func initRoutersCatalogProductBundlePriceIndex(r *gin.Engine, catalogproductbundlepriceindex string) {
	route := r.Group("catalogproductbundlepriceindex")
	route.GET("/", getCatalogProductBundlePriceIndexs)
	route.GET("/:id", getCatalogProductBundlePriceIndex)
	route.POST("/", createCatalogProductBundlePriceIndex)
	route.PUT("/:id", updateCatalogProductBundlePriceIndex)
	route.DELETE("/:id", deleteCatalogProductBundlePriceIndex)
}

func getCatalogProductBundlePriceIndexs(c *gin.Context) {
	var catalogProductBundlePriceIndexs []CatalogProductBundlePriceIndex
	if err := g.Find(&catalogProductBundlePriceIndexs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundlePriceIndexs)
	}
}

func getCatalogProductBundlePriceIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundlePriceIndex CatalogProductBundlePriceIndex
	if err := g.Where("id = ?", id).First(&catalogProductBundlePriceIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundlePriceIndex)
	}
}

func createCatalogProductBundlePriceIndex(c *gin.Context) {
	var catalogProductBundlePriceIndex CatalogProductBundlePriceIndex
	c.BindJSON(&catalogProductBundlePriceIndex)
	g.Create(&catalogProductBundlePriceIndex)
	c.JSON(200, catalogProductBundlePriceIndex)
}

func updateCatalogProductBundlePriceIndex(c *gin.Context) {
	var catalogProductBundlePriceIndex CatalogProductBundlePriceIndex
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductBundlePriceIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductBundlePriceIndex)
	g.Save(&catalogProductBundlePriceIndex)
	c.JSON(200, catalogProductBundlePriceIndex)
}
func deleteCatalogProductBundlePriceIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundlePriceIndex CatalogProductBundlePriceIndex
	d := g.Where("id = ?", id).Delete(&catalogProductBundlePriceIndex)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
