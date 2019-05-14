package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Review :
type Review struct {
	ReviewId      uint64     `gorm:"column:review_id;primary_key" form:"review_id;primary_key" json:"review_id;primary_key"`
	CreatedAt     *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	EntityId      uint16     `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	EntityPkValue uint16     `gorm:"column:entity_pk_value" form:"entity_pk_value" json:"entity_pk_value"`
	StatusId      uint16     `gorm:"column:status_id" form:"status_id" json:"status_id"`
}

// TableName :
func (*Review) TableName() string {
	return "review"
}

// handler create
func initRoutersReview(r *gin.Engine, review string) {
	route := r.Group("review")
	route.GET("/", getReviews)
	route.GET("/:id", getReview)
	route.POST("/", createReview)
	route.PUT("/:id", updateReview)
	route.DELETE("/:id", deleteReview)
}

func getReviews(c *gin.Context) {
	var reviews []Review
	if err := g.Find(&reviews).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviews)
	}
}

func getReview(c *gin.Context) {
	id := c.Params.ByName("id")
	var review Review
	if err := g.Where("id = ?", id).First(&review).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, review)
	}
}

func createReview(c *gin.Context) {
	var review Review
	c.BindJSON(&review)
	g.Create(&review)
	c.JSON(200, review)
}

func updateReview(c *gin.Context) {
	var review Review
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&review).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&review)
	g.Save(&review)
	c.JSON(200, review)
}
func deleteReview(c *gin.Context) {
	id := c.Params.ByName("id")
	var review Review
	d := g.Where("id = ?", id).Delete(&review)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
