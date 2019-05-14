package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CmsPageStore :
type CmsPageStore struct {
	PageId  uint16 `gorm:"column:page_id;primary_key" form:"page_id;primary_key" json:"page_id;primary_key"`
	StoreId uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
}

// TableName :
func (*CmsPageStore) TableName() string {
	return "cms_page_store"
}

// handler create
func initRoutersCmsPageStore(r *gin.Engine, cmspagestore string) {
	route := r.Group("cmspagestore")
	route.GET("/", getCmsPageStores)
	route.GET("/:id", getCmsPageStore)
	route.POST("/", createCmsPageStore)
	route.PUT("/:id", updateCmsPageStore)
	route.DELETE("/:id", deleteCmsPageStore)
}

func getCmsPageStores(c *gin.Context) {
	var cmsPageStores []CmsPageStore
	if err := g.Find(&cmsPageStores).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cmsPageStores)
	}
}

func getCmsPageStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var cmsPageStore CmsPageStore
	if err := g.Where("id = ?", id).First(&cmsPageStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cmsPageStore)
	}
}

func createCmsPageStore(c *gin.Context) {
	var cmsPageStore CmsPageStore
	c.BindJSON(&cmsPageStore)
	g.Create(&cmsPageStore)
	c.JSON(200, cmsPageStore)
}

func updateCmsPageStore(c *gin.Context) {
	var cmsPageStore CmsPageStore
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&cmsPageStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cmsPageStore)
	g.Save(&cmsPageStore)
	c.JSON(200, cmsPageStore)
}
func deleteCmsPageStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var cmsPageStore CmsPageStore
	d := g.Where("id = ?", id).Delete(&cmsPageStore)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
