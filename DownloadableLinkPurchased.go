package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// DownloadableLinkPurchased :
type DownloadableLinkPurchased struct {
	PurchasedId      uint16     `gorm:"column:purchased_id;primary_key" form:"purchased_id;primary_key" json:"purchased_id;primary_key"`
	OrderId          uint16     `gorm:"column:order_id" form:"order_id" json:"order_id"`
	OrderIncrementId string     `gorm:"column:order_increment_id" form:"order_increment_id" json:"order_increment_id"`
	OrderItemId      uint16     `gorm:"column:order_item_id" form:"order_item_id" json:"order_item_id"`
	CreatedAt        *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	CustomerId       uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	ProductName      string     `gorm:"column:product_name" form:"product_name" json:"product_name"`
	ProductSku       string     `gorm:"column:product_sku" form:"product_sku" json:"product_sku"`
	LinkSectionTitle string     `gorm:"column:link_section_title" form:"link_section_title" json:"link_section_title"`
}

// TableName :
func (*DownloadableLinkPurchased) TableName() string {
	return "downloadable_link_purchased"
}

// handler create
func initRoutersDownloadableLinkPurchased(r *gin.Engine, downloadablelinkpurchased string) {
	route := r.Group("downloadablelinkpurchased")
	route.GET("/", getDownloadableLinkPurchaseds)
	route.GET("/:id", getDownloadableLinkPurchased)
	route.POST("/", createDownloadableLinkPurchased)
	route.PUT("/:id", updateDownloadableLinkPurchased)
	route.DELETE("/:id", deleteDownloadableLinkPurchased)
}

func getDownloadableLinkPurchaseds(c *gin.Context) {
	var downloadableLinkPurchaseds []DownloadableLinkPurchased
	if err := g.Find(&downloadableLinkPurchaseds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableLinkPurchaseds)
	}
}

func getDownloadableLinkPurchased(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableLinkPurchased DownloadableLinkPurchased
	if err := g.Where("id = ?", id).First(&downloadableLinkPurchased).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableLinkPurchased)
	}
}

func createDownloadableLinkPurchased(c *gin.Context) {
	var downloadableLinkPurchased DownloadableLinkPurchased
	c.BindJSON(&downloadableLinkPurchased)
	g.Create(&downloadableLinkPurchased)
	c.JSON(200, downloadableLinkPurchased)
}

func updateDownloadableLinkPurchased(c *gin.Context) {
	var downloadableLinkPurchased DownloadableLinkPurchased
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&downloadableLinkPurchased).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&downloadableLinkPurchased)
	g.Save(&downloadableLinkPurchased)
	c.JSON(200, downloadableLinkPurchased)
}
func deleteDownloadableLinkPurchased(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableLinkPurchased DownloadableLinkPurchased
	d := g.Where("id = ?", id).Delete(&downloadableLinkPurchased)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
