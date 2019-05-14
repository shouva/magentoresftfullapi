package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// NewsletterProblem :
type NewsletterProblem struct {
	ProblemId        uint16 `gorm:"column:problem_id;primary_key" form:"problem_id;primary_key" json:"problem_id;primary_key"`
	SubscriberId     uint16 `gorm:"column:subscriber_id" form:"subscriber_id" json:"subscriber_id"`
	QueueId          uint16 `gorm:"column:queue_id" form:"queue_id" json:"queue_id"`
	ProblemErrorCode uint16 `gorm:"column:problem_error_code" form:"problem_error_code" json:"problem_error_code"`
	ProblemErrorText string `gorm:"column:problem_error_text" form:"problem_error_text" json:"problem_error_text"`
}

// TableName :
func (*NewsletterProblem) TableName() string {
	return "newsletter_problem"
}

// handler create
func initRoutersNewsletterProblem(r *gin.Engine, newsletterproblem string) {
	route := r.Group("newsletterproblem")
	route.GET("/", getNewsletterProblems)
	route.GET("/:id", getNewsletterProblem)
	route.POST("/", createNewsletterProblem)
	route.PUT("/:id", updateNewsletterProblem)
	route.DELETE("/:id", deleteNewsletterProblem)
}

func getNewsletterProblems(c *gin.Context) {
	var newsletterProblems []NewsletterProblem
	if err := g.Find(&newsletterProblems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterProblems)
	}
}

func getNewsletterProblem(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterProblem NewsletterProblem
	if err := g.Where("id = ?", id).First(&newsletterProblem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterProblem)
	}
}

func createNewsletterProblem(c *gin.Context) {
	var newsletterProblem NewsletterProblem
	c.BindJSON(&newsletterProblem)
	g.Create(&newsletterProblem)
	c.JSON(200, newsletterProblem)
}

func updateNewsletterProblem(c *gin.Context) {
	var newsletterProblem NewsletterProblem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&newsletterProblem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&newsletterProblem)
	g.Save(&newsletterProblem)
	c.JSON(200, newsletterProblem)
}
func deleteNewsletterProblem(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterProblem NewsletterProblem
	d := g.Where("id = ?", id).Delete(&newsletterProblem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
