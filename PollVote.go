package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// PollVote :
type PollVote struct {
	VoteId       uint16     `gorm:"column:vote_id;primary_key" form:"vote_id;primary_key" json:"vote_id;primary_key"`
	PollId       uint16     `gorm:"column:poll_id" form:"poll_id" json:"poll_id"`
	PollAnswerId uint16     `gorm:"column:poll_answer_id" form:"poll_answer_id" json:"poll_answer_id"`
	IpAddress    []byte     `gorm:"column:ip_address" form:"ip_address" json:"ip_address"`
	CustomerId   uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	VoteTime     *time.Time `gorm:"column:vote_time" form:"vote_time" json:"vote_time"`
}

// TableName :
func (*PollVote) TableName() string {
	return "poll_vote"
}

// handler create
func initRoutersPollVote(r *gin.Engine, pollvote string) {
	route := r.Group("pollvote")
	route.GET("/", getPollVotes)
	route.GET("/:id", getPollVote)
	route.POST("/", createPollVote)
	route.PUT("/:id", updatePollVote)
	route.DELETE("/:id", deletePollVote)
}

func getPollVotes(c *gin.Context) {
	var pollVotes []PollVote
	if err := g.Find(&pollVotes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pollVotes)
	}
}

func getPollVote(c *gin.Context) {
	id := c.Params.ByName("id")
	var pollVote PollVote
	if err := g.Where("id = ?", id).First(&pollVote).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pollVote)
	}
}

func createPollVote(c *gin.Context) {
	var pollVote PollVote
	c.BindJSON(&pollVote)
	g.Create(&pollVote)
	c.JSON(200, pollVote)
}

func updatePollVote(c *gin.Context) {
	var pollVote PollVote
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&pollVote).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&pollVote)
	g.Save(&pollVote)
	c.JSON(200, pollVote)
}
func deletePollVote(c *gin.Context) {
	id := c.Params.ByName("id")
	var pollVote PollVote
	d := g.Where("id = ?", id).Delete(&pollVote)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
