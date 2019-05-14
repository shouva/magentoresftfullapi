package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RatingTitle :
type RatingTitle struct {
	RatingId uint16 `gorm:"column:rating_id;primary_key" form:"rating_id;primary_key" json:"rating_id;primary_key"`
	StoreId  uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Value    string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*RatingTitle) TableName() string {
	return "rating_title"
}

// handler create
func initRoutersRatingTitle(r *gin.Engine, ratingtitle string) {
	route := r.Group("ratingtitle")
	route.GET("/", getRatingTitles)
	route.GET("/:id", getRatingTitle)
	route.POST("/", createRatingTitle)
	route.PUT("/:id", updateRatingTitle)
	route.DELETE("/:id", deleteRatingTitle)
}

func getRatingTitles(c *gin.Context) {
	var ratingTitles []RatingTitle
	if err := g.Find(&ratingTitles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingTitles)
	}
}

func getRatingTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingTitle RatingTitle
	if err := g.Where("id = ?", id).First(&ratingTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingTitle)
	}
}

func createRatingTitle(c *gin.Context) {
	var ratingTitle RatingTitle
	c.BindJSON(&ratingTitle)
	g.Create(&ratingTitle)
	c.JSON(200, ratingTitle)
}

func updateRatingTitle(c *gin.Context) {
	var ratingTitle RatingTitle
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&ratingTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&ratingTitle)
	g.Save(&ratingTitle)
	c.JSON(200, ratingTitle)
}
func deleteRatingTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingTitle RatingTitle
	d := g.Where("id = ?", id).Delete(&ratingTitle)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
