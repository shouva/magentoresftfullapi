package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RatingOptionVoteAggregated :
type RatingOptionVoteAggregated struct {
	PrimaryId       uint16 `gorm:"column:primary_id;primary_key" form:"primary_id;primary_key" json:"primary_id;primary_key"`
	RatingId        uint16 `gorm:"column:rating_id" form:"rating_id" json:"rating_id"`
	EntityPkValue   uint64 `gorm:"column:entity_pk_value" form:"entity_pk_value" json:"entity_pk_value"`
	VoteCount       uint16 `gorm:"column:vote_count" form:"vote_count" json:"vote_count"`
	VoteValueSum    uint16 `gorm:"column:vote_value_sum" form:"vote_value_sum" json:"vote_value_sum"`
	Percent         uint16 `gorm:"column:percent" form:"percent" json:"percent"`
	PercentApproved uint16 `gorm:"column:percent_approved" form:"percent_approved" json:"percent_approved"`
	StoreId         uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
}

// TableName :
func (*RatingOptionVoteAggregated) TableName() string {
	return "rating_option_vote_aggregated"
}

// handler create
func initRoutersRatingOptionVoteAggregated(r *gin.Engine, ratingoptionvoteaggregated string) {
	route := r.Group("ratingoptionvoteaggregated")
	route.GET("/", getRatingOptionVoteAggregateds)
	route.GET("/:id", getRatingOptionVoteAggregated)
	route.POST("/", createRatingOptionVoteAggregated)
	route.PUT("/:id", updateRatingOptionVoteAggregated)
	route.DELETE("/:id", deleteRatingOptionVoteAggregated)
}

func getRatingOptionVoteAggregateds(c *gin.Context) {
	var ratingOptionVoteAggregateds []RatingOptionVoteAggregated
	if err := g.Find(&ratingOptionVoteAggregateds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingOptionVoteAggregateds)
	}
}

func getRatingOptionVoteAggregated(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingOptionVoteAggregated RatingOptionVoteAggregated
	if err := g.Where("id = ?", id).First(&ratingOptionVoteAggregated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingOptionVoteAggregated)
	}
}

func createRatingOptionVoteAggregated(c *gin.Context) {
	var ratingOptionVoteAggregated RatingOptionVoteAggregated
	c.BindJSON(&ratingOptionVoteAggregated)
	g.Create(&ratingOptionVoteAggregated)
	c.JSON(200, ratingOptionVoteAggregated)
}

func updateRatingOptionVoteAggregated(c *gin.Context) {
	var ratingOptionVoteAggregated RatingOptionVoteAggregated
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&ratingOptionVoteAggregated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&ratingOptionVoteAggregated)
	g.Save(&ratingOptionVoteAggregated)
	c.JSON(200, ratingOptionVoteAggregated)
}
func deleteRatingOptionVoteAggregated(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingOptionVoteAggregated RatingOptionVoteAggregated
	d := g.Where("id = ?", id).Delete(&ratingOptionVoteAggregated)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
