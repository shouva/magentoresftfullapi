package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CmsBlock :
type CmsBlock struct {
	BlockId      uint16     `gorm:"column:block_id;primary_key" form:"block_id;primary_key" json:"block_id;primary_key"`
	Title        string     `gorm:"column:title" form:"title" json:"title"`
	Identifier   string     `gorm:"column:identifier" form:"identifier" json:"identifier"`
	Content      string     `gorm:"column:content" form:"content" json:"content"`
	CreationTime *time.Time `gorm:"column:creation_time" form:"creation_time" json:"creation_time"`
	UpdateTime   *time.Time `gorm:"column:update_time" form:"update_time" json:"update_time"`
	IsActive     uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
}

// TableName :
func (*CmsBlock) TableName() string {
	return "cms_block"
}

// handler create
func initRoutersCmsBlock(r *gin.Engine, cmsblock string) {
	route := r.Group("cmsblock")
	route.GET("/", getCmsBlocks)
	route.GET("/:id", getCmsBlock)
	route.POST("/", createCmsBlock)
	route.PUT("/:id", updateCmsBlock)
	route.DELETE("/:id", deleteCmsBlock)
}

func getCmsBlocks(c *gin.Context) {
	var cmsBlocks []CmsBlock
	if err := g.Find(&cmsBlocks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cmsBlocks)
	}
}

func getCmsBlock(c *gin.Context) {
	id := c.Params.ByName("id")
	var cmsBlock CmsBlock
	if err := g.Where("id = ?", id).First(&cmsBlock).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cmsBlock)
	}
}

func createCmsBlock(c *gin.Context) {
	var cmsBlock CmsBlock
	c.BindJSON(&cmsBlock)
	g.Create(&cmsBlock)
	c.JSON(200, cmsBlock)
}

func updateCmsBlock(c *gin.Context) {
	var cmsBlock CmsBlock
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&cmsBlock).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cmsBlock)
	g.Save(&cmsBlock)
	c.JSON(200, cmsBlock)
}
func deleteCmsBlock(c *gin.Context) {
	id := c.Params.ByName("id")
	var cmsBlock CmsBlock
	d := g.Where("id = ?", id).Delete(&cmsBlock)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
