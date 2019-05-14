package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DownloadableLink :
type DownloadableLink struct {
	LinkId            uint16 `gorm:"column:link_id;primary_key" form:"link_id;primary_key" json:"link_id;primary_key"`
	ProductId         uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	SortOrder         uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	NumberOfDownloads uint16 `gorm:"column:number_of_downloads" form:"number_of_downloads" json:"number_of_downloads"`
	IsShareable       uint16 `gorm:"column:is_shareable" form:"is_shareable" json:"is_shareable"`
	LinkUrl           string `gorm:"column:link_url" form:"link_url" json:"link_url"`
	LinkFile          string `gorm:"column:link_file" form:"link_file" json:"link_file"`
	LinkType          string `gorm:"column:link_type" form:"link_type" json:"link_type"`
	SampleUrl         string `gorm:"column:sample_url" form:"sample_url" json:"sample_url"`
	SampleFile        string `gorm:"column:sample_file" form:"sample_file" json:"sample_file"`
	SampleType        string `gorm:"column:sample_type" form:"sample_type" json:"sample_type"`
}

// TableName :
func (*DownloadableLink) TableName() string {
	return "downloadable_link"
}

// handler create
func initRoutersDownloadableLink(r *gin.Engine, downloadablelink string) {
	route := r.Group("downloadablelink")
	route.GET("/", getDownloadableLinks)
	route.GET("/:id", getDownloadableLink)
	route.POST("/", createDownloadableLink)
	route.PUT("/:id", updateDownloadableLink)
	route.DELETE("/:id", deleteDownloadableLink)
}

func getDownloadableLinks(c *gin.Context) {
	var downloadableLinks []DownloadableLink
	if err := g.Find(&downloadableLinks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableLinks)
	}
}

func getDownloadableLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableLink DownloadableLink
	if err := g.Where("id = ?", id).First(&downloadableLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableLink)
	}
}

func createDownloadableLink(c *gin.Context) {
	var downloadableLink DownloadableLink
	c.BindJSON(&downloadableLink)
	g.Create(&downloadableLink)
	c.JSON(200, downloadableLink)
}

func updateDownloadableLink(c *gin.Context) {
	var downloadableLink DownloadableLink
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&downloadableLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&downloadableLink)
	g.Save(&downloadableLink)
	c.JSON(200, downloadableLink)
}
func deleteDownloadableLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableLink DownloadableLink
	d := g.Where("id = ?", id).Delete(&downloadableLink)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
