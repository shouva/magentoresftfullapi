package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreCacheOption :
type CoreCacheOption struct {
	Code  string `gorm:"column:code;primary_key" form:"code;primary_key" json:"code;primary_key"`
	Value uint16 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CoreCacheOption) TableName() string {
	return "core_cache_option"
}

// handler create
func initRoutersCoreCacheOption(r *gin.Engine, corecacheoption string) {
	route := r.Group("corecacheoption")
	route.GET("/", getCoreCacheOptions)
	route.GET("/:id", getCoreCacheOption)
	route.POST("/", createCoreCacheOption)
	route.PUT("/:id", updateCoreCacheOption)
	route.DELETE("/:id", deleteCoreCacheOption)
}

func getCoreCacheOptions(c *gin.Context) {
	var coreCacheOptions []CoreCacheOption
	if err := g.Find(&coreCacheOptions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreCacheOptions)
	}
}

func getCoreCacheOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreCacheOption CoreCacheOption
	if err := g.Where("id = ?", id).First(&coreCacheOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreCacheOption)
	}
}

func createCoreCacheOption(c *gin.Context) {
	var coreCacheOption CoreCacheOption
	c.BindJSON(&coreCacheOption)
	g.Create(&coreCacheOption)
	c.JSON(200, coreCacheOption)
}

func updateCoreCacheOption(c *gin.Context) {
	var coreCacheOption CoreCacheOption
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreCacheOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreCacheOption)
	g.Save(&coreCacheOption)
	c.JSON(200, coreCacheOption)
}
func deleteCoreCacheOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreCacheOption CoreCacheOption
	d := g.Where("id = ?", id).Delete(&coreCacheOption)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
