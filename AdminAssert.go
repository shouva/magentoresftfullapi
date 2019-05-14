package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// AdminAssert :
type AdminAssert struct {
	AssertId   uint16 `gorm:"column:assert_id;primary_key" form:"assert_id;primary_key" json:"assert_id;primary_key"`
	AssertType string `gorm:"column:assert_type" form:"assert_type" json:"assert_type"`
	AssertData string `gorm:"column:assert_data" form:"assert_data" json:"assert_data"`
}

// TableName :
func (*AdminAssert) TableName() string {
	return "admin_assert"
}

// handler create
func initRoutersAdminAssert(r *gin.Engine, adminassert string) {
	route := r.Group("adminassert")
	route.GET("/", getAdminAsserts)
	route.GET("/:id", getAdminAssert)
	route.POST("/", createAdminAssert)
	route.PUT("/:id", updateAdminAssert)
	route.DELETE("/:id", deleteAdminAssert)
}

func getAdminAsserts(c *gin.Context) {
	var adminAsserts []AdminAssert
	if err := g.Find(&adminAsserts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, adminAsserts)
	}
}

func getAdminAssert(c *gin.Context) {
	id := c.Params.ByName("id")
	var adminAssert AdminAssert
	if err := g.Where("id = ?", id).First(&adminAssert).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, adminAssert)
	}
}

func createAdminAssert(c *gin.Context) {
	var adminAssert AdminAssert
	c.BindJSON(&adminAssert)
	g.Create(&adminAssert)
	c.JSON(200, adminAssert)
}

func updateAdminAssert(c *gin.Context) {
	var adminAssert AdminAssert
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&adminAssert).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&adminAssert)
	g.Save(&adminAssert)
	c.JSON(200, adminAssert)
}
func deleteAdminAssert(c *gin.Context) {
	id := c.Params.ByName("id")
	var adminAssert AdminAssert
	d := g.Where("id = ?", id).Delete(&adminAssert)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
