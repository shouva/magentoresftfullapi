package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// TagRelation :
type TagRelation struct {
	TagRelationId uint16     `gorm:"column:tag_relation_id;primary_key" form:"tag_relation_id;primary_key" json:"tag_relation_id;primary_key"`
	TagId         uint16     `gorm:"column:tag_id" form:"tag_id" json:"tag_id"`
	CustomerId    uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	ProductId     uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	StoreId       uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Active        uint16     `gorm:"column:active" form:"active" json:"active"`
	CreatedAt     *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}

// TableName :
func (*TagRelation) TableName() string {
	return "tag_relation"
}

// handler create
func initRoutersTagRelation(r *gin.Engine, tagrelation string) {
	route := r.Group("tagrelation")
	route.GET("/", getTagRelations)
	route.GET("/:id", getTagRelation)
	route.POST("/", createTagRelation)
	route.PUT("/:id", updateTagRelation)
	route.DELETE("/:id", deleteTagRelation)
}

func getTagRelations(c *gin.Context) {
	var tagRelations []TagRelation
	if err := g.Find(&tagRelations).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, tagRelations)
	}
}

func getTagRelation(c *gin.Context) {
	id := c.Params.ByName("id")
	var tagRelation TagRelation
	if err := g.Where("id = ?", id).First(&tagRelation).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, tagRelation)
	}
}

func createTagRelation(c *gin.Context) {
	var tagRelation TagRelation
	c.BindJSON(&tagRelation)
	g.Create(&tagRelation)
	c.JSON(200, tagRelation)
}

func updateTagRelation(c *gin.Context) {
	var tagRelation TagRelation
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&tagRelation).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&tagRelation)
	g.Save(&tagRelation)
	c.JSON(200, tagRelation)
}
func deleteTagRelation(c *gin.Context) {
	id := c.Params.ByName("id")
	var tagRelation TagRelation
	d := g.Where("id = ?", id).Delete(&tagRelation)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
