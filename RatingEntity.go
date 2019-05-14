package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RatingEntity :
type RatingEntity struct {
	EntityId   uint16 `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	EntityCode string `gorm:"column:entity_code" form:"entity_code" json:"entity_code"`
}

// TableName :
func (*RatingEntity) TableName() string {
	return "rating_entity"
}

// handler create
func initRoutersRatingEntity(r *gin.Engine, ratingentity string) {
	route := r.Group("ratingentity")
	route.GET("/", getRatingEntitys)
	route.GET("/:id", getRatingEntity)
	route.POST("/", createRatingEntity)
	route.PUT("/:id", updateRatingEntity)
	route.DELETE("/:id", deleteRatingEntity)
}

func getRatingEntitys(c *gin.Context) {
	var ratingEntitys []RatingEntity
	if err := g.Find(&ratingEntitys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingEntitys)
	}
}

func getRatingEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingEntity RatingEntity
	if err := g.Where("id = ?", id).First(&ratingEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingEntity)
	}
}

func createRatingEntity(c *gin.Context) {
	var ratingEntity RatingEntity
	c.BindJSON(&ratingEntity)
	g.Create(&ratingEntity)
	c.JSON(200, ratingEntity)
}

func updateRatingEntity(c *gin.Context) {
	var ratingEntity RatingEntity
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&ratingEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&ratingEntity)
	g.Save(&ratingEntity)
	c.JSON(200, ratingEntity)
}
func deleteRatingEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingEntity RatingEntity
	d := g.Where("id = ?", id).Delete(&ratingEntity)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
