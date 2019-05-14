package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ReviewEntitySummary :
type ReviewEntitySummary struct {
	PrimaryId     uint64 `gorm:"column:primary_id;primary_key" form:"primary_id;primary_key" json:"primary_id;primary_key"`
	EntityPkValue uint64 `gorm:"column:entity_pk_value" form:"entity_pk_value" json:"entity_pk_value"`
	EntityType    uint16 `gorm:"column:entity_type" form:"entity_type" json:"entity_type"`
	ReviewsCount  uint16 `gorm:"column:reviews_count" form:"reviews_count" json:"reviews_count"`
	RatingSummary uint16 `gorm:"column:rating_summary" form:"rating_summary" json:"rating_summary"`
	StoreId       uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
}

// TableName :
func (*ReviewEntitySummary) TableName() string {
	return "review_entity_summary"
}

// handler create
func initRoutersReviewEntitySummary(r *gin.Engine, reviewentitysummary string) {
	route := r.Group("reviewentitysummary")
	route.GET("/", getReviewEntitySummarys)
	route.GET("/:id", getReviewEntitySummary)
	route.POST("/", createReviewEntitySummary)
	route.PUT("/:id", updateReviewEntitySummary)
	route.DELETE("/:id", deleteReviewEntitySummary)
}

func getReviewEntitySummarys(c *gin.Context) {
	var reviewEntitySummarys []ReviewEntitySummary
	if err := g.Find(&reviewEntitySummarys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviewEntitySummarys)
	}
}

func getReviewEntitySummary(c *gin.Context) {
	id := c.Params.ByName("id")
	var reviewEntitySummary ReviewEntitySummary
	if err := g.Where("id = ?", id).First(&reviewEntitySummary).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviewEntitySummary)
	}
}

func createReviewEntitySummary(c *gin.Context) {
	var reviewEntitySummary ReviewEntitySummary
	c.BindJSON(&reviewEntitySummary)
	g.Create(&reviewEntitySummary)
	c.JSON(200, reviewEntitySummary)
}

func updateReviewEntitySummary(c *gin.Context) {
	var reviewEntitySummary ReviewEntitySummary
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reviewEntitySummary).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reviewEntitySummary)
	g.Save(&reviewEntitySummary)
	c.JSON(200, reviewEntitySummary)
}
func deleteReviewEntitySummary(c *gin.Context) {
	id := c.Params.ByName("id")
	var reviewEntitySummary ReviewEntitySummary
	d := g.Where("id = ?", id).Delete(&reviewEntitySummary)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
