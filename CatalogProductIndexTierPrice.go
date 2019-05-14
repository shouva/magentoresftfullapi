package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexTierPrice :
type CatalogProductIndexTierPrice struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
}

// TableName :
func (*CatalogProductIndexTierPrice) TableName() string {
	return "catalog_product_index_tier_price"
}

// handler create
func initRoutersCatalogProductIndexTierPrice(r *gin.Engine, catalogproductindextierprice string) {
	route := r.Group("catalogproductindextierprice")
	route.GET("/", getCatalogProductIndexTierPrices)
	route.GET("/:id", getCatalogProductIndexTierPrice)
	route.POST("/", createCatalogProductIndexTierPrice)
	route.PUT("/:id", updateCatalogProductIndexTierPrice)
	route.DELETE("/:id", deleteCatalogProductIndexTierPrice)
}

func getCatalogProductIndexTierPrices(c *gin.Context) {
	var catalogProductIndexTierPrices []CatalogProductIndexTierPrice
	if err := g.Find(&catalogProductIndexTierPrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexTierPrices)
	}
}

func getCatalogProductIndexTierPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexTierPrice CatalogProductIndexTierPrice
	if err := g.Where("id = ?", id).First(&catalogProductIndexTierPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexTierPrice)
	}
}

func createCatalogProductIndexTierPrice(c *gin.Context) {
	var catalogProductIndexTierPrice CatalogProductIndexTierPrice
	c.BindJSON(&catalogProductIndexTierPrice)
	g.Create(&catalogProductIndexTierPrice)
	c.JSON(200, catalogProductIndexTierPrice)
}

func updateCatalogProductIndexTierPrice(c *gin.Context) {
	var catalogProductIndexTierPrice CatalogProductIndexTierPrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexTierPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexTierPrice)
	g.Save(&catalogProductIndexTierPrice)
	c.JSON(200, catalogProductIndexTierPrice)
}
func deleteCatalogProductIndexTierPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexTierPrice CatalogProductIndexTierPrice
	d := g.Where("id = ?", id).Delete(&catalogProductIndexTierPrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
