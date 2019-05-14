package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesOrderTax :
type SalesOrderTax struct {
	TaxId          uint16  `gorm:"column:tax_id;primary_key" form:"tax_id;primary_key" json:"tax_id;primary_key"`
	OrderId        uint16  `gorm:"column:order_id" form:"order_id" json:"order_id"`
	Code           string  `gorm:"column:code" form:"code" json:"code"`
	Title          string  `gorm:"column:title" form:"title" json:"title"`
	Percent        float64 `gorm:"column:percent" form:"percent" json:"percent"`
	Amount         float64 `gorm:"column:amount" form:"amount" json:"amount"`
	Priority       uint16  `gorm:"column:priority" form:"priority" json:"priority"`
	Position       uint16  `gorm:"column:position" form:"position" json:"position"`
	BaseAmount     float64 `gorm:"column:base_amount" form:"base_amount" json:"base_amount"`
	Process        uint16  `gorm:"column:process" form:"process" json:"process"`
	BaseRealAmount float64 `gorm:"column:base_real_amount" form:"base_real_amount" json:"base_real_amount"`
	Hidden         uint16  `gorm:"column:hidden" form:"hidden" json:"hidden"`
}

// TableName :
func (*SalesOrderTax) TableName() string {
	return "sales_order_tax"
}

// handler create
func initRoutersSalesOrderTax(r *gin.Engine, salesordertax string) {
	route := r.Group("salesordertax")
	route.GET("/", getSalesOrderTaxs)
	route.GET("/:id", getSalesOrderTax)
	route.POST("/", createSalesOrderTax)
	route.PUT("/:id", updateSalesOrderTax)
	route.DELETE("/:id", deleteSalesOrderTax)
}

func getSalesOrderTaxs(c *gin.Context) {
	var salesOrderTaxs []SalesOrderTax
	if err := g.Find(&salesOrderTaxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderTaxs)
	}
}

func getSalesOrderTax(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderTax SalesOrderTax
	if err := g.Where("id = ?", id).First(&salesOrderTax).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderTax)
	}
}

func createSalesOrderTax(c *gin.Context) {
	var salesOrderTax SalesOrderTax
	c.BindJSON(&salesOrderTax)
	g.Create(&salesOrderTax)
	c.JSON(200, salesOrderTax)
}

func updateSalesOrderTax(c *gin.Context) {
	var salesOrderTax SalesOrderTax
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesOrderTax).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesOrderTax)
	g.Save(&salesOrderTax)
	c.JSON(200, salesOrderTax)
}
func deleteSalesOrderTax(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderTax SalesOrderTax
	d := g.Where("id = ?", id).Delete(&salesOrderTax)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
