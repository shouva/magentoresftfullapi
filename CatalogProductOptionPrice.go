package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductOptionPrice :
type CatalogProductOptionPrice struct {
	OptionPriceId uint16  `gorm:"column:option_price_id;primary_key" form:"option_price_id;primary_key" json:"option_price_id;primary_key"`
	OptionId      uint16  `gorm:"column:option_id" form:"option_id" json:"option_id"`
	StoreId       uint16  `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Price         float64 `gorm:"column:price" form:"price" json:"price"`
	PriceType     string  `gorm:"column:price_type" form:"price_type" json:"price_type"`
}

// TableName :
func (*CatalogProductOptionPrice) TableName() string {
	return "catalog_product_option_price"
}

// handler create
func initRoutersCatalogProductOptionPrice(r *gin.Engine, catalogproductoptionprice string) {
	route := r.Group("catalogproductoptionprice")
	route.GET("/", getCatalogProductOptionPrices)
	route.GET("/:id", getCatalogProductOptionPrice)
	route.POST("/", createCatalogProductOptionPrice)
	route.PUT("/:id", updateCatalogProductOptionPrice)
	route.DELETE("/:id", deleteCatalogProductOptionPrice)
}

func getCatalogProductOptionPrices(c *gin.Context) {
	var catalogProductOptionPrices []CatalogProductOptionPrice
	if err := g.Find(&catalogProductOptionPrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptionPrices)
	}
}

func getCatalogProductOptionPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOptionPrice CatalogProductOptionPrice
	if err := g.Where("id = ?", id).First(&catalogProductOptionPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptionPrice)
	}
}

func createCatalogProductOptionPrice(c *gin.Context) {
	var catalogProductOptionPrice CatalogProductOptionPrice
	c.BindJSON(&catalogProductOptionPrice)
	g.Create(&catalogProductOptionPrice)
	c.JSON(200, catalogProductOptionPrice)
}

func updateCatalogProductOptionPrice(c *gin.Context) {
	var catalogProductOptionPrice CatalogProductOptionPrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductOptionPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductOptionPrice)
	g.Save(&catalogProductOptionPrice)
	c.JSON(200, catalogProductOptionPrice)
}
func deleteCatalogProductOptionPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOptionPrice CatalogProductOptionPrice
	d := g.Where("id = ?", id).Delete(&catalogProductOptionPrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
