package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Widget :
type Widget struct {
	WidgetId   uint16 `gorm:"column:widget_id;primary_key" form:"widget_id;primary_key" json:"widget_id;primary_key"`
	WidgetCode string `gorm:"column:widget_code" form:"widget_code" json:"widget_code"`
	WidgetType string `gorm:"column:widget_type" form:"widget_type" json:"widget_type"`
	Parameters string `gorm:"column:parameters" form:"parameters" json:"parameters"`
}

// TableName :
func (*Widget) TableName() string {
	return "widget"
}

// handler create
func initRoutersWidget(r *gin.Engine, widget string) {
	route := r.Group("widget")
	route.GET("/", getWidgets)
	route.GET("/:id", getWidget)
	route.POST("/", createWidget)
	route.PUT("/:id", updateWidget)
	route.DELETE("/:id", deleteWidget)
}

func getWidgets(c *gin.Context) {
	var widgets []Widget
	if err := g.Find(&widgets).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, widgets)
	}
}

func getWidget(c *gin.Context) {
	id := c.Params.ByName("id")
	var widget Widget
	if err := g.Where("id = ?", id).First(&widget).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, widget)
	}
}

func createWidget(c *gin.Context) {
	var widget Widget
	c.BindJSON(&widget)
	g.Create(&widget)
	c.JSON(200, widget)
}

func updateWidget(c *gin.Context) {
	var widget Widget
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&widget).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&widget)
	g.Save(&widget)
	c.JSON(200, widget)
}
func deleteWidget(c *gin.Context) {
	id := c.Params.ByName("id")
	var widget Widget
	d := g.Where("id = ?", id).Delete(&widget)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
