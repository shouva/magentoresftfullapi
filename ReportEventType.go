package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ReportEventType :
type ReportEventType struct {
	EventTypeId   uint16 `gorm:"column:event_type_id;primary_key" form:"event_type_id;primary_key" json:"event_type_id;primary_key"`
	EventName     string `gorm:"column:event_name" form:"event_name" json:"event_name"`
	CustomerLogin uint16 `gorm:"column:customer_login" form:"customer_login" json:"customer_login"`
}

// TableName :
func (*ReportEventType) TableName() string {
	return "report_event_types"
}

// handler create
func initRoutersReportEventType(r *gin.Engine, reporteventtype string) {
	route := r.Group("reporteventtype")
	route.GET("/", getReportEventTypes)
	route.GET("/:id", getReportEventType)
	route.POST("/", createReportEventType)
	route.PUT("/:id", updateReportEventType)
	route.DELETE("/:id", deleteReportEventType)
}

func getReportEventTypes(c *gin.Context) {
	var reportEventTypes []ReportEventType
	if err := g.Find(&reportEventTypes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportEventTypes)
	}
}

func getReportEventType(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportEventType ReportEventType
	if err := g.Where("id = ?", id).First(&reportEventType).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportEventType)
	}
}

func createReportEventType(c *gin.Context) {
	var reportEventType ReportEventType
	c.BindJSON(&reportEventType)
	g.Create(&reportEventType)
	c.JSON(200, reportEventType)
}

func updateReportEventType(c *gin.Context) {
	var reportEventType ReportEventType
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reportEventType).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reportEventType)
	g.Save(&reportEventType)
	c.JSON(200, reportEventType)
}
func deleteReportEventType(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportEventType ReportEventType
	d := g.Where("id = ?", id).Delete(&reportEventType)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
