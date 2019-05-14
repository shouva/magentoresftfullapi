package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// XmlconnectHistory :
type XmlconnectHistory struct {
	HistoryId     uint16     `gorm:"column:history_id;primary_key" form:"history_id;primary_key" json:"history_id;primary_key"`
	ApplicationId uint16     `gorm:"column:application_id" form:"application_id" json:"application_id"`
	CreatedAt     *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	StoreId       uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Params        []byte     `gorm:"column:params" form:"params" json:"params"`
	Title         string     `gorm:"column:title" form:"title" json:"title"`
	ActivationKey string     `gorm:"column:activation_key" form:"activation_key" json:"activation_key"`
	Name          string     `gorm:"column:name" form:"name" json:"name"`
	Code          string     `gorm:"column:code" form:"code" json:"code"`
}

// TableName :
func (*XmlconnectHistory) TableName() string {
	return "xmlconnect_history"
}

// handler create
func initRoutersXmlconnectHistory(r *gin.Engine, xmlconnecthistory string) {
	route := r.Group("xmlconnecthistory")
	route.GET("/", getXmlconnectHistorys)
	route.GET("/:id", getXmlconnectHistory)
	route.POST("/", createXmlconnectHistory)
	route.PUT("/:id", updateXmlconnectHistory)
	route.DELETE("/:id", deleteXmlconnectHistory)
}

func getXmlconnectHistorys(c *gin.Context) {
	var xmlconnectHistorys []XmlconnectHistory
	if err := g.Find(&xmlconnectHistorys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectHistorys)
	}
}

func getXmlconnectHistory(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectHistory XmlconnectHistory
	if err := g.Where("id = ?", id).First(&xmlconnectHistory).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectHistory)
	}
}

func createXmlconnectHistory(c *gin.Context) {
	var xmlconnectHistory XmlconnectHistory
	c.BindJSON(&xmlconnectHistory)
	g.Create(&xmlconnectHistory)
	c.JSON(200, xmlconnectHistory)
}

func updateXmlconnectHistory(c *gin.Context) {
	var xmlconnectHistory XmlconnectHistory
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&xmlconnectHistory).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&xmlconnectHistory)
	g.Save(&xmlconnectHistory)
	c.JSON(200, xmlconnectHistory)
}
func deleteXmlconnectHistory(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectHistory XmlconnectHistory
	d := g.Where("id = ?", id).Delete(&xmlconnectHistory)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
