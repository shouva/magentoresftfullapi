package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DirectoryCountry :
type DirectoryCountry struct {
	CountryId string `gorm:"column:country_id;primary_key" form:"country_id;primary_key" json:"country_id;primary_key"`
	Iso2Code  string `gorm:"column:iso2_code" form:"iso2_code" json:"iso2_code"`
	Iso3Code  string `gorm:"column:iso3_code" form:"iso3_code" json:"iso3_code"`
}

// TableName :
func (*DirectoryCountry) TableName() string {
	return "directory_country"
}

// handler create
func initRoutersDirectoryCountry(r *gin.Engine, directorycountry string) {
	route := r.Group("directorycountry")
	route.GET("/", getDirectoryCountrys)
	route.GET("/:id", getDirectoryCountry)
	route.POST("/", createDirectoryCountry)
	route.PUT("/:id", updateDirectoryCountry)
	route.DELETE("/:id", deleteDirectoryCountry)
}

func getDirectoryCountrys(c *gin.Context) {
	var directoryCountrys []DirectoryCountry
	if err := g.Find(&directoryCountrys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, directoryCountrys)
	}
}

func getDirectoryCountry(c *gin.Context) {
	id := c.Params.ByName("id")
	var directoryCountry DirectoryCountry
	if err := g.Where("id = ?", id).First(&directoryCountry).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, directoryCountry)
	}
}

func createDirectoryCountry(c *gin.Context) {
	var directoryCountry DirectoryCountry
	c.BindJSON(&directoryCountry)
	g.Create(&directoryCountry)
	c.JSON(200, directoryCountry)
}

func updateDirectoryCountry(c *gin.Context) {
	var directoryCountry DirectoryCountry
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&directoryCountry).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&directoryCountry)
	g.Save(&directoryCountry)
	c.JSON(200, directoryCountry)
}
func deleteDirectoryCountry(c *gin.Context) {
	id := c.Params.ByName("id")
	var directoryCountry DirectoryCountry
	d := g.Where("id = ?", id).Delete(&directoryCountry)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
