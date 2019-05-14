package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPrice :
type CatalogProductIndexPrice struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	TaxClassId      uint16  `gorm:"column:tax_class_id" form:"tax_class_id" json:"tax_class_id"`
	Price           float64 `gorm:"column:price" form:"price" json:"price"`
	FinalPrice      float64 `gorm:"column:final_price" form:"final_price" json:"final_price"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
}

// TableName :
func (*CatalogProductIndexPrice) TableName() string {
	return "catalog_product_index_price"
}

// handler create
func initRoutersCatalogProductIndexPrice(r *gin.Engine, catalogproductindexprice string) {
	route := r.Group("catalogproductindexprice")
	route.GET("/", getCatalogProductIndexPrices)
	route.GET("/:id", getCatalogProductIndexPrice)
	route.POST("/", createCatalogProductIndexPrice)
	route.PUT("/:id", updateCatalogProductIndexPrice)
	route.DELETE("/:id", deleteCatalogProductIndexPrice)
}

func getCatalogProductIndexPrices(c *gin.Context) {
	var catalogProductIndexPrices []CatalogProductIndexPrice
	if err := g.Find(&catalogProductIndexPrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPrices)
	}
}

func getCatalogProductIndexPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPrice CatalogProductIndexPrice
	if err := g.Where("id = ?", id).First(&catalogProductIndexPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPrice)
	}
}

func createCatalogProductIndexPrice(c *gin.Context) {
	var catalogProductIndexPrice CatalogProductIndexPrice
	c.BindJSON(&catalogProductIndexPrice)
	g.Create(&catalogProductIndexPrice)
	c.JSON(200, catalogProductIndexPrice)
}

func updateCatalogProductIndexPrice(c *gin.Context) {
	var catalogProductIndexPrice CatalogProductIndexPrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPrice)
	g.Save(&catalogProductIndexPrice)
	c.JSON(200, catalogProductIndexPrice)
}
func deleteCatalogProductIndexPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPrice CatalogProductIndexPrice
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
