package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DirectoryCountryRegionName :
type DirectoryCountryRegionName struct {
	Locale   string `gorm:"column:locale;primary_key" form:"locale;primary_key" json:"locale;primary_key"`
	RegionId uint16 `gorm:"column:region_id;primary_key" form:"region_id;primary_key" json:"region_id;primary_key"`
	Name     string `gorm:"column:name" form:"name" json:"name"`
}

// TableName :
func (*DirectoryCountryRegionName) TableName() string {
	return "directory_country_region_name"
}

// handler create
func initRoutersDirectoryCountryRegionName(r *gin.Engine, directorycountryregionname string) {
	route := r.Group("directorycountryregionname")
	route.GET("/", getDirectoryCountryRegionNames)
	route.GET("/:id", getDirectoryCountryRegionName)
	route.POST("/", createDirectoryCountryRegionName)
	route.PUT("/:id", updateDirectoryCountryRegionName)
	route.DELETE("/:id", deleteDirectoryCountryRegionName)
}

func getDirectoryCountryRegionNames(c *gin.Context) {
	var directoryCountryRegionNames []DirectoryCountryRegionName
	if err := g.Find(&directoryCountryRegionNames).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, directoryCountryRegionNames)
	}
}

func getDirectoryCountryRegionName(c *gin.Context) {
	id := c.Params.ByName("id")
	var directoryCountryRegionName DirectoryCountryRegionName
	if err := g.Where("id = ?", id).First(&directoryCountryRegionName).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, directoryCountryRegionName)
	}
}

func createDirectoryCountryRegionName(c *gin.Context) {
	var directoryCountryRegionName DirectoryCountryRegionName
	c.BindJSON(&directoryCountryRegionName)
	g.Create(&directoryCountryRegionName)
	c.JSON(200, directoryCountryRegionName)
}

func updateDirectoryCountryRegionName(c *gin.Context) {
	var directoryCountryRegionName DirectoryCountryRegionName
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&directoryCountryRegionName).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&directoryCountryRegionName)
	g.Save(&directoryCountryRegionName)
	c.JSON(200, directoryCountryRegionName)
}
func deleteDirectoryCountryRegionName(c *gin.Context) {
	id := c.Params.ByName("id")
	var directoryCountryRegionName DirectoryCountryRegionName
	d := g.Where("id = ?", id).Delete(&directoryCountryRegionName)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
