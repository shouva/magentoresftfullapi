package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ApiRole :
type ApiRole struct {
	RoleId    uint16 `gorm:"column:role_id;primary_key" form:"role_id;primary_key" json:"role_id;primary_key"`
	ParentId  uint16 `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	TreeLevel uint16 `gorm:"column:tree_level" form:"tree_level" json:"tree_level"`
	SortOrder uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	RoleType  string `gorm:"column:role_type" form:"role_type" json:"role_type"`
	UserId    uint16 `gorm:"column:user_id" form:"user_id" json:"user_id"`
	RoleName  string `gorm:"column:role_name" form:"role_name" json:"role_name"`
}

// TableName :
func (*ApiRole) TableName() string {
	return "api_role"
}

// handler create
func initRoutersApiRole(r *gin.Engine, apirole string) {
	route := r.Group("apirole")
	route.GET("/", getApiRoles)
	route.GET("/:id", getApiRole)
	route.POST("/", createApiRole)
	route.PUT("/:id", updateApiRole)
	route.DELETE("/:id", deleteApiRole)
}

func getApiRoles(c *gin.Context) {
	var apiRoles []ApiRole
	if err := g.Find(&apiRoles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, apiRoles)
	}
}

func getApiRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var apiRole ApiRole
	if err := g.Where("id = ?", id).First(&apiRole).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, apiRole)
	}
}

func createApiRole(c *gin.Context) {
	var apiRole ApiRole
	c.BindJSON(&apiRole)
	g.Create(&apiRole)
	c.JSON(200, apiRole)
}

func updateApiRole(c *gin.Context) {
	var apiRole ApiRole
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&apiRole).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&apiRole)
	g.Save(&apiRole)
	c.JSON(200, apiRole)
}
func deleteApiRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var apiRole ApiRole
	d := g.Where("id = ?", id).Delete(&apiRole)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
