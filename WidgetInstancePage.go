package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// WidgetInstancePage :
type WidgetInstancePage struct {
	PageId         uint16 `gorm:"column:page_id;primary_key" form:"page_id;primary_key" json:"page_id;primary_key"`
	InstanceId     uint16 `gorm:"column:instance_id" form:"instance_id" json:"instance_id"`
	PageGroup      string `gorm:"column:page_group" form:"page_group" json:"page_group"`
	LayoutHandle   string `gorm:"column:layout_handle" form:"layout_handle" json:"layout_handle"`
	BlockReference string `gorm:"column:block_reference" form:"block_reference" json:"block_reference"`
	PageFor        string `gorm:"column:page_for" form:"page_for" json:"page_for"`
	Entities       string `gorm:"column:entities" form:"entities" json:"entities"`
	PageTemplate   string `gorm:"column:page_template" form:"page_template" json:"page_template"`
}

// TableName :
func (*WidgetInstancePage) TableName() string {
	return "widget_instance_page"
}

// handler create
func initRoutersWidgetInstancePage(r *gin.Engine, widgetinstancepage string) {
	route := r.Group("widgetinstancepage")
	route.GET("/", getWidgetInstancePages)
	route.GET("/:id", getWidgetInstancePage)
	route.POST("/", createWidgetInstancePage)
	route.PUT("/:id", updateWidgetInstancePage)
	route.DELETE("/:id", deleteWidgetInstancePage)
}

func getWidgetInstancePages(c *gin.Context) {
	var widgetInstancePages []WidgetInstancePage
	if err := g.Find(&widgetInstancePages).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, widgetInstancePages)
	}
}

func getWidgetInstancePage(c *gin.Context) {
	id := c.Params.ByName("id")
	var widgetInstancePage WidgetInstancePage
	if err := g.Where("id = ?", id).First(&widgetInstancePage).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, widgetInstancePage)
	}
}

func createWidgetInstancePage(c *gin.Context) {
	var widgetInstancePage WidgetInstancePage
	c.BindJSON(&widgetInstancePage)
	g.Create(&widgetInstancePage)
	c.JSON(200, widgetInstancePage)
}

func updateWidgetInstancePage(c *gin.Context) {
	var widgetInstancePage WidgetInstancePage
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&widgetInstancePage).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&widgetInstancePage)
	g.Save(&widgetInstancePage)
	c.JSON(200, widgetInstancePage)
}
func deleteWidgetInstancePage(c *gin.Context) {
	id := c.Params.ByName("id")
	var widgetInstancePage WidgetInstancePage
	d := g.Where("id = ?", id).Delete(&widgetInstancePage)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
