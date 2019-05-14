package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesruleWebsite :
type SalesruleWebsite struct {
	RuleId    uint16 `gorm:"column:rule_id;primary_key" form:"rule_id;primary_key" json:"rule_id;primary_key"`
	WebsiteId uint16 `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
}

// TableName :
func (*SalesruleWebsite) TableName() string {
	return "salesrule_website"
}

// handler create
func initRoutersSalesruleWebsite(r *gin.Engine, salesrulewebsite string) {
	route := r.Group("salesrulewebsite")
	route.GET("/", getSalesruleWebsites)
	route.GET("/:id", getSalesruleWebsite)
	route.POST("/", createSalesruleWebsite)
	route.PUT("/:id", updateSalesruleWebsite)
	route.DELETE("/:id", deleteSalesruleWebsite)
}

func getSalesruleWebsites(c *gin.Context) {
	var salesruleWebsites []SalesruleWebsite
	if err := g.Find(&salesruleWebsites).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleWebsites)
	}
}

func getSalesruleWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleWebsite SalesruleWebsite
	if err := g.Where("id = ?", id).First(&salesruleWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleWebsite)
	}
}

func createSalesruleWebsite(c *gin.Context) {
	var salesruleWebsite SalesruleWebsite
	c.BindJSON(&salesruleWebsite)
	g.Create(&salesruleWebsite)
	c.JSON(200, salesruleWebsite)
}

func updateSalesruleWebsite(c *gin.Context) {
	var salesruleWebsite SalesruleWebsite
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesruleWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesruleWebsite)
	g.Save(&salesruleWebsite)
	c.JSON(200, salesruleWebsite)
}
func deleteSalesruleWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleWebsite SalesruleWebsite
	d := g.Where("id = ?", id).Delete(&salesruleWebsite)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
