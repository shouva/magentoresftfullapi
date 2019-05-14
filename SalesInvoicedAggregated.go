package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesInvoicedAggregated :
type SalesInvoicedAggregated struct {
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
func (*SalesInvoicedAggregated) TableName() string {
	return "sales_invoiced_aggregated"
}

// handler create
func initRoutersSalesInvoicedAggregated(r *gin.Engine, salesinvoicedaggregated string) {
	route := r.Group("salesinvoicedaggregated")
	route.GET("/", getSalesInvoicedAggregateds)
	route.GET("/:id", getSalesInvoicedAggregated)
	route.POST("/", createSalesInvoicedAggregated)
	route.PUT("/:id", updateSalesInvoicedAggregated)
	route.DELETE("/:id", deleteSalesInvoicedAggregated)
}

func getSalesInvoicedAggregateds(c *gin.Context) {
	var salesInvoicedAggregateds []SalesInvoicedAggregated
	if err := g.Find(&salesInvoicedAggregateds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesInvoicedAggregateds)
	}
}

func getSalesInvoicedAggregated(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesInvoicedAggregated SalesInvoicedAggregated
	if err := g.Where("id = ?", id).First(&salesInvoicedAggregated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesInvoicedAggregated)
	}
}

func createSalesInvoicedAggregated(c *gin.Context) {
	var salesInvoicedAggregated SalesInvoicedAggregated
	c.BindJSON(&salesInvoicedAggregated)
	g.Create(&salesInvoicedAggregated)
	c.JSON(200, salesInvoicedAggregated)
}

func updateSalesInvoicedAggregated(c *gin.Context) {
	var salesInvoicedAggregated SalesInvoicedAggregated
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesInvoicedAggregated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesInvoicedAggregated)
	g.Save(&salesInvoicedAggregated)
	c.JSON(200, salesInvoicedAggregated)
}
func deleteSalesInvoicedAggregated(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesInvoicedAggregated SalesInvoicedAggregated
	d := g.Where("id = ?", id).Delete(&salesInvoicedAggregated)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
