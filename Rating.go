package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Rating :
type Rating struct {
	RatingId   uint16 `gorm:"column:rating_id;primary_key" form:"rating_id;primary_key" json:"rating_id;primary_key"`
	EntityId   uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	RatingCode string `gorm:"column:rating_code" form:"rating_code" json:"rating_code"`
	Position   uint16 `gorm:"column:position" form:"position" json:"position"`
}

// TableName :
func (*Rating) TableName() string {
	return "rating"
}

// handler create
func initRoutersRating(r *gin.Engine, rating string) {
	route := r.Group("rating")
	route.GET("/", getRatings)
	route.GET("/:id", getRating)
	route.POST("/", createRating)
	route.PUT("/:id", updateRating)
	route.DELETE("/:id", deleteRating)
}

func getRatings(c *gin.Context) {
	var ratings []Rating
	if err := g.Find(&ratings).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratings)
	}
}

func getRating(c *gin.Context) {
	id := c.Params.ByName("id")
	var rating Rating
	if err := g.Where("id = ?", id).First(&rating).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, rating)
	}
}

func createRating(c *gin.Context) {
	var rating Rating
	c.BindJSON(&rating)
	g.Create(&rating)
	c.JSON(200, rating)
}

func updateRating(c *gin.Context) {
	var rating Rating
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&rating).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&rating)
	g.Save(&rating)
	c.JSON(200, rating)
}
func deleteRating(c *gin.Context) {
	id := c.Params.ByName("id")
	var rating Rating
	d := g.Where("id = ?", id).Delete(&rating)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
