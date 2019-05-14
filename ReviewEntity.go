package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ReviewEntity :
type ReviewEntity struct {
	EntityId   uint16 `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	EntityCode string `gorm:"column:entity_code" form:"entity_code" json:"entity_code"`
}

// TableName :
func (*ReviewEntity) TableName() string {
	return "review_entity"
}

// handler create
func initRoutersReviewEntity(r *gin.Engine, reviewentity string) {
	route := r.Group("reviewentity")
	route.GET("/", getReviewEntitys)
	route.GET("/:id", getReviewEntity)
	route.POST("/", createReviewEntity)
	route.PUT("/:id", updateReviewEntity)
	route.DELETE("/:id", deleteReviewEntity)
}

func getReviewEntitys(c *gin.Context) {
	var reviewEntitys []ReviewEntity
	if err := g.Find(&reviewEntitys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviewEntitys)
	}
}

func getReviewEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var reviewEntity ReviewEntity
	if err := g.Where("id = ?", id).First(&reviewEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviewEntity)
	}
}

func createReviewEntity(c *gin.Context) {
	var reviewEntity ReviewEntity
	c.BindJSON(&reviewEntity)
	g.Create(&reviewEntity)
	c.JSON(200, reviewEntity)
}

func updateReviewEntity(c *gin.Context) {
	var reviewEntity ReviewEntity
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reviewEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reviewEntity)
	g.Save(&reviewEntity)
	c.JSON(200, reviewEntity)
}
func deleteReviewEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var reviewEntity ReviewEntity
	d := g.Where("id = ?", id).Delete(&reviewEntity)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
