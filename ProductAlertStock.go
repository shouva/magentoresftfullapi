package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ProductAlertStock :
type ProductAlertStock struct {
	AlertStockId uint16     `gorm:"column:alert_stock_id;primary_key" form:"alert_stock_id;primary_key" json:"alert_stock_id;primary_key"`
	CustomerId   uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	ProductId    uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	WebsiteId    uint16     `gorm:"column:website_id" form:"website_id" json:"website_id"`
	AddDate      *time.Time `gorm:"column:add_date" form:"add_date" json:"add_date"`
	SendDate     *time.Time `gorm:"column:send_date" form:"send_date" json:"send_date"`
	SendCount    uint16     `gorm:"column:send_count" form:"send_count" json:"send_count"`
	Status       uint16     `gorm:"column:status" form:"status" json:"status"`
}

// TableName :
func (*ProductAlertStock) TableName() string {
	return "product_alert_stock"
}

// handler create
func initRoutersProductAlertStock(r *gin.Engine, productalertstock string) {
	route := r.Group("productalertstock")
	route.GET("/", getProductAlertStocks)
	route.GET("/:id", getProductAlertStock)
	route.POST("/", createProductAlertStock)
	route.PUT("/:id", updateProductAlertStock)
	route.DELETE("/:id", deleteProductAlertStock)
}

func getProductAlertStocks(c *gin.Context) {
	var productAlertStocks []ProductAlertStock
	if err := g.Find(&productAlertStocks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, productAlertStocks)
	}
}

func getProductAlertStock(c *gin.Context) {
	id := c.Params.ByName("id")
	var productAlertStock ProductAlertStock
	if err := g.Where("id = ?", id).First(&productAlertStock).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, productAlertStock)
	}
}

func createProductAlertStock(c *gin.Context) {
	var productAlertStock ProductAlertStock
	c.BindJSON(&productAlertStock)
	g.Create(&productAlertStock)
	c.JSON(200, productAlertStock)
}

func updateProductAlertStock(c *gin.Context) {
	var productAlertStock ProductAlertStock
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&productAlertStock).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&productAlertStock)
	g.Save(&productAlertStock)
	c.JSON(200, productAlertStock)
}
func deleteProductAlertStock(c *gin.Context) {
	id := c.Params.ByName("id")
	var productAlertStock ProductAlertStock
	d := g.Where("id = ?", id).Delete(&productAlertStock)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
