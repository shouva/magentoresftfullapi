package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ProductAlertPrice :
type ProductAlertPrice struct {
	AlertPriceId uint16     `gorm:"column:alert_price_id;primary_key" form:"alert_price_id;primary_key" json:"alert_price_id;primary_key"`
	CustomerId   uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	ProductId    uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Price        float64    `gorm:"column:price" form:"price" json:"price"`
	WebsiteId    uint16     `gorm:"column:website_id" form:"website_id" json:"website_id"`
	AddDate      *time.Time `gorm:"column:add_date" form:"add_date" json:"add_date"`
	LastSendDate *time.Time `gorm:"column:last_send_date" form:"last_send_date" json:"last_send_date"`
	SendCount    uint16     `gorm:"column:send_count" form:"send_count" json:"send_count"`
	Status       uint16     `gorm:"column:status" form:"status" json:"status"`
}

// TableName :
func (*ProductAlertPrice) TableName() string {
	return "product_alert_price"
}

// handler create
func initRoutersProductAlertPrice(r *gin.Engine, productalertprice string) {
	route := r.Group("productalertprice")
	route.GET("/", getProductAlertPrices)
	route.GET("/:id", getProductAlertPrice)
	route.POST("/", createProductAlertPrice)
	route.PUT("/:id", updateProductAlertPrice)
	route.DELETE("/:id", deleteProductAlertPrice)
}

func getProductAlertPrices(c *gin.Context) {
	var productAlertPrices []ProductAlertPrice
	if err := g.Find(&productAlertPrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, productAlertPrices)
	}
}

func getProductAlertPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var productAlertPrice ProductAlertPrice
	if err := g.Where("id = ?", id).First(&productAlertPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, productAlertPrice)
	}
}

func createProductAlertPrice(c *gin.Context) {
	var productAlertPrice ProductAlertPrice
	c.BindJSON(&productAlertPrice)
	g.Create(&productAlertPrice)
	c.JSON(200, productAlertPrice)
}

func updateProductAlertPrice(c *gin.Context) {
	var productAlertPrice ProductAlertPrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&productAlertPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&productAlertPrice)
	g.Save(&productAlertPrice)
	c.JSON(200, productAlertPrice)
}
func deleteProductAlertPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var productAlertPrice ProductAlertPrice
	d := g.Where("id = ?", id).Delete(&productAlertPrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
