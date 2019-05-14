package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesFlatQuoteItemOption :
type SalesFlatQuoteItemOption struct {
	OptionId  uint16 `gorm:"column:option_id;primary_key" form:"option_id;primary_key" json:"option_id;primary_key"`
	ItemId    uint16 `gorm:"column:item_id" form:"item_id" json:"item_id"`
	ProductId uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Code      string `gorm:"column:code" form:"code" json:"code"`
	Value     string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*SalesFlatQuoteItemOption) TableName() string {
	return "sales_flat_quote_item_option"
}

// handler create
func initRoutersSalesFlatQuoteItemOption(r *gin.Engine, salesflatquoteitemoption string) {
	route := r.Group("salesflatquoteitemoption")
	route.GET("/", getSalesFlatQuoteItemOptions)
	route.GET("/:id", getSalesFlatQuoteItemOption)
	route.POST("/", createSalesFlatQuoteItemOption)
	route.PUT("/:id", updateSalesFlatQuoteItemOption)
	route.DELETE("/:id", deleteSalesFlatQuoteItemOption)
}

func getSalesFlatQuoteItemOptions(c *gin.Context) {
	var salesFlatQuoteItemOptions []SalesFlatQuoteItemOption
	if err := g.Find(&salesFlatQuoteItemOptions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuoteItemOptions)
	}
}

func getSalesFlatQuoteItemOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuoteItemOption SalesFlatQuoteItemOption
	if err := g.Where("id = ?", id).First(&salesFlatQuoteItemOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuoteItemOption)
	}
}

func createSalesFlatQuoteItemOption(c *gin.Context) {
	var salesFlatQuoteItemOption SalesFlatQuoteItemOption
	c.BindJSON(&salesFlatQuoteItemOption)
	g.Create(&salesFlatQuoteItemOption)
	c.JSON(200, salesFlatQuoteItemOption)
}

func updateSalesFlatQuoteItemOption(c *gin.Context) {
	var salesFlatQuoteItemOption SalesFlatQuoteItemOption
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatQuoteItemOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatQuoteItemOption)
	g.Save(&salesFlatQuoteItemOption)
	c.JSON(200, salesFlatQuoteItemOption)
}
func deleteSalesFlatQuoteItemOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuoteItemOption SalesFlatQuoteItemOption
	d := g.Where("id = ?", id).Delete(&salesFlatQuoteItemOption)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
