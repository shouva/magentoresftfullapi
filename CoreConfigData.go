package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreConfigData :
type CoreConfigData struct {
	ConfigId uint16 `gorm:"column:config_id;primary_key" form:"config_id;primary_key" json:"config_id;primary_key"`
	Scope    string `gorm:"column:scope" form:"scope" json:"scope"`
	ScopeId  uint16 `gorm:"column:scope_id" form:"scope_id" json:"scope_id"`
	Path     string `gorm:"column:path" form:"path" json:"path"`
	Value    string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CoreConfigData) TableName() string {
	return "core_config_data"
}

// handler create
func initRoutersCoreConfigData(r *gin.Engine, coreconfigdata string) {
	route := r.Group("coreconfigdata")
	route.GET("/", getCoreConfigDatas)
	route.GET("/:id", getCoreConfigData)
	route.POST("/", createCoreConfigData)
	route.PUT("/:id", updateCoreConfigData)
	route.DELETE("/:id", deleteCoreConfigData)
}

func getCoreConfigDatas(c *gin.Context) {
	var coreConfigDatas []CoreConfigData
	if err := g.Find(&coreConfigDatas).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreConfigDatas)
	}
}

func getCoreConfigData(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreConfigData CoreConfigData
	if err := g.Where("id = ?", id).First(&coreConfigData).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreConfigData)
	}
}

func createCoreConfigData(c *gin.Context) {
	var coreConfigData CoreConfigData
	c.BindJSON(&coreConfigData)
	g.Create(&coreConfigData)
	c.JSON(200, coreConfigData)
}

func updateCoreConfigData(c *gin.Context) {
	var coreConfigData CoreConfigData
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreConfigData).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreConfigData)
	g.Save(&coreConfigData)
	c.JSON(200, coreConfigData)
}
func deleteCoreConfigData(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreConfigData CoreConfigData
	d := g.Where("id = ?", id).Delete(&coreConfigData)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
