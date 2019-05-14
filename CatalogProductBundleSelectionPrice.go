package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductBundleSelectionPrice :
type CatalogProductBundleSelectionPrice struct {
	SelectionId         uint16  `gorm:"column:selection_id;primary_key" form:"selection_id;primary_key" json:"selection_id;primary_key"`
	WebsiteId           uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	SelectionPriceType  uint16  `gorm:"column:selection_price_type" form:"selection_price_type" json:"selection_price_type"`
	SelectionPriceValue float64 `gorm:"column:selection_price_value" form:"selection_price_value" json:"selection_price_value"`
}

// TableName :
func (*CatalogProductBundleSelectionPrice) TableName() string {
	return "catalog_product_bundle_selection_price"
}

// handler create
func initRoutersCatalogProductBundleSelectionPrice(r *gin.Engine, catalogproductbundleselectionprice string) {
	route := r.Group("catalogproductbundleselectionprice")
	route.GET("/", getCatalogProductBundleSelectionPrices)
	route.GET("/:id", getCatalogProductBundleSelectionPrice)
	route.POST("/", createCatalogProductBundleSelectionPrice)
	route.PUT("/:id", updateCatalogProductBundleSelectionPrice)
	route.DELETE("/:id", deleteCatalogProductBundleSelectionPrice)
}

func getCatalogProductBundleSelectionPrices(c *gin.Context) {
	var catalogProductBundleSelectionPrices []CatalogProductBundleSelectionPrice
	if err := g.Find(&catalogProductBundleSelectionPrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundleSelectionPrices)
	}
}

func getCatalogProductBundleSelectionPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundleSelectionPrice CatalogProductBundleSelectionPrice
	if err := g.Where("id = ?", id).First(&catalogProductBundleSelectionPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundleSelectionPrice)
	}
}

func createCatalogProductBundleSelectionPrice(c *gin.Context) {
	var catalogProductBundleSelectionPrice CatalogProductBundleSelectionPrice
	c.BindJSON(&catalogProductBundleSelectionPrice)
	g.Create(&catalogProductBundleSelectionPrice)
	c.JSON(200, catalogProductBundleSelectionPrice)
}

func updateCatalogProductBundleSelectionPrice(c *gin.Context) {
	var catalogProductBundleSelectionPrice CatalogProductBundleSelectionPrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductBundleSelectionPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductBundleSelectionPrice)
	g.Save(&catalogProductBundleSelectionPrice)
	c.JSON(200, catalogProductBundleSelectionPrice)
}
func deleteCatalogProductBundleSelectionPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundleSelectionPrice CatalogProductBundleSelectionPrice
	d := g.Where("id = ?", id).Delete(&catalogProductBundleSelectionPrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
