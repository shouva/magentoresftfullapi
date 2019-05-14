package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// WidgetInstance :
type WidgetInstance struct {
	InstanceId       uint16 `gorm:"column:instance_id;primary_key" form:"instance_id;primary_key" json:"instance_id;primary_key"`
	InstanceType     string `gorm:"column:instance_type" form:"instance_type" json:"instance_type"`
	PackageTheme     string `gorm:"column:package_theme" form:"package_theme" json:"package_theme"`
	Title            string `gorm:"column:title" form:"title" json:"title"`
	StoreIds         string `gorm:"column:store_ids" form:"store_ids" json:"store_ids"`
	WidgetParameters string `gorm:"column:widget_parameters" form:"widget_parameters" json:"widget_parameters"`
	SortOrder        uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
}

// TableName :
func (*WidgetInstance) TableName() string {
	return "widget_instance"
}

// handler create
func initRoutersWidgetInstance(r *gin.Engine, widgetinstance string) {
	route := r.Group("widgetinstance")
	route.GET("/", getWidgetInstances)
	route.GET("/:id", getWidgetInstance)
	route.POST("/", createWidgetInstance)
	route.PUT("/:id", updateWidgetInstance)
	route.DELETE("/:id", deleteWidgetInstance)
}

func getWidgetInstances(c *gin.Context) {
	var widgetInstances []WidgetInstance
	if err := g.Find(&widgetInstances).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, widgetInstances)
	}
}

func getWidgetInstance(c *gin.Context) {
	id := c.Params.ByName("id")
	var widgetInstance WidgetInstance
	if err := g.Where("id = ?", id).First(&widgetInstance).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, widgetInstance)
	}
}

func createWidgetInstance(c *gin.Context) {
	var widgetInstance WidgetInstance
	c.BindJSON(&widgetInstance)
	g.Create(&widgetInstance)
	c.JSON(200, widgetInstance)
}

func updateWidgetInstance(c *gin.Context) {
	var widgetInstance WidgetInstance
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&widgetInstance).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&widgetInstance)
	g.Save(&widgetInstance)
	c.JSON(200, widgetInstance)
}
func deleteWidgetInstance(c *gin.Context) {
	id := c.Params.ByName("id")
	var widgetInstance WidgetInstance
	d := g.Where("id = ?", id).Delete(&widgetInstance)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
