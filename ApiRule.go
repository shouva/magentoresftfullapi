package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ApiRule :
type ApiRule struct {
	RuleId        uint16 `gorm:"column:rule_id;primary_key" form:"rule_id;primary_key" json:"rule_id;primary_key"`
	RoleId        uint16 `gorm:"column:role_id" form:"role_id" json:"role_id"`
	ResourceId    string `gorm:"column:resource_id" form:"resource_id" json:"resource_id"`
	ApiPrivileges string `gorm:"column:api_privileges" form:"api_privileges" json:"api_privileges"`
	AssertId      uint16 `gorm:"column:assert_id" form:"assert_id" json:"assert_id"`
	RoleType      string `gorm:"column:role_type" form:"role_type" json:"role_type"`
	ApiPermission string `gorm:"column:api_permission" form:"api_permission" json:"api_permission"`
}

// TableName :
func (*ApiRule) TableName() string {
	return "api_rule"
}

// handler create
func initRoutersApiRule(r *gin.Engine, apirule string) {
	route := r.Group("apirule")
	route.GET("/", getApiRules)
	route.GET("/:id", getApiRule)
	route.POST("/", createApiRule)
	route.PUT("/:id", updateApiRule)
	route.DELETE("/:id", deleteApiRule)
}

func getApiRules(c *gin.Context) {
	var apiRules []ApiRule
	if err := g.Find(&apiRules).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, apiRules)
	}
}

func getApiRule(c *gin.Context) {
	id := c.Params.ByName("id")
	var apiRule ApiRule
	if err := g.Where("id = ?", id).First(&apiRule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, apiRule)
	}
}

func createApiRule(c *gin.Context) {
	var apiRule ApiRule
	c.BindJSON(&apiRule)
	g.Create(&apiRule)
	c.JSON(200, apiRule)
}

func updateApiRule(c *gin.Context) {
	var apiRule ApiRule
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&apiRule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&apiRule)
	g.Save(&apiRule)
	c.JSON(200, apiRule)
}
func deleteApiRule(c *gin.Context) {
	id := c.Params.ByName("id")
	var apiRule ApiRule
	d := g.Where("id = ?", id).Delete(&apiRule)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
