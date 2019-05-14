package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LogSummary :
type LogSummary struct {
	SummaryId     uint64     `gorm:"column:summary_id;primary_key" form:"summary_id;primary_key" json:"summary_id;primary_key"`
	StoreId       uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	TypeId        uint16     `gorm:"column:type_id" form:"type_id" json:"type_id"`
	VisitorCount  uint16     `gorm:"column:visitor_count" form:"visitor_count" json:"visitor_count"`
	CustomerCount uint16     `gorm:"column:customer_count" form:"customer_count" json:"customer_count"`
	AddDate       *time.Time `gorm:"column:add_date" form:"add_date" json:"add_date"`
}

// TableName :
func (*LogSummary) TableName() string {
	return "log_summary"
}

// handler create
func initRoutersLogSummary(r *gin.Engine, logsummary string) {
	route := r.Group("logsummary")
	route.GET("/", getLogSummarys)
	route.GET("/:id", getLogSummary)
	route.POST("/", createLogSummary)
	route.PUT("/:id", updateLogSummary)
	route.DELETE("/:id", deleteLogSummary)
}

func getLogSummarys(c *gin.Context) {
	var logSummarys []LogSummary
	if err := g.Find(&logSummarys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logSummarys)
	}
}

func getLogSummary(c *gin.Context) {
	id := c.Params.ByName("id")
	var logSummary LogSummary
	if err := g.Where("id = ?", id).First(&logSummary).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logSummary)
	}
}

func createLogSummary(c *gin.Context) {
	var logSummary LogSummary
	c.BindJSON(&logSummary)
	g.Create(&logSummary)
	c.JSON(200, logSummary)
}

func updateLogSummary(c *gin.Context) {
	var logSummary LogSummary
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&logSummary).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&logSummary)
	g.Save(&logSummary)
	c.JSON(200, logSummary)
}
func deleteLogSummary(c *gin.Context) {
	id := c.Params.ByName("id")
	var logSummary LogSummary
	d := g.Where("id = ?", id).Delete(&logSummary)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
