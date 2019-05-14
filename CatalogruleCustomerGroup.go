package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogruleCustomerGroup :
type CatalogruleCustomerGroup struct {
	RuleId          uint16 `gorm:"column:rule_id;primary_key" form:"rule_id;primary_key" json:"rule_id;primary_key"`
	CustomerGroupId uint16 `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
}

// TableName :
func (*CatalogruleCustomerGroup) TableName() string {
	return "catalogrule_customer_group"
}

// handler create
func initRoutersCatalogruleCustomerGroup(r *gin.Engine, catalogrulecustomergroup string) {
	route := r.Group("catalogrulecustomergroup")
	route.GET("/", getCatalogruleCustomerGroups)
	route.GET("/:id", getCatalogruleCustomerGroup)
	route.POST("/", createCatalogruleCustomerGroup)
	route.PUT("/:id", updateCatalogruleCustomerGroup)
	route.DELETE("/:id", deleteCatalogruleCustomerGroup)
}

func getCatalogruleCustomerGroups(c *gin.Context) {
	var catalogruleCustomerGroups []CatalogruleCustomerGroup
	if err := g.Find(&catalogruleCustomerGroups).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleCustomerGroups)
	}
}

func getCatalogruleCustomerGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleCustomerGroup CatalogruleCustomerGroup
	if err := g.Where("id = ?", id).First(&catalogruleCustomerGroup).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleCustomerGroup)
	}
}

func createCatalogruleCustomerGroup(c *gin.Context) {
	var catalogruleCustomerGroup CatalogruleCustomerGroup
	c.BindJSON(&catalogruleCustomerGroup)
	g.Create(&catalogruleCustomerGroup)
	c.JSON(200, catalogruleCustomerGroup)
}

func updateCatalogruleCustomerGroup(c *gin.Context) {
	var catalogruleCustomerGroup CatalogruleCustomerGroup
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogruleCustomerGroup).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogruleCustomerGroup)
	g.Save(&catalogruleCustomerGroup)
	c.JSON(200, catalogruleCustomerGroup)
}
func deleteCatalogruleCustomerGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleCustomerGroup CatalogruleCustomerGroup
	d := g.Where("id = ?", id).Delete(&catalogruleCustomerGroup)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
