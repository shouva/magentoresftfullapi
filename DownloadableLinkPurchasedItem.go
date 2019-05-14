package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// DownloadableLinkPurchasedItem :
type DownloadableLinkPurchasedItem struct {
	ItemId                  uint16     `gorm:"column:item_id;primary_key" form:"item_id;primary_key" json:"item_id;primary_key"`
	PurchasedId             uint16     `gorm:"column:purchased_id" form:"purchased_id" json:"purchased_id"`
	OrderItemId             uint16     `gorm:"column:order_item_id" form:"order_item_id" json:"order_item_id"`
	ProductId               uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	LinkHash                string     `gorm:"column:link_hash" form:"link_hash" json:"link_hash"`
	NumberOfDownloadsBought uint16     `gorm:"column:number_of_downloads_bought" form:"number_of_downloads_bought" json:"number_of_downloads_bought"`
	NumberOfDownloadsUsed   uint16     `gorm:"column:number_of_downloads_used" form:"number_of_downloads_used" json:"number_of_downloads_used"`
	LinkId                  uint16     `gorm:"column:link_id" form:"link_id" json:"link_id"`
	LinkTitle               string     `gorm:"column:link_title" form:"link_title" json:"link_title"`
	IsShareable             uint16     `gorm:"column:is_shareable" form:"is_shareable" json:"is_shareable"`
	LinkUrl                 string     `gorm:"column:link_url" form:"link_url" json:"link_url"`
	LinkFile                string     `gorm:"column:link_file" form:"link_file" json:"link_file"`
	LinkType                string     `gorm:"column:link_type" form:"link_type" json:"link_type"`
	Status                  string     `gorm:"column:status" form:"status" json:"status"`
	CreatedAt               *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt               *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
}

// TableName :
func (*DownloadableLinkPurchasedItem) TableName() string {
	return "downloadable_link_purchased_item"
}

// handler create
func initRoutersDownloadableLinkPurchasedItem(r *gin.Engine, downloadablelinkpurchaseditem string) {
	route := r.Group("downloadablelinkpurchaseditem")
	route.GET("/", getDownloadableLinkPurchasedItems)
	route.GET("/:id", getDownloadableLinkPurchasedItem)
	route.POST("/", createDownloadableLinkPurchasedItem)
	route.PUT("/:id", updateDownloadableLinkPurchasedItem)
	route.DELETE("/:id", deleteDownloadableLinkPurchasedItem)
}

func getDownloadableLinkPurchasedItems(c *gin.Context) {
	var downloadableLinkPurchasedItems []DownloadableLinkPurchasedItem
	if err := g.Find(&downloadableLinkPurchasedItems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableLinkPurchasedItems)
	}
}

func getDownloadableLinkPurchasedItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableLinkPurchasedItem DownloadableLinkPurchasedItem
	if err := g.Where("id = ?", id).First(&downloadableLinkPurchasedItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableLinkPurchasedItem)
	}
}

func createDownloadableLinkPurchasedItem(c *gin.Context) {
	var downloadableLinkPurchasedItem DownloadableLinkPurchasedItem
	c.BindJSON(&downloadableLinkPurchasedItem)
	g.Create(&downloadableLinkPurchasedItem)
	c.JSON(200, downloadableLinkPurchasedItem)
}

func updateDownloadableLinkPurchasedItem(c *gin.Context) {
	var downloadableLinkPurchasedItem DownloadableLinkPurchasedItem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&downloadableLinkPurchasedItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&downloadableLinkPurchasedItem)
	g.Save(&downloadableLinkPurchasedItem)
	c.JSON(200, downloadableLinkPurchasedItem)
}
func deleteDownloadableLinkPurchasedItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableLinkPurchasedItem DownloadableLinkPurchasedItem
	d := g.Where("id = ?", id).Delete(&downloadableLinkPurchasedItem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
