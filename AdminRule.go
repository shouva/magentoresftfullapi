package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// AdminRule :
type AdminRule struct {
	RuleId     uint16 `gorm:"column:rule_id;primary_key" form:"rule_id;primary_key" json:"rule_id;primary_key"`
	RoleId     uint16 `gorm:"column:role_id" form:"role_id" json:"role_id"`
	ResourceId string `gorm:"column:resource_id" form:"resource_id" json:"resource_id"`
	Privileges string `gorm:"column:privileges" form:"privileges" json:"privileges"`
	AssertId   uint16 `gorm:"column:assert_id" form:"assert_id" json:"assert_id"`
	RoleType   string `gorm:"column:role_type" form:"role_type" json:"role_type"`
	Permission string `gorm:"column:permission" form:"permission" json:"permission"`
}

// TableName :
func (*AdminRule) TableName() string {
	return "admin_rule"
}

// handler create
func initRoutersAdminRule(r *gin.Engine, adminrule string) {
	route := r.Group("adminrule")
	route.GET("/", getAdminRules)
	route.GET("/:id", getAdminRule)
	route.POST("/", createAdminRule)
	route.PUT("/:id", updateAdminRule)
	route.DELETE("/:id", deleteAdminRule)
}

func getAdminRules(c *gin.Context) {
	var adminRules []AdminRule
	if err := g.Find(&adminRules).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, adminRules)
	}
}

func getAdminRule(c *gin.Context) {
	id := c.Params.ByName("id")
	var adminRule AdminRule
	if err := g.Where("id = ?", id).First(&adminRule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, adminRule)
	}
}

func createAdminRule(c *gin.Context) {
	var adminRule AdminRule
	c.BindJSON(&adminRule)
	g.Create(&adminRule)
	c.JSON(200, adminRule)
}

func updateAdminRule(c *gin.Context) {
	var adminRule AdminRule
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&adminRule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&adminRule)
	g.Save(&adminRule)
	c.JSON(200, adminRule)
}
func deleteAdminRule(c *gin.Context) {
	id := c.Params.ByName("id")
	var adminRule AdminRule
	d := g.Where("id = ?", id).Delete(&adminRule)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
