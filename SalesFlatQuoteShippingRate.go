package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatQuoteShippingRate :
type SalesFlatQuoteShippingRate struct {
	RateId            uint16     `gorm:"column:rate_id;primary_key" form:"rate_id;primary_key" json:"rate_id;primary_key"`
	AddressId         uint16     `gorm:"column:address_id" form:"address_id" json:"address_id"`
	CreatedAt         *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	Carrier           string     `gorm:"column:carrier" form:"carrier" json:"carrier"`
	CarrierTitle      string     `gorm:"column:carrier_title" form:"carrier_title" json:"carrier_title"`
	Code              string     `gorm:"column:code" form:"code" json:"code"`
	Method            string     `gorm:"column:method" form:"method" json:"method"`
	MethodDescription string     `gorm:"column:method_description" form:"method_description" json:"method_description"`
	Price             float64    `gorm:"column:price" form:"price" json:"price"`
	ErrorMessage      string     `gorm:"column:error_message" form:"error_message" json:"error_message"`
	MethodTitle       string     `gorm:"column:method_title" form:"method_title" json:"method_title"`
}

// TableName :
func (*SalesFlatQuoteShippingRate) TableName() string {
	return "sales_flat_quote_shipping_rate"
}

// handler create
func initRoutersSalesFlatQuoteShippingRate(r *gin.Engine, salesflatquoteshippingrate string) {
	route := r.Group("salesflatquoteshippingrate")
	route.GET("/", getSalesFlatQuoteShippingRates)
	route.GET("/:id", getSalesFlatQuoteShippingRate)
	route.POST("/", createSalesFlatQuoteShippingRate)
	route.PUT("/:id", updateSalesFlatQuoteShippingRate)
	route.DELETE("/:id", deleteSalesFlatQuoteShippingRate)
}

func getSalesFlatQuoteShippingRates(c *gin.Context) {
	var salesFlatQuoteShippingRates []SalesFlatQuoteShippingRate
	if err := g.Find(&salesFlatQuoteShippingRates).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuoteShippingRates)
	}
}

func getSalesFlatQuoteShippingRate(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuoteShippingRate SalesFlatQuoteShippingRate
	if err := g.Where("id = ?", id).First(&salesFlatQuoteShippingRate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuoteShippingRate)
	}
}

func createSalesFlatQuoteShippingRate(c *gin.Context) {
	var salesFlatQuoteShippingRate SalesFlatQuoteShippingRate
	c.BindJSON(&salesFlatQuoteShippingRate)
	g.Create(&salesFlatQuoteShippingRate)
	c.JSON(200, salesFlatQuoteShippingRate)
}

func updateSalesFlatQuoteShippingRate(c *gin.Context) {
	var salesFlatQuoteShippingRate SalesFlatQuoteShippingRate
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatQuoteShippingRate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatQuoteShippingRate)
	g.Save(&salesFlatQuoteShippingRate)
	c.JSON(200, salesFlatQuoteShippingRate)
}
func deleteSalesFlatQuoteShippingRate(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuoteShippingRate SalesFlatQuoteShippingRate
	d := g.Where("id = ?", id).Delete(&salesFlatQuoteShippingRate)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
