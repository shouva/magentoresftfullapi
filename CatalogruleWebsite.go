package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogruleWebsite :
type CatalogruleWebsite struct {
	RuleId    uint16 `gorm:"column:rule_id;primary_key" form:"rule_id;primary_key" json:"rule_id;primary_key"`
	WebsiteId uint16 `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
}

// TableName :
func (*CatalogruleWebsite) TableName() string {
	return "catalogrule_website"
}

// handler create
func initRoutersCatalogruleWebsite(r *gin.Engine, catalogrulewebsite string) {
	route := r.Group("catalogrulewebsite")
	route.GET("/", getCatalogruleWebsites)
	route.GET("/:id", getCatalogruleWebsite)
	route.POST("/", createCatalogruleWebsite)
	route.PUT("/:id", updateCatalogruleWebsite)
	route.DELETE("/:id", deleteCatalogruleWebsite)
}

func getCatalogruleWebsites(c *gin.Context) {
	var catalogruleWebsites []CatalogruleWebsite
	if err := g.Find(&catalogruleWebsites).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleWebsites)
	}
}

func getCatalogruleWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleWebsite CatalogruleWebsite
	if err := g.Where("id = ?", id).First(&catalogruleWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleWebsite)
	}
}

func createCatalogruleWebsite(c *gin.Context) {
	var catalogruleWebsite CatalogruleWebsite
	c.BindJSON(&catalogruleWebsite)
	g.Create(&catalogruleWebsite)
	c.JSON(200, catalogruleWebsite)
}

func updateCatalogruleWebsite(c *gin.Context) {
	var catalogruleWebsite CatalogruleWebsite
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogruleWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogruleWebsite)
	g.Save(&catalogruleWebsite)
	c.JSON(200, catalogruleWebsite)
}
func deleteCatalogruleWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleWebsite CatalogruleWebsite
	d := g.Where("id = ?", id).Delete(&catalogruleWebsite)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
