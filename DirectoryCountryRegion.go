package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DirectoryCountryRegion :
type DirectoryCountryRegion struct {
	RegionId    uint16 `gorm:"column:region_id;primary_key" form:"region_id;primary_key" json:"region_id;primary_key"`
	CountryId   string `gorm:"column:country_id" form:"country_id" json:"country_id"`
	Code        string `gorm:"column:code" form:"code" json:"code"`
	DefaultName string `gorm:"column:default_name" form:"default_name" json:"default_name"`
}

// TableName :
func (*DirectoryCountryRegion) TableName() string {
	return "directory_country_region"
}

// handler create
func initRoutersDirectoryCountryRegion(r *gin.Engine, directorycountryregion string) {
	route := r.Group("directorycountryregion")
	route.GET("/", getDirectoryCountryRegions)
	route.GET("/:id", getDirectoryCountryRegion)
	route.POST("/", createDirectoryCountryRegion)
	route.PUT("/:id", updateDirectoryCountryRegion)
	route.DELETE("/:id", deleteDirectoryCountryRegion)
}

func getDirectoryCountryRegions(c *gin.Context) {
	var directoryCountryRegions []DirectoryCountryRegion
	if err := g.Find(&directoryCountryRegions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, directoryCountryRegions)
	}
}

func getDirectoryCountryRegion(c *gin.Context) {
	id := c.Params.ByName("id")
	var directoryCountryRegion DirectoryCountryRegion
	if err := g.Where("id = ?", id).First(&directoryCountryRegion).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, directoryCountryRegion)
	}
}

func createDirectoryCountryRegion(c *gin.Context) {
	var directoryCountryRegion DirectoryCountryRegion
	c.BindJSON(&directoryCountryRegion)
	g.Create(&directoryCountryRegion)
	c.JSON(200, directoryCountryRegion)
}

func updateDirectoryCountryRegion(c *gin.Context) {
	var directoryCountryRegion DirectoryCountryRegion
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&directoryCountryRegion).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&directoryCountryRegion)
	g.Save(&directoryCountryRegion)
	c.JSON(200, directoryCountryRegion)
}
func deleteDirectoryCountryRegion(c *gin.Context) {
	id := c.Params.ByName("id")
	var directoryCountryRegion DirectoryCountryRegion
	d := g.Where("id = ?", id).Delete(&directoryCountryRegion)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
