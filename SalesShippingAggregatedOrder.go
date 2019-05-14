package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesShippingAggregatedOrder :
type SalesShippingAggregatedOrder struct {
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
func (*SalesShippingAggregatedOrder) TableName() string {
	return "sales_shipping_aggregated_order"
}

// handler create
func initRoutersSalesShippingAggregatedOrder(r *gin.Engine, salesshippingaggregatedorder string) {
	route := r.Group("salesshippingaggregatedorder")
	route.GET("/", getSalesShippingAggregatedOrders)
	route.GET("/:id", getSalesShippingAggregatedOrder)
	route.POST("/", createSalesShippingAggregatedOrder)
	route.PUT("/:id", updateSalesShippingAggregatedOrder)
	route.DELETE("/:id", deleteSalesShippingAggregatedOrder)
}

func getSalesShippingAggregatedOrders(c *gin.Context) {
	var salesShippingAggregatedOrders []SalesShippingAggregatedOrder
	if err := g.Find(&salesShippingAggregatedOrders).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesShippingAggregatedOrders)
	}
}

func getSalesShippingAggregatedOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesShippingAggregatedOrder SalesShippingAggregatedOrder
	if err := g.Where("id = ?", id).First(&salesShippingAggregatedOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesShippingAggregatedOrder)
	}
}

func createSalesShippingAggregatedOrder(c *gin.Context) {
	var salesShippingAggregatedOrder SalesShippingAggregatedOrder
	c.BindJSON(&salesShippingAggregatedOrder)
	g.Create(&salesShippingAggregatedOrder)
	c.JSON(200, salesShippingAggregatedOrder)
}

func updateSalesShippingAggregatedOrder(c *gin.Context) {
	var salesShippingAggregatedOrder SalesShippingAggregatedOrder
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesShippingAggregatedOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesShippingAggregatedOrder)
	g.Save(&salesShippingAggregatedOrder)
	c.JSON(200, salesShippingAggregatedOrder)
}
func deleteSalesShippingAggregatedOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesShippingAggregatedOrder SalesShippingAggregatedOrder
	d := g.Where("id = ?", id).Delete(&salesShippingAggregatedOrder)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
