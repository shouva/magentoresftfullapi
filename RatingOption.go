package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RatingOption :
type RatingOption struct {
	OptionId uint16 `gorm:"column:option_id;primary_key" form:"option_id;primary_key" json:"option_id;primary_key"`
	RatingId uint16 `gorm:"column:rating_id" form:"rating_id" json:"rating_id"`
	Code     string `gorm:"column:code" form:"code" json:"code"`
	Value    uint16 `gorm:"column:value" form:"value" json:"value"`
	Position uint16 `gorm:"column:position" form:"position" json:"position"`
}

// TableName :
func (*RatingOption) TableName() string {
	return "rating_option"
}

// handler create
func initRoutersRatingOption(r *gin.Engine, ratingoption string) {
	route := r.Group("ratingoption")
	route.GET("/", getRatingOptions)
	route.GET("/:id", getRatingOption)
	route.POST("/", createRatingOption)
	route.PUT("/:id", updateRatingOption)
	route.DELETE("/:id", deleteRatingOption)
}

func getRatingOptions(c *gin.Context) {
	var ratingOptions []RatingOption
	if err := g.Find(&ratingOptions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingOptions)
	}
}

func getRatingOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingOption RatingOption
	if err := g.Where("id = ?", id).First(&ratingOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingOption)
	}
}

func createRatingOption(c *gin.Context) {
	var ratingOption RatingOption
	c.BindJSON(&ratingOption)
	g.Create(&ratingOption)
	c.JSON(200, ratingOption)
}

func updateRatingOption(c *gin.Context) {
	var ratingOption RatingOption
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&ratingOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&ratingOption)
	g.Save(&ratingOption)
	c.JSON(200, ratingOption)
}
func deleteRatingOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingOption RatingOption
	d := g.Where("id = ?", id).Delete(&ratingOption)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
