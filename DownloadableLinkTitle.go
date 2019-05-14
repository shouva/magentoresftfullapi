package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DownloadableLinkTitle :
type DownloadableLinkTitle struct {
	TitleId uint16 `gorm:"column:title_id;primary_key" form:"title_id;primary_key" json:"title_id;primary_key"`
	LinkId  uint16 `gorm:"column:link_id" form:"link_id" json:"link_id"`
	StoreId uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Title   string `gorm:"column:title" form:"title" json:"title"`
}

// TableName :
func (*DownloadableLinkTitle) TableName() string {
	return "downloadable_link_title"
}

// handler create
func initRoutersDownloadableLinkTitle(r *gin.Engine, downloadablelinktitle string) {
	route := r.Group("downloadablelinktitle")
	route.GET("/", getDownloadableLinkTitles)
	route.GET("/:id", getDownloadableLinkTitle)
	route.POST("/", createDownloadableLinkTitle)
	route.PUT("/:id", updateDownloadableLinkTitle)
	route.DELETE("/:id", deleteDownloadableLinkTitle)
}

func getDownloadableLinkTitles(c *gin.Context) {
	var downloadableLinkTitles []DownloadableLinkTitle
	if err := g.Find(&downloadableLinkTitles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableLinkTitles)
	}
}

func getDownloadableLinkTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableLinkTitle DownloadableLinkTitle
	if err := g.Where("id = ?", id).First(&downloadableLinkTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableLinkTitle)
	}
}

func createDownloadableLinkTitle(c *gin.Context) {
	var downloadableLinkTitle DownloadableLinkTitle
	c.BindJSON(&downloadableLinkTitle)
	g.Create(&downloadableLinkTitle)
	c.JSON(200, downloadableLinkTitle)
}

func updateDownloadableLinkTitle(c *gin.Context) {
	var downloadableLinkTitle DownloadableLinkTitle
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&downloadableLinkTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&downloadableLinkTitle)
	g.Save(&downloadableLinkTitle)
	c.JSON(200, downloadableLinkTitle)
}
func deleteDownloadableLinkTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableLinkTitle DownloadableLinkTitle
	d := g.Where("id = ?", id).Delete(&downloadableLinkTitle)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
