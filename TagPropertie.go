package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TagPropertie :
type TagPropertie struct {
	TagId          uint16 `gorm:"column:tag_id;primary_key" form:"tag_id;primary_key" json:"tag_id;primary_key"`
	StoreId        uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	BasePopularity uint16 `gorm:"column:base_popularity" form:"base_popularity" json:"base_popularity"`
}

// TableName :
func (*TagPropertie) TableName() string {
	return "tag_properties"
}

// handler create
func initRoutersTagPropertie(r *gin.Engine, tagpropertie string) {
	route := r.Group("tagpropertie")
	route.GET("/", getTagProperties)
	route.GET("/:id", getTagPropertie)
	route.POST("/", createTagPropertie)
	route.PUT("/:id", updateTagPropertie)
	route.DELETE("/:id", deleteTagPropertie)
}

func getTagProperties(c *gin.Context) {
	var tagProperties []TagPropertie
	if err := g.Find(&tagProperties).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, tagProperties)
	}
}

func getTagPropertie(c *gin.Context) {
	id := c.Params.ByName("id")
	var tagPropertie TagPropertie
	if err := g.Where("id = ?", id).First(&tagPropertie).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, tagPropertie)
	}
}

func createTagPropertie(c *gin.Context) {
	var tagPropertie TagPropertie
	c.BindJSON(&tagPropertie)
	g.Create(&tagPropertie)
	c.JSON(200, tagPropertie)
}

func updateTagPropertie(c *gin.Context) {
	var tagPropertie TagPropertie
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&tagPropertie).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&tagPropertie)
	g.Save(&tagPropertie)
	c.JSON(200, tagPropertie)
}
func deleteTagPropertie(c *gin.Context) {
	id := c.Params.ByName("id")
	var tagPropertie TagPropertie
	d := g.Where("id = ?", id).Delete(&tagPropertie)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
