package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ApiAssert :
type ApiAssert struct {
	AssertId   uint16 `gorm:"column:assert_id;primary_key" form:"assert_id;primary_key" json:"assert_id;primary_key"`
	AssertType string `gorm:"column:assert_type" form:"assert_type" json:"assert_type"`
	AssertData string `gorm:"column:assert_data" form:"assert_data" json:"assert_data"`
}

// TableName :
func (*ApiAssert) TableName() string {
	return "api_assert"
}

// handler create
func initRoutersApiAssert(r *gin.Engine, apiassert string) {
	route := r.Group("apiassert")
	route.GET("/", getApiAsserts)
	route.GET("/:id", getApiAssert)
	route.POST("/", createApiAssert)
	route.PUT("/:id", updateApiAssert)
	route.DELETE("/:id", deleteApiAssert)
}

func getApiAsserts(c *gin.Context) {
	var apiAsserts []ApiAssert
	if err := g.Find(&apiAsserts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, apiAsserts)
	}
}

func getApiAssert(c *gin.Context) {
	id := c.Params.ByName("id")
	var apiAssert ApiAssert
	if err := g.Where("id = ?", id).First(&apiAssert).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, apiAssert)
	}
}

func createApiAssert(c *gin.Context) {
	var apiAssert ApiAssert
	c.BindJSON(&apiAssert)
	g.Create(&apiAssert)
	c.JSON(200, apiAssert)
}

func updateApiAssert(c *gin.Context) {
	var apiAssert ApiAssert
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&apiAssert).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&apiAssert)
	g.Save(&apiAssert)
	c.JSON(200, apiAssert)
}
func deleteApiAssert(c *gin.Context) {
	id := c.Params.ByName("id")
	var apiAssert ApiAssert
	d := g.Where("id = ?", id).Delete(&apiAssert)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
