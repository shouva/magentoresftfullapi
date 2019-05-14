package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DownloadableSample :
type DownloadableSample struct {
	SampleId   uint16 `gorm:"column:sample_id;primary_key" form:"sample_id;primary_key" json:"sample_id;primary_key"`
	ProductId  uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	SampleUrl  string `gorm:"column:sample_url" form:"sample_url" json:"sample_url"`
	SampleFile string `gorm:"column:sample_file" form:"sample_file" json:"sample_file"`
	SampleType string `gorm:"column:sample_type" form:"sample_type" json:"sample_type"`
	SortOrder  uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
}

// TableName :
func (*DownloadableSample) TableName() string {
	return "downloadable_sample"
}

// handler create
func initRoutersDownloadableSample(r *gin.Engine, downloadablesample string) {
	route := r.Group("downloadablesample")
	route.GET("/", getDownloadableSamples)
	route.GET("/:id", getDownloadableSample)
	route.POST("/", createDownloadableSample)
	route.PUT("/:id", updateDownloadableSample)
	route.DELETE("/:id", deleteDownloadableSample)
}

func getDownloadableSamples(c *gin.Context) {
	var downloadableSamples []DownloadableSample
	if err := g.Find(&downloadableSamples).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableSamples)
	}
}

func getDownloadableSample(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableSample DownloadableSample
	if err := g.Where("id = ?", id).First(&downloadableSample).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableSample)
	}
}

func createDownloadableSample(c *gin.Context) {
	var downloadableSample DownloadableSample
	c.BindJSON(&downloadableSample)
	g.Create(&downloadableSample)
	c.JSON(200, downloadableSample)
}

func updateDownloadableSample(c *gin.Context) {
	var downloadableSample DownloadableSample
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&downloadableSample).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&downloadableSample)
	g.Save(&downloadableSample)
	c.JSON(200, downloadableSample)
}
func deleteDownloadableSample(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableSample DownloadableSample
	d := g.Where("id = ?", id).Delete(&downloadableSample)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
