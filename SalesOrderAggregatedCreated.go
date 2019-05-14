package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesOrderAggregatedCreated :
type SalesOrderAggregatedCreated struct {
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
func (*SalesOrderAggregatedCreated) TableName() string {
	return "sales_order_aggregated_created"
}

// handler create
func initRoutersSalesOrderAggregatedCreated(r *gin.Engine, salesorderaggregatedcreated string) {
	route := r.Group("salesorderaggregatedcreated")
	route.GET("/", getSalesOrderAggregatedCreateds)
	route.GET("/:id", getSalesOrderAggregatedCreated)
	route.POST("/", createSalesOrderAggregatedCreated)
	route.PUT("/:id", updateSalesOrderAggregatedCreated)
	route.DELETE("/:id", deleteSalesOrderAggregatedCreated)
}

func getSalesOrderAggregatedCreateds(c *gin.Context) {
	var salesOrderAggregatedCreateds []SalesOrderAggregatedCreated
	if err := g.Find(&salesOrderAggregatedCreateds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderAggregatedCreateds)
	}
}

func getSalesOrderAggregatedCreated(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderAggregatedCreated SalesOrderAggregatedCreated
	if err := g.Where("id = ?", id).First(&salesOrderAggregatedCreated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderAggregatedCreated)
	}
}

func createSalesOrderAggregatedCreated(c *gin.Context) {
	var salesOrderAggregatedCreated SalesOrderAggregatedCreated
	c.BindJSON(&salesOrderAggregatedCreated)
	g.Create(&salesOrderAggregatedCreated)
	c.JSON(200, salesOrderAggregatedCreated)
}

func updateSalesOrderAggregatedCreated(c *gin.Context) {
	var salesOrderAggregatedCreated SalesOrderAggregatedCreated
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesOrderAggregatedCreated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesOrderAggregatedCreated)
	g.Save(&salesOrderAggregatedCreated)
	c.JSON(200, salesOrderAggregatedCreated)
}
func deleteSalesOrderAggregatedCreated(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderAggregatedCreated SalesOrderAggregatedCreated
	d := g.Where("id = ?", id).Delete(&salesOrderAggregatedCreated)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
