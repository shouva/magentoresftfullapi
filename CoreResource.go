package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreResource :
type CoreResource struct {
	Code        string `gorm:"column:code;primary_key" form:"code;primary_key" json:"code;primary_key"`
	Version     string `gorm:"column:version" form:"version" json:"version"`
	DataVersion string `gorm:"column:data_version" form:"data_version" json:"data_version"`
}

// TableName :
func (*CoreResource) TableName() string {
	return "core_resource"
}

// handler create
func initRoutersCoreResource(r *gin.Engine, coreresource string) {
	route := r.Group("coreresource")
	route.GET("/", getCoreResources)
	route.GET("/:id", getCoreResource)
	route.POST("/", createCoreResource)
	route.PUT("/:id", updateCoreResource)
	route.DELETE("/:id", deleteCoreResource)
}

func getCoreResources(c *gin.Context) {
	var coreResources []CoreResource
	if err := g.Find(&coreResources).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreResources)
	}
}

func getCoreResource(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreResource CoreResource
	if err := g.Where("id = ?", id).First(&coreResource).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreResource)
	}
}

func createCoreResource(c *gin.Context) {
	var coreResource CoreResource
	c.BindJSON(&coreResource)
	g.Create(&coreResource)
	c.JSON(200, coreResource)
}

func updateCoreResource(c *gin.Context) {
	var coreResource CoreResource
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreResource).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreResource)
	g.Save(&coreResource)
	c.JSON(200, coreResource)
}
func deleteCoreResource(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreResource CoreResource
	d := g.Where("id = ?", id).Delete(&coreResource)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
