package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCompareItem :
type CatalogCompareItem struct {
	CatalogCompareItemId uint16 `gorm:"column:catalog_compare_item_id;primary_key" form:"catalog_compare_item_id;primary_key" json:"catalog_compare_item_id;primary_key"`
	VisitorId            uint16 `gorm:"column:visitor_id" form:"visitor_id" json:"visitor_id"`
	CustomerId           uint16 `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	ProductId            uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	StoreId              uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
}

// TableName :
func (*CatalogCompareItem) TableName() string {
	return "catalog_compare_item"
}

// handler create
func initRoutersCatalogCompareItem(r *gin.Engine, catalogcompareitem string) {
	route := r.Group("catalogcompareitem")
	route.GET("/", getCatalogCompareItems)
	route.GET("/:id", getCatalogCompareItem)
	route.POST("/", createCatalogCompareItem)
	route.PUT("/:id", updateCatalogCompareItem)
	route.DELETE("/:id", deleteCatalogCompareItem)
}

func getCatalogCompareItems(c *gin.Context) {
	var catalogCompareItems []CatalogCompareItem
	if err := g.Find(&catalogCompareItems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCompareItems)
	}
}

func getCatalogCompareItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCompareItem CatalogCompareItem
	if err := g.Where("id = ?", id).First(&catalogCompareItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCompareItem)
	}
}

func createCatalogCompareItem(c *gin.Context) {
	var catalogCompareItem CatalogCompareItem
	c.BindJSON(&catalogCompareItem)
	g.Create(&catalogCompareItem)
	c.JSON(200, catalogCompareItem)
}

func updateCatalogCompareItem(c *gin.Context) {
	var catalogCompareItem CatalogCompareItem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCompareItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCompareItem)
	g.Save(&catalogCompareItem)
	c.JSON(200, catalogCompareItem)
}
func deleteCatalogCompareItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCompareItem CatalogCompareItem
	d := g.Where("id = ?", id).Delete(&catalogCompareItem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
