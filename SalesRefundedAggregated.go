package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesRefundedAggregated :
type SalesRefundedAggregated struct {
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
func (*SalesRefundedAggregated) TableName() string {
	return "sales_refunded_aggregated"
}

// handler create
func initRoutersSalesRefundedAggregated(r *gin.Engine, salesrefundedaggregated string) {
	route := r.Group("salesrefundedaggregated")
	route.GET("/", getSalesRefundedAggregateds)
	route.GET("/:id", getSalesRefundedAggregated)
	route.POST("/", createSalesRefundedAggregated)
	route.PUT("/:id", updateSalesRefundedAggregated)
	route.DELETE("/:id", deleteSalesRefundedAggregated)
}

func getSalesRefundedAggregateds(c *gin.Context) {
	var salesRefundedAggregateds []SalesRefundedAggregated
	if err := g.Find(&salesRefundedAggregateds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesRefundedAggregateds)
	}
}

func getSalesRefundedAggregated(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesRefundedAggregated SalesRefundedAggregated
	if err := g.Where("id = ?", id).First(&salesRefundedAggregated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesRefundedAggregated)
	}
}

func createSalesRefundedAggregated(c *gin.Context) {
	var salesRefundedAggregated SalesRefundedAggregated
	c.BindJSON(&salesRefundedAggregated)
	g.Create(&salesRefundedAggregated)
	c.JSON(200, salesRefundedAggregated)
}

func updateSalesRefundedAggregated(c *gin.Context) {
	var salesRefundedAggregated SalesRefundedAggregated
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesRefundedAggregated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesRefundedAggregated)
	g.Save(&salesRefundedAggregated)
	c.JSON(200, salesRefundedAggregated)
}
func deleteSalesRefundedAggregated(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesRefundedAggregated SalesRefundedAggregated
	d := g.Where("id = ?", id).Delete(&salesRefundedAggregated)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
