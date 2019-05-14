package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// WishlistItem :
type WishlistItem struct {
	WishlistItemId uint16     `gorm:"column:wishlist_item_id;primary_key" form:"wishlist_item_id;primary_key" json:"wishlist_item_id;primary_key"`
	WishlistId     uint16     `gorm:"column:wishlist_id" form:"wishlist_id" json:"wishlist_id"`
	ProductId      uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	StoreId        uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	AddedAt        *time.Time `gorm:"column:added_at" form:"added_at" json:"added_at"`
	Description    string     `gorm:"column:description" form:"description" json:"description"`
	Qty            float64    `gorm:"column:qty" form:"qty" json:"qty"`
}

// TableName :
func (*WishlistItem) TableName() string {
	return "wishlist_item"
}

// handler create
func initRoutersWishlistItem(r *gin.Engine, wishlistitem string) {
	route := r.Group("wishlistitem")
	route.GET("/", getWishlistItems)
	route.GET("/:id", getWishlistItem)
	route.POST("/", createWishlistItem)
	route.PUT("/:id", updateWishlistItem)
	route.DELETE("/:id", deleteWishlistItem)
}

func getWishlistItems(c *gin.Context) {
	var wishlistItems []WishlistItem
	if err := g.Find(&wishlistItems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, wishlistItems)
	}
}

func getWishlistItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var wishlistItem WishlistItem
	if err := g.Where("id = ?", id).First(&wishlistItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, wishlistItem)
	}
}

func createWishlistItem(c *gin.Context) {
	var wishlistItem WishlistItem
	c.BindJSON(&wishlistItem)
	g.Create(&wishlistItem)
	c.JSON(200, wishlistItem)
}

func updateWishlistItem(c *gin.Context) {
	var wishlistItem WishlistItem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&wishlistItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&wishlistItem)
	g.Save(&wishlistItem)
	c.JSON(200, wishlistItem)
}
func deleteWishlistItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var wishlistItem WishlistItem
	d := g.Where("id = ?", id).Delete(&wishlistItem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
