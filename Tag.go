package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Tag :
type Tag struct {
	TagId           uint16 `gorm:"column:tag_id;primary_key" form:"tag_id;primary_key" json:"tag_id;primary_key"`
	Name            string `gorm:"column:name" form:"name" json:"name"`
	Status          uint16 `gorm:"column:status" form:"status" json:"status"`
	FirstCustomerId uint16 `gorm:"column:first_customer_id" form:"first_customer_id" json:"first_customer_id"`
	FirstStoreId    uint16 `gorm:"column:first_store_id" form:"first_store_id" json:"first_store_id"`
}

// TableName :
func (*Tag) TableName() string {
	return "tag"
}

// handler create
func initRoutersTag(r *gin.Engine, tag string) {
	route := r.Group("tag")
	route.GET("/", getTags)
	route.GET("/:id", getTag)
	route.POST("/", createTag)
	route.PUT("/:id", updateTag)
	route.DELETE("/:id", deleteTag)
}

func getTags(c *gin.Context) {
	var tags []Tag
	if err := g.Find(&tags).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, tags)
	}
}

func getTag(c *gin.Context) {
	id := c.Params.ByName("id")
	var tag Tag
	if err := g.Where("id = ?", id).First(&tag).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, tag)
	}
}

func createTag(c *gin.Context) {
	var tag Tag
	c.BindJSON(&tag)
	g.Create(&tag)
	c.JSON(200, tag)
}

func updateTag(c *gin.Context) {
	var tag Tag
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&tag).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&tag)
	g.Save(&tag)
	c.JSON(200, tag)
}
func deleteTag(c *gin.Context) {
	id := c.Params.ByName("id")
	var tag Tag
	d := g.Where("id = ?", id).Delete(&tag)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
