package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreLayoutUpdate :
type CoreLayoutUpdate struct {
	LayoutUpdateId uint16 `gorm:"column:layout_update_id;primary_key" form:"layout_update_id;primary_key" json:"layout_update_id;primary_key"`
	Handle         string `gorm:"column:handle" form:"handle" json:"handle"`
	Xml            string `gorm:"column:xml" form:"xml" json:"xml"`
	SortOrder      uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
}

// TableName :
func (*CoreLayoutUpdate) TableName() string {
	return "core_layout_update"
}

// handler create
func initRoutersCoreLayoutUpdate(r *gin.Engine, corelayoutupdate string) {
	route := r.Group("corelayoutupdate")
	route.GET("/", getCoreLayoutUpdates)
	route.GET("/:id", getCoreLayoutUpdate)
	route.POST("/", createCoreLayoutUpdate)
	route.PUT("/:id", updateCoreLayoutUpdate)
	route.DELETE("/:id", deleteCoreLayoutUpdate)
}

func getCoreLayoutUpdates(c *gin.Context) {
	var coreLayoutUpdates []CoreLayoutUpdate
	if err := g.Find(&coreLayoutUpdates).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreLayoutUpdates)
	}
}

func getCoreLayoutUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreLayoutUpdate CoreLayoutUpdate
	if err := g.Where("id = ?", id).First(&coreLayoutUpdate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreLayoutUpdate)
	}
}

func createCoreLayoutUpdate(c *gin.Context) {
	var coreLayoutUpdate CoreLayoutUpdate
	c.BindJSON(&coreLayoutUpdate)
	g.Create(&coreLayoutUpdate)
	c.JSON(200, coreLayoutUpdate)
}

func updateCoreLayoutUpdate(c *gin.Context) {
	var coreLayoutUpdate CoreLayoutUpdate
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreLayoutUpdate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreLayoutUpdate)
	g.Save(&coreLayoutUpdate)
	c.JSON(200, coreLayoutUpdate)
}
func deleteCoreLayoutUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreLayoutUpdate CoreLayoutUpdate
	d := g.Where("id = ?", id).Delete(&coreLayoutUpdate)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
