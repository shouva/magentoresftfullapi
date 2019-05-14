package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// WidgetInstancePageLayout :
type WidgetInstancePageLayout struct {
	PageId         uint16 `gorm:"column:page_id;primary_key" form:"page_id;primary_key" json:"page_id;primary_key"`
	LayoutUpdateId uint16 `gorm:"column:layout_update_id;primary_key" form:"layout_update_id;primary_key" json:"layout_update_id;primary_key"`
}

// TableName :
func (*WidgetInstancePageLayout) TableName() string {
	return "widget_instance_page_layout"
}

// handler create
func initRoutersWidgetInstancePageLayout(r *gin.Engine, widgetinstancepagelayout string) {
	route := r.Group("widgetinstancepagelayout")
	route.GET("/", getWidgetInstancePageLayouts)
	route.GET("/:id", getWidgetInstancePageLayout)
	route.POST("/", createWidgetInstancePageLayout)
	route.PUT("/:id", updateWidgetInstancePageLayout)
	route.DELETE("/:id", deleteWidgetInstancePageLayout)
}

func getWidgetInstancePageLayouts(c *gin.Context) {
	var widgetInstancePageLayouts []WidgetInstancePageLayout
	if err := g.Find(&widgetInstancePageLayouts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, widgetInstancePageLayouts)
	}
}

func getWidgetInstancePageLayout(c *gin.Context) {
	id := c.Params.ByName("id")
	var widgetInstancePageLayout WidgetInstancePageLayout
	if err := g.Where("id = ?", id).First(&widgetInstancePageLayout).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, widgetInstancePageLayout)
	}
}

func createWidgetInstancePageLayout(c *gin.Context) {
	var widgetInstancePageLayout WidgetInstancePageLayout
	c.BindJSON(&widgetInstancePageLayout)
	g.Create(&widgetInstancePageLayout)
	c.JSON(200, widgetInstancePageLayout)
}

func updateWidgetInstancePageLayout(c *gin.Context) {
	var widgetInstancePageLayout WidgetInstancePageLayout
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&widgetInstancePageLayout).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&widgetInstancePageLayout)
	g.Save(&widgetInstancePageLayout)
	c.JSON(200, widgetInstancePageLayout)
}
func deleteWidgetInstancePageLayout(c *gin.Context) {
	id := c.Params.ByName("id")
	var widgetInstancePageLayout WidgetInstancePageLayout
	d := g.Where("id = ?", id).Delete(&widgetInstancePageLayout)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
