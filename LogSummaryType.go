package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LogSummaryType :
type LogSummaryType struct {
	TypeId     uint16 `gorm:"column:type_id;primary_key" form:"type_id;primary_key" json:"type_id;primary_key"`
	TypeCode   string `gorm:"column:type_code" form:"type_code" json:"type_code"`
	Period     uint16 `gorm:"column:period" form:"period" json:"period"`
	PeriodType string `gorm:"column:period_type" form:"period_type" json:"period_type"`
}

// TableName :
func (*LogSummaryType) TableName() string {
	return "log_summary_type"
}

// handler create
func initRoutersLogSummaryType(r *gin.Engine, logsummarytype string) {
	route := r.Group("logsummarytype")
	route.GET("/", getLogSummaryTypes)
	route.GET("/:id", getLogSummaryType)
	route.POST("/", createLogSummaryType)
	route.PUT("/:id", updateLogSummaryType)
	route.DELETE("/:id", deleteLogSummaryType)
}

func getLogSummaryTypes(c *gin.Context) {
	var logSummaryTypes []LogSummaryType
	if err := g.Find(&logSummaryTypes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logSummaryTypes)
	}
}

func getLogSummaryType(c *gin.Context) {
	id := c.Params.ByName("id")
	var logSummaryType LogSummaryType
	if err := g.Where("id = ?", id).First(&logSummaryType).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logSummaryType)
	}
}

func createLogSummaryType(c *gin.Context) {
	var logSummaryType LogSummaryType
	c.BindJSON(&logSummaryType)
	g.Create(&logSummaryType)
	c.JSON(200, logSummaryType)
}

func updateLogSummaryType(c *gin.Context) {
	var logSummaryType LogSummaryType
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&logSummaryType).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&logSummaryType)
	g.Save(&logSummaryType)
	c.JSON(200, logSummaryType)
}
func deleteLogSummaryType(c *gin.Context) {
	id := c.Params.ByName("id")
	var logSummaryType LogSummaryType
	d := g.Where("id = ?", id).Delete(&logSummaryType)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
