package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DownloadableSampleTitle :
type DownloadableSampleTitle struct {
	TitleId  uint16 `gorm:"column:title_id;primary_key" form:"title_id;primary_key" json:"title_id;primary_key"`
	SampleId uint16 `gorm:"column:sample_id" form:"sample_id" json:"sample_id"`
	StoreId  uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Title    string `gorm:"column:title" form:"title" json:"title"`
}

// TableName :
func (*DownloadableSampleTitle) TableName() string {
	return "downloadable_sample_title"
}

// handler create
func initRoutersDownloadableSampleTitle(r *gin.Engine, downloadablesampletitle string) {
	route := r.Group("downloadablesampletitle")
	route.GET("/", getDownloadableSampleTitles)
	route.GET("/:id", getDownloadableSampleTitle)
	route.POST("/", createDownloadableSampleTitle)
	route.PUT("/:id", updateDownloadableSampleTitle)
	route.DELETE("/:id", deleteDownloadableSampleTitle)
}

func getDownloadableSampleTitles(c *gin.Context) {
	var downloadableSampleTitles []DownloadableSampleTitle
	if err := g.Find(&downloadableSampleTitles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableSampleTitles)
	}
}

func getDownloadableSampleTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableSampleTitle DownloadableSampleTitle
	if err := g.Where("id = ?", id).First(&downloadableSampleTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableSampleTitle)
	}
}

func createDownloadableSampleTitle(c *gin.Context) {
	var downloadableSampleTitle DownloadableSampleTitle
	c.BindJSON(&downloadableSampleTitle)
	g.Create(&downloadableSampleTitle)
	c.JSON(200, downloadableSampleTitle)
}

func updateDownloadableSampleTitle(c *gin.Context) {
	var downloadableSampleTitle DownloadableSampleTitle
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&downloadableSampleTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&downloadableSampleTitle)
	g.Save(&downloadableSampleTitle)
	c.JSON(200, downloadableSampleTitle)
}
func deleteDownloadableSampleTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableSampleTitle DownloadableSampleTitle
	d := g.Where("id = ?", id).Delete(&downloadableSampleTitle)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
