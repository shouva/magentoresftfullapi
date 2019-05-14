package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesOrderAggregatedUpdated :
type SalesOrderAggregatedUpdated struct {
	ID                        uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Period                    *time.Time `gorm:"column:period" form:"period" json:"period"`
	StoreId                   uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	OrderStatus               string     `gorm:"column:order_status" form:"order_status" json:"order_status"`
	OrdersCount               uint16     `gorm:"column:orders_count" form:"orders_count" json:"orders_count"`
	TotalQtyOrdered           float64    `gorm:"column:total_qty_ordered" form:"total_qty_ordered" json:"total_qty_ordered"`
	TotalQtyInvoiced          float64    `gorm:"column:total_qty_invoiced" form:"total_qty_invoiced" json:"total_qty_invoiced"`
	TotalIncomeAmount         float64    `gorm:"column:total_income_amount" form:"total_income_amount" json:"total_income_amount"`
	TotalRevenueAmount        float64    `gorm:"column:total_revenue_amount" form:"total_revenue_amount" json:"total_revenue_amount"`
	TotalProfitAmount         float64    `gorm:"column:total_profit_amount" form:"total_profit_amount" json:"total_profit_amount"`
	TotalInvoicedAmount       float64    `gorm:"column:total_invoiced_amount" form:"total_invoiced_amount" json:"total_invoiced_amount"`
	TotalCanceledAmount       float64    `gorm:"column:total_canceled_amount" form:"total_canceled_amount" json:"total_canceled_amount"`
	TotalPaidAmount           float64    `gorm:"column:total_paid_amount" form:"total_paid_amount" json:"total_paid_amount"`
	TotalRefundedAmount       float64    `gorm:"column:total_refunded_amount" form:"total_refunded_amount" json:"total_refunded_amount"`
	TotalTaxAmount            float64    `gorm:"column:total_tax_amount" form:"total_tax_amount" json:"total_tax_amount"`
	TotalTaxAmountActual      float64    `gorm:"column:total_tax_amount_actual" form:"total_tax_amount_actual" json:"total_tax_amount_actual"`
	TotalShippingAmount       float64    `gorm:"column:total_shipping_amount" form:"total_shipping_amount" json:"total_shipping_amount"`
	TotalShippingAmountActual float64    `gorm:"column:total_shipping_amount_actual" form:"total_shipping_amount_actual" json:"total_shipping_amount_actual"`
	TotalDiscountAmount       float64    `gorm:"column:total_discount_amount" form:"total_discount_amount" json:"total_discount_amount"`
	TotalDiscountAmountActual float64    `gorm:"column:total_discount_amount_actual" form:"total_discount_amount_actual" json:"total_discount_amount_actual"`
}

// TableName :
func (*SalesOrderAggregatedUpdated) TableName() string {
	return "sales_order_aggregated_updated"
}

// handler create
func initRoutersSalesOrderAggregatedUpdated(r *gin.Engine, salesorderaggregatedupdated string) {
	route := r.Group("salesorderaggregatedupdated")
	route.GET("/", getSalesOrderAggregatedUpdateds)
	route.GET("/:id", getSalesOrderAggregatedUpdated)
	route.POST("/", createSalesOrderAggregatedUpdated)
	route.PUT("/:id", updateSalesOrderAggregatedUpdated)
	route.DELETE("/:id", deleteSalesOrderAggregatedUpdated)
}

func getSalesOrderAggregatedUpdateds(c *gin.Context) {
	var salesOrderAggregatedUpdateds []SalesOrderAggregatedUpdated
	if err := g.Find(&salesOrderAggregatedUpdateds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderAggregatedUpdateds)
	}
}

func getSalesOrderAggregatedUpdated(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderAggregatedUpdated SalesOrderAggregatedUpdated
	if err := g.Where("id = ?", id).First(&salesOrderAggregatedUpdated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderAggregatedUpdated)
	}
}

func createSalesOrderAggregatedUpdated(c *gin.Context) {
	var salesOrderAggregatedUpdated SalesOrderAggregatedUpdated
	c.BindJSON(&salesOrderAggregatedUpdated)
	g.Create(&salesOrderAggregatedUpdated)
	c.JSON(200, salesOrderAggregatedUpdated)
}

func updateSalesOrderAggregatedUpdated(c *gin.Context) {
	var salesOrderAggregatedUpdated SalesOrderAggregatedUpdated
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesOrderAggregatedUpdated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesOrderAggregatedUpdated)
	g.Save(&salesOrderAggregatedUpdated)
	c.JSON(200, salesOrderAggregatedUpdated)
}
func deleteSalesOrderAggregatedUpdated(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderAggregatedUpdated SalesOrderAggregatedUpdated
	d := g.Where("id = ?", id).Delete(&salesOrderAggregatedUpdated)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
