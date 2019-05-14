package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// PollAnswer :
type PollAnswer struct {
	AnswerId    uint16 `gorm:"column:answer_id;primary_key" form:"answer_id;primary_key" json:"answer_id;primary_key"`
	PollId      uint16 `gorm:"column:poll_id" form:"poll_id" json:"poll_id"`
	AnswerTitle string `gorm:"column:answer_title" form:"answer_title" json:"answer_title"`
	VotesCount  uint16 `gorm:"column:votes_count" form:"votes_count" json:"votes_count"`
	AnswerOrder uint16 `gorm:"column:answer_order" form:"answer_order" json:"answer_order"`
}

// TableName :
func (*PollAnswer) TableName() string {
	return "poll_answer"
}

// handler create
func initRoutersPollAnswer(r *gin.Engine, pollanswer string) {
	route := r.Group("pollanswer")
	route.GET("/", getPollAnswers)
	route.GET("/:id", getPollAnswer)
	route.POST("/", createPollAnswer)
	route.PUT("/:id", updatePollAnswer)
	route.DELETE("/:id", deletePollAnswer)
}

func getPollAnswers(c *gin.Context) {
	var pollAnswers []PollAnswer
	if err := g.Find(&pollAnswers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pollAnswers)
	}
}

func getPollAnswer(c *gin.Context) {
	id := c.Params.ByName("id")
	var pollAnswer PollAnswer
	if err := g.Where("id = ?", id).First(&pollAnswer).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pollAnswer)
	}
}

func createPollAnswer(c *gin.Context) {
	var pollAnswer PollAnswer
	c.BindJSON(&pollAnswer)
	g.Create(&pollAnswer)
	c.JSON(200, pollAnswer)
}

func updatePollAnswer(c *gin.Context) {
	var pollAnswer PollAnswer
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&pollAnswer).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&pollAnswer)
	g.Save(&pollAnswer)
	c.JSON(200, pollAnswer)
}
func deletePollAnswer(c *gin.Context) {
	id := c.Params.ByName("id")
	var pollAnswer PollAnswer
	d := g.Where("id = ?", id).Delete(&pollAnswer)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
