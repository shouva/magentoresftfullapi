package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// TaxOrderAggregatedUpdated :
type TaxOrderAggregatedUpdated struct {
	ID               uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Period           *time.Time `gorm:"column:period" form:"period" json:"period"`
	StoreId          uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Code             string     `gorm:"column:code" form:"code" json:"code"`
	OrderStatus      string     `gorm:"column:order_status" form:"order_status" json:"order_status"`
	Percent          float32    `gorm:"column:percent" form:"percent" json:"percent"`
	OrdersCount      uint16     `gorm:"column:orders_count" form:"orders_count" json:"orders_count"`
	TaxBaseAmountSum float32    `gorm:"column:tax_base_amount_sum" form:"tax_base_amount_sum" json:"tax_base_amount_sum"`
}

// TableName :
func (*TaxOrderAggregatedUpdated) TableName() string {
	return "tax_order_aggregated_updated"
}

// handler create
func initRoutersTaxOrderAggregatedUpdated(r *gin.Engine, taxorderaggregatedupdated string) {
	route := r.Group("taxorderaggregatedupdated")
	route.GET("/", getTaxOrderAggregatedUpdateds)
	route.GET("/:id", getTaxOrderAggregatedUpdated)
	route.POST("/", createTaxOrderAggregatedUpdated)
	route.PUT("/:id", updateTaxOrderAggregatedUpdated)
	route.DELETE("/:id", deleteTaxOrderAggregatedUpdated)
}

func getTaxOrderAggregatedUpdateds(c *gin.Context) {
	var taxOrderAggregatedUpdateds []TaxOrderAggregatedUpdated
	if err := g.Find(&taxOrderAggregatedUpdateds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxOrderAggregatedUpdateds)
	}
}

func getTaxOrderAggregatedUpdated(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxOrderAggregatedUpdated TaxOrderAggregatedUpdated
	if err := g.Where("id = ?", id).First(&taxOrderAggregatedUpdated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxOrderAggregatedUpdated)
	}
}

func createTaxOrderAggregatedUpdated(c *gin.Context) {
	var taxOrderAggregatedUpdated TaxOrderAggregatedUpdated
	c.BindJSON(&taxOrderAggregatedUpdated)
	g.Create(&taxOrderAggregatedUpdated)
	c.JSON(200, taxOrderAggregatedUpdated)
}

func updateTaxOrderAggregatedUpdated(c *gin.Context) {
	var taxOrderAggregatedUpdated TaxOrderAggregatedUpdated
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&taxOrderAggregatedUpdated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&taxOrderAggregatedUpdated)
	g.Save(&taxOrderAggregatedUpdated)
	c.JSON(200, taxOrderAggregatedUpdated)
}
func deleteTaxOrderAggregatedUpdated(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxOrderAggregatedUpdated TaxOrderAggregatedUpdated
	d := g.Where("id = ?", id).Delete(&taxOrderAggregatedUpdated)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
