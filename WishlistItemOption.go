package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// WishlistItemOption :
type WishlistItemOption struct {
	OptionId       uint16 `gorm:"column:option_id;primary_key" form:"option_id;primary_key" json:"option_id;primary_key"`
	WishlistItemId uint16 `gorm:"column:wishlist_item_id" form:"wishlist_item_id" json:"wishlist_item_id"`
	ProductId      uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Code           string `gorm:"column:code" form:"code" json:"code"`
	Value          string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*WishlistItemOption) TableName() string {
	return "wishlist_item_option"
}

// handler create
func initRoutersWishlistItemOption(r *gin.Engine, wishlistitemoption string) {
	route := r.Group("wishlistitemoption")
	route.GET("/", getWishlistItemOptions)
	route.GET("/:id", getWishlistItemOption)
	route.POST("/", createWishlistItemOption)
	route.PUT("/:id", updateWishlistItemOption)
	route.DELETE("/:id", deleteWishlistItemOption)
}

func getWishlistItemOptions(c *gin.Context) {
	var wishlistItemOptions []WishlistItemOption
	if err := g.Find(&wishlistItemOptions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, wishlistItemOptions)
	}
}

func getWishlistItemOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var wishlistItemOption WishlistItemOption
	if err := g.Where("id = ?", id).First(&wishlistItemOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, wishlistItemOption)
	}
}

func createWishlistItemOption(c *gin.Context) {
	var wishlistItemOption WishlistItemOption
	c.BindJSON(&wishlistItemOption)
	g.Create(&wishlistItemOption)
	c.JSON(200, wishlistItemOption)
}

func updateWishlistItemOption(c *gin.Context) {
	var wishlistItemOption WishlistItemOption
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&wishlistItemOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&wishlistItemOption)
	g.Save(&wishlistItemOption)
	c.JSON(200, wishlistItemOption)
}
func deleteWishlistItemOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var wishlistItemOption WishlistItemOption
	d := g.Where("id = ?", id).Delete(&wishlistItemOption)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
