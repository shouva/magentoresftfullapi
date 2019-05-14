package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RatingOptionVote :
type RatingOptionVote struct {
	VoteId        uint64 `gorm:"column:vote_id;primary_key" form:"vote_id;primary_key" json:"vote_id;primary_key"`
	OptionId      uint16 `gorm:"column:option_id" form:"option_id" json:"option_id"`
	RemoteIp      string `gorm:"column:remote_ip" form:"remote_ip" json:"remote_ip"`
	RemoteIpLong  []byte `gorm:"column:remote_ip_long" form:"remote_ip_long" json:"remote_ip_long"`
	CustomerId    uint16 `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	EntityPkValue uint64 `gorm:"column:entity_pk_value" form:"entity_pk_value" json:"entity_pk_value"`
	RatingId      uint16 `gorm:"column:rating_id" form:"rating_id" json:"rating_id"`
	ReviewId      uint64 `gorm:"column:review_id" form:"review_id" json:"review_id"`
	Percent       uint16 `gorm:"column:percent" form:"percent" json:"percent"`
	Value         uint16 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*RatingOptionVote) TableName() string {
	return "rating_option_vote"
}

// handler create
func initRoutersRatingOptionVote(r *gin.Engine, ratingoptionvote string) {
	route := r.Group("ratingoptionvote")
	route.GET("/", getRatingOptionVotes)
	route.GET("/:id", getRatingOptionVote)
	route.POST("/", createRatingOptionVote)
	route.PUT("/:id", updateRatingOptionVote)
	route.DELETE("/:id", deleteRatingOptionVote)
}

func getRatingOptionVotes(c *gin.Context) {
	var ratingOptionVotes []RatingOptionVote
	if err := g.Find(&ratingOptionVotes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingOptionVotes)
	}
}

func getRatingOptionVote(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingOptionVote RatingOptionVote
	if err := g.Where("id = ?", id).First(&ratingOptionVote).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ratingOptionVote)
	}
}

func createRatingOptionVote(c *gin.Context) {
	var ratingOptionVote RatingOptionVote
	c.BindJSON(&ratingOptionVote)
	g.Create(&ratingOptionVote)
	c.JSON(200, ratingOptionVote)
}

func updateRatingOptionVote(c *gin.Context) {
	var ratingOptionVote RatingOptionVote
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&ratingOptionVote).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&ratingOptionVote)
	g.Save(&ratingOptionVote)
	c.JSON(200, ratingOptionVote)
}
func deleteRatingOptionVote(c *gin.Context) {
	id := c.Params.ByName("id")
	var ratingOptionVote RatingOptionVote
	d := g.Where("id = ?", id).Delete(&ratingOptionVote)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
