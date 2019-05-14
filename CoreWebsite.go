package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreWebsite :
type CoreWebsite struct {
	WebsiteId      uint16 `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	Code           string `gorm:"column:code" form:"code" json:"code"`
	Name           string `gorm:"column:name" form:"name" json:"name"`
	SortOrder      uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	DefaultGroupId uint16 `gorm:"column:default_group_id" form:"default_group_id" json:"default_group_id"`
	IsDefault      uint16 `gorm:"column:is_default" form:"is_default" json:"is_default"`
}

// TableName :
func (*CoreWebsite) TableName() string {
	return "core_website"
}

// handler create
func initRoutersCoreWebsite(r *gin.Engine, corewebsite string) {
	route := r.Group("corewebsite")
	route.GET("/", getCoreWebsites)
	route.GET("/:id", getCoreWebsite)
	route.POST("/", createCoreWebsite)
	route.PUT("/:id", updateCoreWebsite)
	route.DELETE("/:id", deleteCoreWebsite)
}

func getCoreWebsites(c *gin.Context) {
	var coreWebsites []CoreWebsite
	if err := g.Find(&coreWebsites).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreWebsites)
	}
}

func getCoreWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreWebsite CoreWebsite
	if err := g.Where("id = ?", id).First(&coreWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreWebsite)
	}
}

func createCoreWebsite(c *gin.Context) {
	var coreWebsite CoreWebsite
	c.BindJSON(&coreWebsite)
	g.Create(&coreWebsite)
	c.JSON(200, coreWebsite)
}

func updateCoreWebsite(c *gin.Context) {
	var coreWebsite CoreWebsite
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreWebsite)
	g.Save(&coreWebsite)
	c.JSON(200, coreWebsite)
}
func deleteCoreWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreWebsite CoreWebsite
	d := g.Where("id = ?", id).Delete(&coreWebsite)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
