package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductEntityGroupPrice :
type CatalogProductEntityGroupPrice struct {
	ValueId         uint16  `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityId        uint16  `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	AllGroups       uint16  `gorm:"column:all_groups" form:"all_groups" json:"all_groups"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id" form:"customer_group_id" json:"customer_group_id"`
	Value           float64 `gorm:"column:value" form:"value" json:"value"`
	WebsiteId       uint16  `gorm:"column:website_id" form:"website_id" json:"website_id"`
}

// TableName :
func (*CatalogProductEntityGroupPrice) TableName() string {
	return "catalog_product_entity_group_price"
}

// handler create
func initRoutersCatalogProductEntityGroupPrice(r *gin.Engine, catalogproductentitygroupprice string) {
	route := r.Group("catalogproductentitygroupprice")
	route.GET("/", getCatalogProductEntityGroupPrices)
	route.GET("/:id", getCatalogProductEntityGroupPrice)
	route.POST("/", createCatalogProductEntityGroupPrice)
	route.PUT("/:id", updateCatalogProductEntityGroupPrice)
	route.DELETE("/:id", deleteCatalogProductEntityGroupPrice)
}

func getCatalogProductEntityGroupPrices(c *gin.Context) {
	var catalogProductEntityGroupPrices []CatalogProductEntityGroupPrice
	if err := g.Find(&catalogProductEntityGroupPrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityGroupPrices)
	}
}

func getCatalogProductEntityGroupPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityGroupPrice CatalogProductEntityGroupPrice
	if err := g.Where("id = ?", id).First(&catalogProductEntityGroupPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEntityGroupPrice)
	}
}

func createCatalogProductEntityGroupPrice(c *gin.Context) {
	var catalogProductEntityGroupPrice CatalogProductEntityGroupPrice
	c.BindJSON(&catalogProductEntityGroupPrice)
	g.Create(&catalogProductEntityGroupPrice)
	c.JSON(200, catalogProductEntityGroupPrice)
}

func updateCatalogProductEntityGroupPrice(c *gin.Context) {
	var catalogProductEntityGroupPrice CatalogProductEntityGroupPrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEntityGroupPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEntityGroupPrice)
	g.Save(&catalogProductEntityGroupPrice)
	c.JSON(200, catalogProductEntityGroupPrice)
}
func deleteCatalogProductEntityGroupPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEntityGroupPrice CatalogProductEntityGroupPrice
	d := g.Where("id = ?", id).Delete(&catalogProductEntityGroupPrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
