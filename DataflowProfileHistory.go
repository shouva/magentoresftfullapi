package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// DataflowProfileHistory :
type DataflowProfileHistory struct {
	HistoryId   uint16     `gorm:"column:history_id;primary_key" form:"history_id;primary_key" json:"history_id;primary_key"`
	ProfileId   uint16     `gorm:"column:profile_id" form:"profile_id" json:"profile_id"`
	ActionCode  string     `gorm:"column:action_code" form:"action_code" json:"action_code"`
	UserId      uint16     `gorm:"column:user_id" form:"user_id" json:"user_id"`
	PerformedAt *time.Time `gorm:"column:performed_at" form:"performed_at" json:"performed_at"`
}

// TableName :
func (*DataflowProfileHistory) TableName() string {
	return "dataflow_profile_history"
}

// handler create
func initRoutersDataflowProfileHistory(r *gin.Engine, dataflowprofilehistory string) {
	route := r.Group("dataflowprofilehistory")
	route.GET("/", getDataflowProfileHistorys)
	route.GET("/:id", getDataflowProfileHistory)
	route.POST("/", createDataflowProfileHistory)
	route.PUT("/:id", updateDataflowProfileHistory)
	route.DELETE("/:id", deleteDataflowProfileHistory)
}

func getDataflowProfileHistorys(c *gin.Context) {
	var dataflowProfileHistorys []DataflowProfileHistory
	if err := g.Find(&dataflowProfileHistorys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowProfileHistorys)
	}
}

func getDataflowProfileHistory(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowProfileHistory DataflowProfileHistory
	if err := g.Where("id = ?", id).First(&dataflowProfileHistory).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowProfileHistory)
	}
}

func createDataflowProfileHistory(c *gin.Context) {
	var dataflowProfileHistory DataflowProfileHistory
	c.BindJSON(&dataflowProfileHistory)
	g.Create(&dataflowProfileHistory)
	c.JSON(200, dataflowProfileHistory)
}

func updateDataflowProfileHistory(c *gin.Context) {
	var dataflowProfileHistory DataflowProfileHistory
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&dataflowProfileHistory).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&dataflowProfileHistory)
	g.Save(&dataflowProfileHistory)
	c.JSON(200, dataflowProfileHistory)
}
func deleteDataflowProfileHistory(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowProfileHistory DataflowProfileHistory
	d := g.Where("id = ?", id).Delete(&dataflowProfileHistory)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
