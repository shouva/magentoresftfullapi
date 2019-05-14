package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ReportEvent :
type ReportEvent struct {
	EventId     uint64     `gorm:"column:event_id;primary_key" form:"event_id;primary_key" json:"event_id;primary_key"`
	LoggedAt    *time.Time `gorm:"column:logged_at" form:"logged_at" json:"logged_at"`
	EventTypeId uint16     `gorm:"column:event_type_id" form:"event_type_id" json:"event_type_id"`
	ObjectId    uint16     `gorm:"column:object_id" form:"object_id" json:"object_id"`
	SubjectId   uint16     `gorm:"column:subject_id" form:"subject_id" json:"subject_id"`
	Subtype     uint16     `gorm:"column:subtype" form:"subtype" json:"subtype"`
	StoreId     uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
}

// TableName :
func (*ReportEvent) TableName() string {
	return "report_event"
}

// handler create
func initRoutersReportEvent(r *gin.Engine, reportevent string) {
	route := r.Group("reportevent")
	route.GET("/", getReportEvents)
	route.GET("/:id", getReportEvent)
	route.POST("/", createReportEvent)
	route.PUT("/:id", updateReportEvent)
	route.DELETE("/:id", deleteReportEvent)
}

func getReportEvents(c *gin.Context) {
	var reportEvents []ReportEvent
	if err := g.Find(&reportEvents).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportEvents)
	}
}

func getReportEvent(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportEvent ReportEvent
	if err := g.Where("id = ?", id).First(&reportEvent).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportEvent)
	}
}

func createReportEvent(c *gin.Context) {
	var reportEvent ReportEvent
	c.BindJSON(&reportEvent)
	g.Create(&reportEvent)
	c.JSON(200, reportEvent)
}

func updateReportEvent(c *gin.Context) {
	var reportEvent ReportEvent
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reportEvent).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reportEvent)
	g.Save(&reportEvent)
	c.JSON(200, reportEvent)
}
func deleteReportEvent(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportEvent ReportEvent
	d := g.Where("id = ?", id).Delete(&reportEvent)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
