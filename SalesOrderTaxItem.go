package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesOrderTaxItem :
type SalesOrderTaxItem struct {
	TaxItemId  uint16  `gorm:"column:tax_item_id;primary_key" form:"tax_item_id;primary_key" json:"tax_item_id;primary_key"`
	TaxId      uint16  `gorm:"column:tax_id" form:"tax_id" json:"tax_id"`
	ItemId     uint16  `gorm:"column:item_id" form:"item_id" json:"item_id"`
	TaxPercent float64 `gorm:"column:tax_percent" form:"tax_percent" json:"tax_percent"`
}

// TableName :
func (*SalesOrderTaxItem) TableName() string {
	return "sales_order_tax_item"
}

// handler create
func initRoutersSalesOrderTaxItem(r *gin.Engine, salesordertaxitem string) {
	route := r.Group("salesordertaxitem")
	route.GET("/", getSalesOrderTaxItems)
	route.GET("/:id", getSalesOrderTaxItem)
	route.POST("/", createSalesOrderTaxItem)
	route.PUT("/:id", updateSalesOrderTaxItem)
	route.DELETE("/:id", deleteSalesOrderTaxItem)
}

func getSalesOrderTaxItems(c *gin.Context) {
	var salesOrderTaxItems []SalesOrderTaxItem
	if err := g.Find(&salesOrderTaxItems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderTaxItems)
	}
}

func getSalesOrderTaxItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderTaxItem SalesOrderTaxItem
	if err := g.Where("id = ?", id).First(&salesOrderTaxItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderTaxItem)
	}
}

func createSalesOrderTaxItem(c *gin.Context) {
	var salesOrderTaxItem SalesOrderTaxItem
	c.BindJSON(&salesOrderTaxItem)
	g.Create(&salesOrderTaxItem)
	c.JSON(200, salesOrderTaxItem)
}

func updateSalesOrderTaxItem(c *gin.Context) {
	var salesOrderTaxItem SalesOrderTaxItem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesOrderTaxItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesOrderTaxItem)
	g.Save(&salesOrderTaxItem)
	c.JSON(200, salesOrderTaxItem)
}
func deleteSalesOrderTaxItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderTaxItem SalesOrderTaxItem
	d := g.Where("id = ?", id).Delete(&salesOrderTaxItem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
