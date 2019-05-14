package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogruleGroupWebsite :
type CatalogruleGroupWebsite struct {
	RuleId          uint16 `gorm:"column:rule_id;primary_key" form:"rule_id;primary_key" json:"rule_id;primary_key"`
	CustomerGroupId uint16 `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16 `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
}

// TableName :
func (*CatalogruleGroupWebsite) TableName() string {
	return "catalogrule_group_website"
}

// handler create
func initRoutersCatalogruleGroupWebsite(r *gin.Engine, catalogrulegroupwebsite string) {
	route := r.Group("catalogrulegroupwebsite")
	route.GET("/", getCatalogruleGroupWebsites)
	route.GET("/:id", getCatalogruleGroupWebsite)
	route.POST("/", createCatalogruleGroupWebsite)
	route.PUT("/:id", updateCatalogruleGroupWebsite)
	route.DELETE("/:id", deleteCatalogruleGroupWebsite)
}

func getCatalogruleGroupWebsites(c *gin.Context) {
	var catalogruleGroupWebsites []CatalogruleGroupWebsite
	if err := g.Find(&catalogruleGroupWebsites).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleGroupWebsites)
	}
}

func getCatalogruleGroupWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleGroupWebsite CatalogruleGroupWebsite
	if err := g.Where("id = ?", id).First(&catalogruleGroupWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleGroupWebsite)
	}
}

func createCatalogruleGroupWebsite(c *gin.Context) {
	var catalogruleGroupWebsite CatalogruleGroupWebsite
	c.BindJSON(&catalogruleGroupWebsite)
	g.Create(&catalogruleGroupWebsite)
	c.JSON(200, catalogruleGroupWebsite)
}

func updateCatalogruleGroupWebsite(c *gin.Context) {
	var catalogruleGroupWebsite CatalogruleGroupWebsite
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogruleGroupWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogruleGroupWebsite)
	g.Save(&catalogruleGroupWebsite)
	c.JSON(200, catalogruleGroupWebsite)
}
func deleteCatalogruleGroupWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleGroupWebsite CatalogruleGroupWebsite
	d := g.Where("id = ?", id).Delete(&catalogruleGroupWebsite)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
