package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreLayoutLink :
type CoreLayoutLink struct {
	LayoutLinkId   uint16 `gorm:"column:layout_link_id;primary_key" form:"layout_link_id;primary_key" json:"layout_link_id;primary_key"`
	StoreId        uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Area           string `gorm:"column:area" form:"area" json:"area"`
	Package        string `gorm:"column:package" form:"package" json:"package"`
	Theme          string `gorm:"column:theme" form:"theme" json:"theme"`
	LayoutUpdateId uint16 `gorm:"column:layout_update_id" form:"layout_update_id" json:"layout_update_id"`
}

// TableName :
func (*CoreLayoutLink) TableName() string {
	return "core_layout_link"
}

// handler create
func initRoutersCoreLayoutLink(r *gin.Engine, corelayoutlink string) {
	route := r.Group("corelayoutlink")
	route.GET("/", getCoreLayoutLinks)
	route.GET("/:id", getCoreLayoutLink)
	route.POST("/", createCoreLayoutLink)
	route.PUT("/:id", updateCoreLayoutLink)
	route.DELETE("/:id", deleteCoreLayoutLink)
}

func getCoreLayoutLinks(c *gin.Context) {
	var coreLayoutLinks []CoreLayoutLink
	if err := g.Find(&coreLayoutLinks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreLayoutLinks)
	}
}

func getCoreLayoutLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreLayoutLink CoreLayoutLink
	if err := g.Where("id = ?", id).First(&coreLayoutLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreLayoutLink)
	}
}

func createCoreLayoutLink(c *gin.Context) {
	var coreLayoutLink CoreLayoutLink
	c.BindJSON(&coreLayoutLink)
	g.Create(&coreLayoutLink)
	c.JSON(200, coreLayoutLink)
}

func updateCoreLayoutLink(c *gin.Context) {
	var coreLayoutLink CoreLayoutLink
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreLayoutLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreLayoutLink)
	g.Save(&coreLayoutLink)
	c.JSON(200, coreLayoutLink)
}
func deleteCoreLayoutLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreLayoutLink CoreLayoutLink
	d := g.Where("id = ?", id).Delete(&coreLayoutLink)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
