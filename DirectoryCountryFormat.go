package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DirectoryCountryFormat :
type DirectoryCountryFormat struct {
	CountryFormatId uint16 `gorm:"column:country_format_id;primary_key" form:"country_format_id;primary_key" json:"country_format_id;primary_key"`
	CountryId       string `gorm:"column:country_id" form:"country_id" json:"country_id"`
	Type            string `gorm:"column:type" form:"type" json:"type"`
	Format          string `gorm:"column:format" form:"format" json:"format"`
}

// TableName :
func (*DirectoryCountryFormat) TableName() string {
	return "directory_country_format"
}

// handler create
func initRoutersDirectoryCountryFormat(r *gin.Engine, directorycountryformat string) {
	route := r.Group("directorycountryformat")
	route.GET("/", getDirectoryCountryFormats)
	route.GET("/:id", getDirectoryCountryFormat)
	route.POST("/", createDirectoryCountryFormat)
	route.PUT("/:id", updateDirectoryCountryFormat)
	route.DELETE("/:id", deleteDirectoryCountryFormat)
}

func getDirectoryCountryFormats(c *gin.Context) {
	var directoryCountryFormats []DirectoryCountryFormat
	if err := g.Find(&directoryCountryFormats).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, directoryCountryFormats)
	}
}

func getDirectoryCountryFormat(c *gin.Context) {
	id := c.Params.ByName("id")
	var directoryCountryFormat DirectoryCountryFormat
	if err := g.Where("id = ?", id).First(&directoryCountryFormat).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, directoryCountryFormat)
	}
}

func createDirectoryCountryFormat(c *gin.Context) {
	var directoryCountryFormat DirectoryCountryFormat
	c.BindJSON(&directoryCountryFormat)
	g.Create(&directoryCountryFormat)
	c.JSON(200, directoryCountryFormat)
}

func updateDirectoryCountryFormat(c *gin.Context) {
	var directoryCountryFormat DirectoryCountryFormat
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&directoryCountryFormat).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&directoryCountryFormat)
	g.Save(&directoryCountryFormat)
	c.JSON(200, directoryCountryFormat)
}
func deleteDirectoryCountryFormat(c *gin.Context) {
	id := c.Params.ByName("id")
	var directoryCountryFormat DirectoryCountryFormat
	d := g.Where("id = ?", id).Delete(&directoryCountryFormat)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
