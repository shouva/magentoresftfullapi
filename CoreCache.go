package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreCache :
type CoreCache struct {
	ID         string `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Data       []byte `gorm:"column:data" form:"data" json:"data"`
	CreateTime uint16 `gorm:"column:create_time" form:"create_time" json:"create_time"`
	UpdateTime uint16 `gorm:"column:update_time" form:"update_time" json:"update_time"`
	ExpireTime uint16 `gorm:"column:expire_time" form:"expire_time" json:"expire_time"`
}

// TableName :
func (*CoreCache) TableName() string {
	return "core_cache"
}

// handler create
func initRoutersCoreCache(r *gin.Engine, corecache string) {
	route := r.Group("corecache")
	route.GET("/", getCoreCaches)
	route.GET("/:id", getCoreCache)
	route.POST("/", createCoreCache)
	route.PUT("/:id", updateCoreCache)
	route.DELETE("/:id", deleteCoreCache)
}

func getCoreCaches(c *gin.Context) {
	var coreCaches []CoreCache
	if err := g.Find(&coreCaches).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreCaches)
	}
}

func getCoreCache(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreCache CoreCache
	if err := g.Where("id = ?", id).First(&coreCache).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreCache)
	}
}

func createCoreCache(c *gin.Context) {
	var coreCache CoreCache
	c.BindJSON(&coreCache)
	g.Create(&coreCache)
	c.JSON(200, coreCache)
}

func updateCoreCache(c *gin.Context) {
	var coreCache CoreCache
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreCache).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreCache)
	g.Save(&coreCache)
	c.JSON(200, coreCache)
}
func deleteCoreCache(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreCache CoreCache
	d := g.Where("id = ?", id).Delete(&coreCache)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
