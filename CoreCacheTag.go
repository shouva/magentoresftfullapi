package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreCacheTag :
type CoreCacheTag struct {
	Tag     string `gorm:"column:tag;primary_key" form:"tag;primary_key" json:"tag;primary_key"`
	CacheId string `gorm:"column:cache_id;primary_key" form:"cache_id;primary_key" json:"cache_id;primary_key"`
}

// TableName :
func (*CoreCacheTag) TableName() string {
	return "core_cache_tag"
}

// handler create
func initRoutersCoreCacheTag(r *gin.Engine, corecachetag string) {
	route := r.Group("corecachetag")
	route.GET("/", getCoreCacheTags)
	route.GET("/:id", getCoreCacheTag)
	route.POST("/", createCoreCacheTag)
	route.PUT("/:id", updateCoreCacheTag)
	route.DELETE("/:id", deleteCoreCacheTag)
}

func getCoreCacheTags(c *gin.Context) {
	var coreCacheTags []CoreCacheTag
	if err := g.Find(&coreCacheTags).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreCacheTags)
	}
}

func getCoreCacheTag(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreCacheTag CoreCacheTag
	if err := g.Where("id = ?", id).First(&coreCacheTag).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreCacheTag)
	}
}

func createCoreCacheTag(c *gin.Context) {
	var coreCacheTag CoreCacheTag
	c.BindJSON(&coreCacheTag)
	g.Create(&coreCacheTag)
	c.JSON(200, coreCacheTag)
}

func updateCoreCacheTag(c *gin.Context) {
	var coreCacheTag CoreCacheTag
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreCacheTag).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreCacheTag)
	g.Save(&coreCacheTag)
	c.JSON(200, coreCacheTag)
}
func deleteCoreCacheTag(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreCacheTag CoreCacheTag
	d := g.Where("id = ?", id).Delete(&coreCacheTag)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
