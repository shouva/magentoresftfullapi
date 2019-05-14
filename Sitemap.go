package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Sitemap :
type Sitemap struct {
	SitemapId       uint16     `gorm:"column:sitemap_id;primary_key" form:"sitemap_id;primary_key" json:"sitemap_id;primary_key"`
	SitemapType     string     `gorm:"column:sitemap_type" form:"sitemap_type" json:"sitemap_type"`
	SitemapFilename string     `gorm:"column:sitemap_filename" form:"sitemap_filename" json:"sitemap_filename"`
	SitemapPath     string     `gorm:"column:sitemap_path" form:"sitemap_path" json:"sitemap_path"`
	SitemapTime     *time.Time `gorm:"column:sitemap_time" form:"sitemap_time" json:"sitemap_time"`
	StoreId         uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
}

// TableName :
func (*Sitemap) TableName() string {
	return "sitemap"
}

// handler create
func initRoutersSitemap(r *gin.Engine, sitemap string) {
	route := r.Group("sitemap")
	route.GET("/", getSitemaps)
	route.GET("/:id", getSitemap)
	route.POST("/", createSitemap)
	route.PUT("/:id", updateSitemap)
	route.DELETE("/:id", deleteSitemap)
}

func getSitemaps(c *gin.Context) {
	var sitemaps []Sitemap
	if err := g.Find(&sitemaps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, sitemaps)
	}
}

func getSitemap(c *gin.Context) {
	id := c.Params.ByName("id")
	var sitemap Sitemap
	if err := g.Where("id = ?", id).First(&sitemap).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, sitemap)
	}
}

func createSitemap(c *gin.Context) {
	var sitemap Sitemap
	c.BindJSON(&sitemap)
	g.Create(&sitemap)
	c.JSON(200, sitemap)
}

func updateSitemap(c *gin.Context) {
	var sitemap Sitemap
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&sitemap).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&sitemap)
	g.Save(&sitemap)
	c.JSON(200, sitemap)
}
func deleteSitemap(c *gin.Context) {
	id := c.Params.ByName("id")
	var sitemap Sitemap
	d := g.Where("id = ?", id).Delete(&sitemap)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
