package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LogQuote :
type LogQuote struct {
	QuoteId   uint16     `gorm:"column:quote_id;primary_key" form:"quote_id;primary_key" json:"quote_id;primary_key"`
	VisitorId uint64     `gorm:"column:visitor_id" form:"visitor_id" json:"visitor_id"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at"`
}

// TableName :
func (*LogQuote) TableName() string {
	return "log_quote"
}

// handler create
func initRoutersLogQuote(r *gin.Engine, logquote string) {
	route := r.Group("logquote")
	route.GET("/", getLogQuotes)
	route.GET("/:id", getLogQuote)
	route.POST("/", createLogQuote)
	route.PUT("/:id", updateLogQuote)
	route.DELETE("/:id", deleteLogQuote)
}

func getLogQuotes(c *gin.Context) {
	var logQuotes []LogQuote
	if err := g.Find(&logQuotes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logQuotes)
	}
}

func getLogQuote(c *gin.Context) {
	id := c.Params.ByName("id")
	var logQuote LogQuote
	if err := g.Where("id = ?", id).First(&logQuote).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logQuote)
	}
}

func createLogQuote(c *gin.Context) {
	var logQuote LogQuote
	c.BindJSON(&logQuote)
	g.Create(&logQuote)
	c.JSON(200, logQuote)
}

func updateLogQuote(c *gin.Context) {
	var logQuote LogQuote
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&logQuote).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&logQuote)
	g.Save(&logQuote)
	c.JSON(200, logQuote)
}
func deleteLogQuote(c *gin.Context) {
	id := c.Params.ByName("id")
	var logQuote LogQuote
	d := g.Where("id = ?", id).Delete(&logQuote)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
