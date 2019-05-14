package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesShippingAggregated :
type SalesShippingAggregated struct {
	ID                  uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Period              *time.Time `gorm:"column:period" form:"period" json:"period"`
	StoreId             uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	OrderStatus         string     `gorm:"column:order_status" form:"order_status" json:"order_status"`
	ShippingDescription string     `gorm:"column:shipping_description" form:"shipping_description" json:"shipping_description"`
	OrdersCount         uint16     `gorm:"column:orders_count" form:"orders_count" json:"orders_count"`
	TotalShipping       float64    `gorm:"column:total_shipping" form:"total_shipping" json:"total_shipping"`
	TotalShippingActual float64    `gorm:"column:total_shipping_actual" form:"total_shipping_actual" json:"total_shipping_actual"`
}

// TableName :
func (*SalesShippingAggregated) TableName() string {
	return "sales_shipping_aggregated"
}

// handler create
func initRoutersSalesShippingAggregated(r *gin.Engine, salesshippingaggregated string) {
	route := r.Group("salesshippingaggregated")
	route.GET("/", getSalesShippingAggregateds)
	route.GET("/:id", getSalesShippingAggregated)
	route.POST("/", createSalesShippingAggregated)
	route.PUT("/:id", updateSalesShippingAggregated)
	route.DELETE("/:id", deleteSalesShippingAggregated)
}

func getSalesShippingAggregateds(c *gin.Context) {
	var salesShippingAggregateds []SalesShippingAggregated
	if err := g.Find(&salesShippingAggregateds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesShippingAggregateds)
	}
}

func getSalesShippingAggregated(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesShippingAggregated SalesShippingAggregated
	if err := g.Where("id = ?", id).First(&salesShippingAggregated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesShippingAggregated)
	}
}

func createSalesShippingAggregated(c *gin.Context) {
	var salesShippingAggregated SalesShippingAggregated
	c.BindJSON(&salesShippingAggregated)
	g.Create(&salesShippingAggregated)
	c.JSON(200, salesShippingAggregated)
}

func updateSalesShippingAggregated(c *gin.Context) {
	var salesShippingAggregated SalesShippingAggregated
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesShippingAggregated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesShippingAggregated)
	g.Save(&salesShippingAggregated)
	c.JSON(200, salesShippingAggregated)
}
func deleteSalesShippingAggregated(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesShippingAggregated SalesShippingAggregated
	d := g.Where("id = ?", id).Delete(&salesShippingAggregated)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
