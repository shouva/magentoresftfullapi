package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductOptionTypePrice :
type CatalogProductOptionTypePrice struct {
	OptionTypePriceId uint16  `gorm:"column:option_type_price_id;primary_key" form:"option_type_price_id;primary_key" json:"option_type_price_id;primary_key"`
	OptionTypeId      uint16  `gorm:"column:option_type_id" form:"option_type_id" json:"option_type_id"`
	StoreId           uint16  `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Price             float64 `gorm:"column:price" form:"price" json:"price"`
	PriceType         string  `gorm:"column:price_type" form:"price_type" json:"price_type"`
}

// TableName :
func (*CatalogProductOptionTypePrice) TableName() string {
	return "catalog_product_option_type_price"
}

// handler create
func initRoutersCatalogProductOptionTypePrice(r *gin.Engine, catalogproductoptiontypeprice string) {
	route := r.Group("catalogproductoptiontypeprice")
	route.GET("/", getCatalogProductOptionTypePrices)
	route.GET("/:id", getCatalogProductOptionTypePrice)
	route.POST("/", createCatalogProductOptionTypePrice)
	route.PUT("/:id", updateCatalogProductOptionTypePrice)
	route.DELETE("/:id", deleteCatalogProductOptionTypePrice)
}

func getCatalogProductOptionTypePrices(c *gin.Context) {
	var catalogProductOptionTypePrices []CatalogProductOptionTypePrice
	if err := g.Find(&catalogProductOptionTypePrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptionTypePrices)
	}
}

func getCatalogProductOptionTypePrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOptionTypePrice CatalogProductOptionTypePrice
	if err := g.Where("id = ?", id).First(&catalogProductOptionTypePrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptionTypePrice)
	}
}

func createCatalogProductOptionTypePrice(c *gin.Context) {
	var catalogProductOptionTypePrice CatalogProductOptionTypePrice
	c.BindJSON(&catalogProductOptionTypePrice)
	g.Create(&catalogProductOptionTypePrice)
	c.JSON(200, catalogProductOptionTypePrice)
}

func updateCatalogProductOptionTypePrice(c *gin.Context) {
	var catalogProductOptionTypePrice CatalogProductOptionTypePrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductOptionTypePrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductOptionTypePrice)
	g.Save(&catalogProductOptionTypePrice)
	c.JSON(200, catalogProductOptionTypePrice)
}
func deleteCatalogProductOptionTypePrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOptionTypePrice CatalogProductOptionTypePrice
	d := g.Where("id = ?", id).Delete(&catalogProductOptionTypePrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
