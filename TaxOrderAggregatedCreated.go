package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// TaxOrderAggregatedCreated :
type TaxOrderAggregatedCreated struct {
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
func (*TaxOrderAggregatedCreated) TableName() string {
	return "tax_order_aggregated_created"
}

// handler create
func initRoutersTaxOrderAggregatedCreated(r *gin.Engine, taxorderaggregatedcreated string) {
	route := r.Group("taxorderaggregatedcreated")
	route.GET("/", getTaxOrderAggregatedCreateds)
	route.GET("/:id", getTaxOrderAggregatedCreated)
	route.POST("/", createTaxOrderAggregatedCreated)
	route.PUT("/:id", updateTaxOrderAggregatedCreated)
	route.DELETE("/:id", deleteTaxOrderAggregatedCreated)
}

func getTaxOrderAggregatedCreateds(c *gin.Context) {
	var taxOrderAggregatedCreateds []TaxOrderAggregatedCreated
	if err := g.Find(&taxOrderAggregatedCreateds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxOrderAggregatedCreateds)
	}
}

func getTaxOrderAggregatedCreated(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxOrderAggregatedCreated TaxOrderAggregatedCreated
	if err := g.Where("id = ?", id).First(&taxOrderAggregatedCreated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxOrderAggregatedCreated)
	}
}

func createTaxOrderAggregatedCreated(c *gin.Context) {
	var taxOrderAggregatedCreated TaxOrderAggregatedCreated
	c.BindJSON(&taxOrderAggregatedCreated)
	g.Create(&taxOrderAggregatedCreated)
	c.JSON(200, taxOrderAggregatedCreated)
}

func updateTaxOrderAggregatedCreated(c *gin.Context) {
	var taxOrderAggregatedCreated TaxOrderAggregatedCreated
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&taxOrderAggregatedCreated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&taxOrderAggregatedCreated)
	g.Save(&taxOrderAggregatedCreated)
	c.JSON(200, taxOrderAggregatedCreated)
}
func deleteTaxOrderAggregatedCreated(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxOrderAggregatedCreated TaxOrderAggregatedCreated
	d := g.Where("id = ?", id).Delete(&taxOrderAggregatedCreated)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
