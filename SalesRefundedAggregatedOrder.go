package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesRefundedAggregatedOrder :
type SalesRefundedAggregatedOrder struct {
	ID              uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Period          *time.Time `gorm:"column:period" form:"period" json:"period"`
	StoreId         uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	OrderStatus     string     `gorm:"column:order_status" form:"order_status" json:"order_status"`
	OrdersCount     uint16     `gorm:"column:orders_count" form:"orders_count" json:"orders_count"`
	Refunded        float64    `gorm:"column:refunded" form:"refunded" json:"refunded"`
	OnlineRefunded  float64    `gorm:"column:online_refunded" form:"online_refunded" json:"online_refunded"`
	OfflineRefunded float64    `gorm:"column:offline_refunded" form:"offline_refunded" json:"offline_refunded"`
}

// TableName :
func (*SalesRefundedAggregatedOrder) TableName() string {
	return "sales_refunded_aggregated_order"
}

// handler create
func initRoutersSalesRefundedAggregatedOrder(r *gin.Engine, salesrefundedaggregatedorder string) {
	route := r.Group("salesrefundedaggregatedorder")
	route.GET("/", getSalesRefundedAggregatedOrders)
	route.GET("/:id", getSalesRefundedAggregatedOrder)
	route.POST("/", createSalesRefundedAggregatedOrder)
	route.PUT("/:id", updateSalesRefundedAggregatedOrder)
	route.DELETE("/:id", deleteSalesRefundedAggregatedOrder)
}

func getSalesRefundedAggregatedOrders(c *gin.Context) {
	var salesRefundedAggregatedOrders []SalesRefundedAggregatedOrder
	if err := g.Find(&salesRefundedAggregatedOrders).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesRefundedAggregatedOrders)
	}
}

func getSalesRefundedAggregatedOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesRefundedAggregatedOrder SalesRefundedAggregatedOrder
	if err := g.Where("id = ?", id).First(&salesRefundedAggregatedOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesRefundedAggregatedOrder)
	}
}

func createSalesRefundedAggregatedOrder(c *gin.Context) {
	var salesRefundedAggregatedOrder SalesRefundedAggregatedOrder
	c.BindJSON(&salesRefundedAggregatedOrder)
	g.Create(&salesRefundedAggregatedOrder)
	c.JSON(200, salesRefundedAggregatedOrder)
}

func updateSalesRefundedAggregatedOrder(c *gin.Context) {
	var salesRefundedAggregatedOrder SalesRefundedAggregatedOrder
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesRefundedAggregatedOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesRefundedAggregatedOrder)
	g.Save(&salesRefundedAggregatedOrder)
	c.JSON(200, salesRefundedAggregatedOrder)
}
func deleteSalesRefundedAggregatedOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesRefundedAggregatedOrder SalesRefundedAggregatedOrder
	d := g.Where("id = ?", id).Delete(&salesRefundedAggregatedOrder)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
