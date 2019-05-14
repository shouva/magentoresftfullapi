package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// DataflowSession :
type DataflowSession struct {
	SessionId   uint16     `gorm:"column:session_id;primary_key" form:"session_id;primary_key" json:"session_id;primary_key"`
	UserId      uint16     `gorm:"column:user_id" form:"user_id" json:"user_id"`
	CreatedDate *time.Time `gorm:"column:created_date" form:"created_date" json:"created_date"`
	File        string     `gorm:"column:file" form:"file" json:"file"`
	Type        string     `gorm:"column:type" form:"type" json:"type"`
	Direction   string     `gorm:"column:direction" form:"direction" json:"direction"`
	Comment     string     `gorm:"column:comment" form:"comment" json:"comment"`
}

// TableName :
func (*DataflowSession) TableName() string {
	return "dataflow_session"
}

// handler create
func initRoutersDataflowSession(r *gin.Engine, dataflowsession string) {
	route := r.Group("dataflowsession")
	route.GET("/", getDataflowSessions)
	route.GET("/:id", getDataflowSession)
	route.POST("/", createDataflowSession)
	route.PUT("/:id", updateDataflowSession)
	route.DELETE("/:id", deleteDataflowSession)
}

func getDataflowSessions(c *gin.Context) {
	var dataflowSessions []DataflowSession
	if err := g.Find(&dataflowSessions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowSessions)
	}
}

func getDataflowSession(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowSession DataflowSession
	if err := g.Where("id = ?", id).First(&dataflowSession).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowSession)
	}
}

func createDataflowSession(c *gin.Context) {
	var dataflowSession DataflowSession
	c.BindJSON(&dataflowSession)
	g.Create(&dataflowSession)
	c.JSON(200, dataflowSession)
}

func updateDataflowSession(c *gin.Context) {
	var dataflowSession DataflowSession
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&dataflowSession).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&dataflowSession)
	g.Save(&dataflowSession)
	c.JSON(200, dataflowSession)
}
func deleteDataflowSession(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowSession DataflowSession
	d := g.Where("id = ?", id).Delete(&dataflowSession)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
