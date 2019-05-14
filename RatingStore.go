package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RatingStore :
type RatingStore struct {
	RatingId uint16 `gorm:"column:rating_id;primary_key" form:"rating_id;primary_key" json:"rating_id;primary_key"`
	StoreId  uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
}

// TableName :
func (*RatingStore) TableName() string {
	return "rating_store"
}

// handler create
func initRoutersRatingStore(r *gin.Engine, ratingstore string) {
	route := r.Group("ratingstore")
	route.GET("/", getRatingStores)
	route.GET("/:id", getRatingStore)
	route.POST("/", createRatingStore)
	route.PUT("/:id", updateRatingStore)
	route.DELETE("/:id", deleteRatingStore)
}

func getRatingStores(c *gin.Context) {
	var ratingStores []RatingStore
	if err := g.Find(&ratingStores).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingStores)
	}
}

func getRatingStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingStore RatingStore
	if err := g.Where("id = ?", id).First(&ratingStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingStore)
	}
}

func createRatingStore(c *gin.Context) {
	var ratingStore RatingStore
	c.BindJSON(&ratingStore)
	g.Create(&ratingStore)
	c.JSON(200, ratingStore)
}

func updateRatingStore(c *gin.Context) {
	var ratingStore RatingStore
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&ratingStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&ratingStore)
	g.Save(&ratingStore)
	c.JSON(200, ratingStore)
}
func deleteRatingStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingStore RatingStore
	d := g.Where("id = ?", id).Delete(&ratingStore)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
