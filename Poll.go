package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Poll :
type Poll struct {
	PollId         uint16     `gorm:"column:poll_id;primary_key" form:"poll_id;primary_key" json:"poll_id;primary_key"`
	PollTitle      string     `gorm:"column:poll_title" form:"poll_title" json:"poll_title"`
	VotesCount     uint16     `gorm:"column:votes_count" form:"votes_count" json:"votes_count"`
	StoreId        uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	DatePosted     *time.Time `gorm:"column:date_posted" form:"date_posted" json:"date_posted"`
	DateClosed     *time.Time `gorm:"column:date_closed" form:"date_closed" json:"date_closed"`
	Active         uint16     `gorm:"column:active" form:"active" json:"active"`
	Closed         uint16     `gorm:"column:closed" form:"closed" json:"closed"`
	AnswersDisplay uint16     `gorm:"column:answers_display" form:"answers_display" json:"answers_display"`
}

// TableName :
func (*Poll) TableName() string {
	return "poll"
}

// handler create
func initRoutersPoll(r *gin.Engine, poll string) {
	route := r.Group("poll")
	route.GET("/", getPolls)
	route.GET("/:id", getPoll)
	route.POST("/", createPoll)
	route.PUT("/:id", updatePoll)
	route.DELETE("/:id", deletePoll)
}

func getPolls(c *gin.Context) {
	var polls []Poll
	if err := g.Find(&polls).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, polls)
	}
}

func getPoll(c *gin.Context) {
	id := c.Params.ByName("id")
	var poll Poll
	if err := g.Where("id = ?", id).First(&poll).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, poll)
	}
}

func createPoll(c *gin.Context) {
	var poll Poll
	c.BindJSON(&poll)
	g.Create(&poll)
	c.JSON(200, poll)
}

func updatePoll(c *gin.Context) {
	var poll Poll
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&poll).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&poll)
	g.Save(&poll)
	c.JSON(200, poll)
}
func deletePoll(c *gin.Context) {
	id := c.Params.ByName("id")
	var poll Poll
	d := g.Where("id = ?", id).Delete(&poll)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
