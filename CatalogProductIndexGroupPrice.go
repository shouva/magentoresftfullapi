package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexGroupPrice :
type CatalogProductIndexGroupPrice struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	Price           float64 `gorm:"column:price" form:"price" json:"price"`
}

// TableName :
func (*CatalogProductIndexGroupPrice) TableName() string {
	return "catalog_product_index_group_price"
}

// handler create
func initRoutersCatalogProductIndexGroupPrice(r *gin.Engine, catalogproductindexgroupprice string) {
	route := r.Group("catalogproductindexgroupprice")
	route.GET("/", getCatalogProductIndexGroupPrices)
	route.GET("/:id", getCatalogProductIndexGroupPrice)
	route.POST("/", createCatalogProductIndexGroupPrice)
	route.PUT("/:id", updateCatalogProductIndexGroupPrice)
	route.DELETE("/:id", deleteCatalogProductIndexGroupPrice)
}

func getCatalogProductIndexGroupPrices(c *gin.Context) {
	var catalogProductIndexGroupPrices []CatalogProductIndexGroupPrice
	if err := g.Find(&catalogProductIndexGroupPrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexGroupPrices)
	}
}

func getCatalogProductIndexGroupPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexGroupPrice CatalogProductIndexGroupPrice
	if err := g.Where("id = ?", id).First(&catalogProductIndexGroupPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexGroupPrice)
	}
}

func createCatalogProductIndexGroupPrice(c *gin.Context) {
	var catalogProductIndexGroupPrice CatalogProductIndexGroupPrice
	c.BindJSON(&catalogProductIndexGroupPrice)
	g.Create(&catalogProductIndexGroupPrice)
	c.JSON(200, catalogProductIndexGroupPrice)
}

func updateCatalogProductIndexGroupPrice(c *gin.Context) {
	var catalogProductIndexGroupPrice CatalogProductIndexGroupPrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexGroupPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexGroupPrice)
	g.Save(&catalogProductIndexGroupPrice)
	c.JSON(200, catalogProductIndexGroupPrice)
}
func deleteCatalogProductIndexGroupPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexGroupPrice CatalogProductIndexGroupPrice
	d := g.Where("id = ?", id).Delete(&catalogProductIndexGroupPrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
