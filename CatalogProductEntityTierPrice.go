package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntityTierPrice :
type CatalogProductEntityTierPrice struct {
	ValueId         uint16  `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityId        uint16  `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	AllGroups       uint16  `gorm:"column:all_groups" form:"all_groups" json:"all_groups"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id" form:"customer_group_id" json:"customer_group_id"`
	Qty             float64 `gorm:"column:qty" form:"qty" json:"qty"`
	Value           float64 `gorm:"column:value" form:"value" json:"value"`
	WebsiteId       uint16  `gorm:"column:website_id" form:"website_id" json:"website_id"`
}

// TableName :
func (*CatalogProductEntityTierPrice) TableName() string {
	return "catalog_product_entity_tier_price"
}

// handler create
func initRoutersCatalogProductEntityTierPrice(r *gin.Engine, catalogproductentitytierprice string) {
	route := r.Group("catalogproductentitytierprice")
	route.GET("/", getCatalogProductEntityTierPrices)
	route.GET("/:id", getCatalogProductEntityTierPrice)
	route.POST("/", createCatalogProductEntityTierPrice)
	route.PUT("/:id", updateCatalogProductEntityTierPrice)
	route.DELETE("/:id", deleteCatalogProductEntityTierPrice)
}

func getCatalogProductEntityTierPrices(c *gin.Context) {
	var catalogProductEntityTierPrices []CatalogProductEntityTierPrice
	if err := g.Find(&catalogProductEntityTierPrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityTierPrices)
	}
}

func getCatalogProductEntityTierPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityTierPrice CatalogProductEntityTierPrice
	if err := g.Where("id = ?", id).First(&catalogProductEntityTierPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityTierPrice)
	}
}

func createCatalogProductEntityTierPrice(c *gin.Context) {
	var catalogProductEntityTierPrice CatalogProductEntityTierPrice
	c.BindJSON(&catalogProductEntityTierPrice)
	g.Create(&catalogProductEntityTierPrice)
	c.JSON(200, catalogProductEntityTierPrice)
}

func updateCatalogProductEntityTierPrice(c *gin.Context) {
	var catalogProductEntityTierPrice CatalogProductEntityTierPrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntityTierPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntityTierPrice)
	g.Save(&catalogProductEntityTierPrice)
	c.JSON(200, catalogProductEntityTierPrice)
}
func deleteCatalogProductEntityTierPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityTierPrice CatalogProductEntityTierPrice
	d := g.Where("id = ?", id).Delete(&catalogProductEntityTierPrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
