package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesRecurringProfileOrder :
type SalesRecurringProfileOrder struct {
	LinkId    uint16 `gorm:"column:link_id;primary_key" form:"link_id;primary_key" json:"link_id;primary_key"`
	ProfileId uint16 `gorm:"column:profile_id" form:"profile_id" json:"profile_id"`
	OrderId   uint16 `gorm:"column:order_id" form:"order_id" json:"order_id"`
}

// TableName :
func (*SalesRecurringProfileOrder) TableName() string {
	return "sales_recurring_profile_order"
}

// handler create
func initRoutersSalesRecurringProfileOrder(r *gin.Engine, salesrecurringprofileorder string) {
	route := r.Group("salesrecurringprofileorder")
	route.GET("/", getSalesRecurringProfileOrders)
	route.GET("/:id", getSalesRecurringProfileOrder)
	route.POST("/", createSalesRecurringProfileOrder)
	route.PUT("/:id", updateSalesRecurringProfileOrder)
	route.DELETE("/:id", deleteSalesRecurringProfileOrder)
}

func getSalesRecurringProfileOrders(c *gin.Context) {
	var salesRecurringProfileOrders []SalesRecurringProfileOrder
	if err := g.Find(&salesRecurringProfileOrders).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesRecurringProfileOrders)
	}
}

func getSalesRecurringProfileOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesRecurringProfileOrder SalesRecurringProfileOrder
	if err := g.Where("id = ?", id).First(&salesRecurringProfileOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesRecurringProfileOrder)
	}
}

func createSalesRecurringProfileOrder(c *gin.Context) {
	var salesRecurringProfileOrder SalesRecurringProfileOrder
	c.BindJSON(&salesRecurringProfileOrder)
	g.Create(&salesRecurringProfileOrder)
	c.JSON(200, salesRecurringProfileOrder)
}

func updateSalesRecurringProfileOrder(c *gin.Context) {
	var salesRecurringProfileOrder SalesRecurringProfileOrder
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesRecurringProfileOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesRecurringProfileOrder)
	g.Save(&salesRecurringProfileOrder)
	c.JSON(200, salesRecurringProfileOrder)
}
func deleteSalesRecurringProfileOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesRecurringProfileOrder SalesRecurringProfileOrder
	d := g.Where("id = ?", id).Delete(&salesRecurringProfileOrder)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
