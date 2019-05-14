package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Api2AclRule :
type Api2AclRule struct {
	EntityId   uint16 `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	RoleId     uint16 `gorm:"column:role_id" form:"role_id" json:"role_id"`
	ResourceId string `gorm:"column:resource_id" form:"resource_id" json:"resource_id"`
	Privilege  string `gorm:"column:privilege" form:"privilege" json:"privilege"`
}

// TableName :
func (*Api2AclRule) TableName() string {
	return "api2_acl_rule"
}

// handler create
func initRoutersApi2AclRule(r *gin.Engine, api2aclrule string) {
	route := r.Group("api2aclrule")
	route.GET("/", getApi2AclRules)
	route.GET("/:id", getApi2AclRule)
	route.POST("/", createApi2AclRule)
	route.PUT("/:id", updateApi2AclRule)
	route.DELETE("/:id", deleteApi2AclRule)
}

func getApi2AclRules(c *gin.Context) {
	var api2AclRules []Api2AclRule
	if err := g.Find(&api2AclRules).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, api2AclRules)
	}
}

func getApi2AclRule(c *gin.Context) {
	id := c.Params.ByName("id")
	var api2AclRule Api2AclRule
	if err := g.Where("id = ?", id).First(&api2AclRule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, api2AclRule)
	}
}

func createApi2AclRule(c *gin.Context) {
	var api2AclRule Api2AclRule
	c.BindJSON(&api2AclRule)
	g.Create(&api2AclRule)
	c.JSON(200, api2AclRule)
}

func updateApi2AclRule(c *gin.Context) {
	var api2AclRule Api2AclRule
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&api2AclRule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&api2AclRule)
	g.Save(&api2AclRule)
	c.JSON(200, api2AclRule)
}
func deleteApi2AclRule(c *gin.Context) {
	id := c.Params.ByName("id")
	var api2AclRule Api2AclRule
	d := g.Where("id = ?", id).Delete(&api2AclRule)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
