package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesInvoicedAggregatedOrder :
type SalesInvoicedAggregatedOrder struct {
	ID                  uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Period              *time.Time `gorm:"column:period" form:"period" json:"period"`
	StoreId             uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	OrderStatus         string     `gorm:"column:order_status" form:"order_status" json:"order_status"`
	OrdersCount         uint16     `gorm:"column:orders_count" form:"orders_count" json:"orders_count"`
	OrdersInvoiced      float64    `gorm:"column:orders_invoiced" form:"orders_invoiced" json:"orders_invoiced"`
	Invoiced            float64    `gorm:"column:invoiced" form:"invoiced" json:"invoiced"`
	InvoicedCaptured    float64    `gorm:"column:invoiced_captured" form:"invoiced_captured" json:"invoiced_captured"`
	InvoicedNotCaptured float64    `gorm:"column:invoiced_not_captured" form:"invoiced_not_captured" json:"invoiced_not_captured"`
}

// TableName :
func (*SalesInvoicedAggregatedOrder) TableName() string {
	return "sales_invoiced_aggregated_order"
}

// handler create
func initRoutersSalesInvoicedAggregatedOrder(r *gin.Engine, salesinvoicedaggregatedorder string) {
	route := r.Group("salesinvoicedaggregatedorder")
	route.GET("/", getSalesInvoicedAggregatedOrders)
	route.GET("/:id", getSalesInvoicedAggregatedOrder)
	route.POST("/", createSalesInvoicedAggregatedOrder)
	route.PUT("/:id", updateSalesInvoicedAggregatedOrder)
	route.DELETE("/:id", deleteSalesInvoicedAggregatedOrder)
}

func getSalesInvoicedAggregatedOrders(c *gin.Context) {
	var salesInvoicedAggregatedOrders []SalesInvoicedAggregatedOrder
	if err := g.Find(&salesInvoicedAggregatedOrders).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesInvoicedAggregatedOrders)
	}
}

func getSalesInvoicedAggregatedOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesInvoicedAggregatedOrder SalesInvoicedAggregatedOrder
	if err := g.Where("id = ?", id).First(&salesInvoicedAggregatedOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesInvoicedAggregatedOrder)
	}
}

func createSalesInvoicedAggregatedOrder(c *gin.Context) {
	var salesInvoicedAggregatedOrder SalesInvoicedAggregatedOrder
	c.BindJSON(&salesInvoicedAggregatedOrder)
	g.Create(&salesInvoicedAggregatedOrder)
	c.JSON(200, salesInvoicedAggregatedOrder)
}

func updateSalesInvoicedAggregatedOrder(c *gin.Context) {
	var salesInvoicedAggregatedOrder SalesInvoicedAggregatedOrder
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesInvoicedAggregatedOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesInvoicedAggregatedOrder)
	g.Save(&salesInvoicedAggregatedOrder)
	c.JSON(200, salesInvoicedAggregatedOrder)
}
func deleteSalesInvoicedAggregatedOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesInvoicedAggregatedOrder SalesInvoicedAggregatedOrder
	d := g.Where("id = ?", id).Delete(&salesInvoicedAggregatedOrder)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
