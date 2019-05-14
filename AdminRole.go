package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// AdminRole :
type AdminRole struct {
	RoleId    uint16 `gorm:"column:role_id;primary_key" form:"role_id;primary_key" json:"role_id;primary_key"`
	ParentId  uint16 `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	TreeLevel uint16 `gorm:"column:tree_level" form:"tree_level" json:"tree_level"`
	SortOrder uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	RoleType  string `gorm:"column:role_type" form:"role_type" json:"role_type"`
	UserId    uint16 `gorm:"column:user_id" form:"user_id" json:"user_id"`
	RoleName  string `gorm:"column:role_name" form:"role_name" json:"role_name"`
}

// TableName :
func (*AdminRole) TableName() string {
	return "admin_role"
}

// handler create
func initRoutersAdminRole(r *gin.Engine, adminrole string) {
	route := r.Group("adminrole")
	route.GET("/", getAdminRoles)
	route.GET("/:id", getAdminRole)
	route.POST("/", createAdminRole)
	route.PUT("/:id", updateAdminRole)
	route.DELETE("/:id", deleteAdminRole)
}

func getAdminRoles(c *gin.Context) {
	var adminRoles []AdminRole
	if err := g.Find(&adminRoles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, adminRoles)
	}
}

func getAdminRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var adminRole AdminRole
	if err := g.Where("id = ?", id).First(&adminRole).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, adminRole)
	}
}

func createAdminRole(c *gin.Context) {
	var adminRole AdminRole
	c.BindJSON(&adminRole)
	g.Create(&adminRole)
	c.JSON(200, adminRole)
}

func updateAdminRole(c *gin.Context) {
	var adminRole AdminRole
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&adminRole).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&adminRole)
	g.Save(&adminRole)
	c.JSON(200, adminRole)
}
func deleteAdminRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var adminRole AdminRole
	d := g.Where("id = ?", id).Delete(&adminRole)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
