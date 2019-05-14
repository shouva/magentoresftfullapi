package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CmsBlockStore :
type CmsBlockStore struct {
	BlockId uint16 `gorm:"column:block_id;primary_key" form:"block_id;primary_key" json:"block_id;primary_key"`
	StoreId uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
}

// TableName :
func (*CmsBlockStore) TableName() string {
	return "cms_block_store"
}

// handler create
func initRoutersCmsBlockStore(r *gin.Engine, cmsblockstore string) {
	route := r.Group("cmsblockstore")
	route.GET("/", getCmsBlockStores)
	route.GET("/:id", getCmsBlockStore)
	route.POST("/", createCmsBlockStore)
	route.PUT("/:id", updateCmsBlockStore)
	route.DELETE("/:id", deleteCmsBlockStore)
}

func getCmsBlockStores(c *gin.Context) {
	var cmsBlockStores []CmsBlockStore
	if err := g.Find(&cmsBlockStores).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cmsBlockStores)
	}
}

func getCmsBlockStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var cmsBlockStore CmsBlockStore
	if err := g.Where("id = ?", id).First(&cmsBlockStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cmsBlockStore)
	}
}

func createCmsBlockStore(c *gin.Context) {
	var cmsBlockStore CmsBlockStore
	c.BindJSON(&cmsBlockStore)
	g.Create(&cmsBlockStore)
	c.JSON(200, cmsBlockStore)
}

func updateCmsBlockStore(c *gin.Context) {
	var cmsBlockStore CmsBlockStore
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&cmsBlockStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cmsBlockStore)
	g.Save(&cmsBlockStore)
	c.JSON(200, cmsBlockStore)
}
func deleteCmsBlockStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var cmsBlockStore CmsBlockStore
	d := g.Where("id = ?", id).Delete(&cmsBlockStore)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
