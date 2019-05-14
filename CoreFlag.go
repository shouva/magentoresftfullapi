package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CoreFlag :
type CoreFlag struct {
	FlagId     uint16     `gorm:"column:flag_id;primary_key" form:"flag_id;primary_key" json:"flag_id;primary_key"`
	FlagCode   string     `gorm:"column:flag_code" form:"flag_code" json:"flag_code"`
	State      uint16     `gorm:"column:state" form:"state" json:"state"`
	FlagData   string     `gorm:"column:flag_data" form:"flag_data" json:"flag_data"`
	LastUpdate *time.Time `gorm:"column:last_update" form:"last_update" json:"last_update"`
}

// TableName :
func (*CoreFlag) TableName() string {
	return "core_flag"
}

// handler create
func initRoutersCoreFlag(r *gin.Engine, coreflag string) {
	route := r.Group("coreflag")
	route.GET("/", getCoreFlags)
	route.GET("/:id", getCoreFlag)
	route.POST("/", createCoreFlag)
	route.PUT("/:id", updateCoreFlag)
	route.DELETE("/:id", deleteCoreFlag)
}

func getCoreFlags(c *gin.Context) {
	var coreFlags []CoreFlag
	if err := g.Find(&coreFlags).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreFlags)
	}
}

func getCoreFlag(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreFlag CoreFlag
	if err := g.Where("id = ?", id).First(&coreFlag).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreFlag)
	}
}

func createCoreFlag(c *gin.Context) {
	var coreFlag CoreFlag
	c.BindJSON(&coreFlag)
	g.Create(&coreFlag)
	c.JSON(200, coreFlag)
}

func updateCoreFlag(c *gin.Context) {
	var coreFlag CoreFlag
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreFlag).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreFlag)
	g.Save(&coreFlag)
	c.JSON(200, coreFlag)
}
func deleteCoreFlag(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreFlag CoreFlag
	d := g.Where("id = ?", id).Delete(&coreFlag)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
