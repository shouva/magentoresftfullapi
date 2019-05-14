package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// DesignChange :
type DesignChange struct {
	DesignChangeId uint16     `gorm:"column:design_change_id;primary_key" form:"design_change_id;primary_key" json:"design_change_id;primary_key"`
	StoreId        uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Design         string     `gorm:"column:design" form:"design" json:"design"`
	DateFrom       *time.Time `gorm:"column:date_from" form:"date_from" json:"date_from"`
	DateTo         *time.Time `gorm:"column:date_to" form:"date_to" json:"date_to"`
}

// TableName :
func (*DesignChange) TableName() string {
	return "design_change"
}

// handler create
func initRoutersDesignChange(r *gin.Engine, designchange string) {
	route := r.Group("designchange")
	route.GET("/", getDesignChanges)
	route.GET("/:id", getDesignChange)
	route.POST("/", createDesignChange)
	route.PUT("/:id", updateDesignChange)
	route.DELETE("/:id", deleteDesignChange)
}

func getDesignChanges(c *gin.Context) {
	var designChanges []DesignChange
	if err := g.Find(&designChanges).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, designChanges)
	}
}

func getDesignChange(c *gin.Context) {
	id := c.Params.ByName("id")
	var designChange DesignChange
	if err := g.Where("id = ?", id).First(&designChange).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, designChange)
	}
}

func createDesignChange(c *gin.Context) {
	var designChange DesignChange
	c.BindJSON(&designChange)
	g.Create(&designChange)
	c.JSON(200, designChange)
}

func updateDesignChange(c *gin.Context) {
	var designChange DesignChange
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&designChange).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&designChange)
	g.Save(&designChange)
	c.JSON(200, designChange)
}
func deleteDesignChange(c *gin.Context) {
	id := c.Params.ByName("id")
	var designChange DesignChange
	d := g.Where("id = ?", id).Delete(&designChange)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
