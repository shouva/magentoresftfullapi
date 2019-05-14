package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductBundleStockIndex :
type CatalogProductBundleStockIndex struct {
	EntityId    uint16 `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	WebsiteId   uint16 `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	StockId     uint16 `gorm:"column:stock_id;primary_key" form:"stock_id;primary_key" json:"stock_id;primary_key"`
	OptionId    uint16 `gorm:"column:option_id;primary_key" form:"option_id;primary_key" json:"option_id;primary_key"`
	StockStatus uint16 `gorm:"column:stock_status" form:"stock_status" json:"stock_status"`
}

// TableName :
func (*CatalogProductBundleStockIndex) TableName() string {
	return "catalog_product_bundle_stock_index"
}

// handler create
func initRoutersCatalogProductBundleStockIndex(r *gin.Engine, catalogproductbundlestockindex string) {
	route := r.Group("catalogproductbundlestockindex")
	route.GET("/", getCatalogProductBundleStockIndexs)
	route.GET("/:id", getCatalogProductBundleStockIndex)
	route.POST("/", createCatalogProductBundleStockIndex)
	route.PUT("/:id", updateCatalogProductBundleStockIndex)
	route.DELETE("/:id", deleteCatalogProductBundleStockIndex)
}

func getCatalogProductBundleStockIndexs(c *gin.Context) {
	var catalogProductBundleStockIndexs []CatalogProductBundleStockIndex
	if err := g.Find(&catalogProductBundleStockIndexs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundleStockIndexs)
	}
}

func getCatalogProductBundleStockIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundleStockIndex CatalogProductBundleStockIndex
	if err := g.Where("id = ?", id).First(&catalogProductBundleStockIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundleStockIndex)
	}
}

func createCatalogProductBundleStockIndex(c *gin.Context) {
	var catalogProductBundleStockIndex CatalogProductBundleStockIndex
	c.BindJSON(&catalogProductBundleStockIndex)
	g.Create(&catalogProductBundleStockIndex)
	c.JSON(200, catalogProductBundleStockIndex)
}

func updateCatalogProductBundleStockIndex(c *gin.Context) {
	var catalogProductBundleStockIndex CatalogProductBundleStockIndex
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductBundleStockIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductBundleStockIndex)
	g.Save(&catalogProductBundleStockIndex)
	c.JSON(200, catalogProductBundleStockIndex)
}
func deleteCatalogProductBundleStockIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundleStockIndex CatalogProductBundleStockIndex
	d := g.Where("id = ?", id).Delete(&catalogProductBundleStockIndex)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
