package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Wishlist :
type Wishlist struct {
	WishlistId  uint16     `gorm:"column:wishlist_id;primary_key" form:"wishlist_id;primary_key" json:"wishlist_id;primary_key"`
	CustomerId  uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	Shared      uint16     `gorm:"column:shared" form:"shared" json:"shared"`
	SharingCode string     `gorm:"column:sharing_code" form:"sharing_code" json:"sharing_code"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
}

// TableName :
func (*Wishlist) TableName() string {
	return "wishlist"
}

// handler create
func initRoutersWishlist(r *gin.Engine, wishlist string) {
	route := r.Group("wishlist")
	route.GET("/", getWishlists)
	route.GET("/:id", getWishlist)
	route.POST("/", createWishlist)
	route.PUT("/:id", updateWishlist)
	route.DELETE("/:id", deleteWishlist)
}

func getWishlists(c *gin.Context) {
	var wishlists []Wishlist
	if err := g.Find(&wishlists).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, wishlists)
	}
}

func getWishlist(c *gin.Context) {
	id := c.Params.ByName("id")
	var wishlist Wishlist
	if err := g.Where("id = ?", id).First(&wishlist).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, wishlist)
	}
}

func createWishlist(c *gin.Context) {
	var wishlist Wishlist
	c.BindJSON(&wishlist)
	g.Create(&wishlist)
	c.JSON(200, wishlist)
}

func updateWishlist(c *gin.Context) {
	var wishlist Wishlist
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&wishlist).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&wishlist)
	g.Save(&wishlist)
	c.JSON(200, wishlist)
}
func deleteWishlist(c *gin.Context) {
	id := c.Params.ByName("id")
	var wishlist Wishlist
	d := g.Where("id = ?", id).Delete(&wishlist)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
