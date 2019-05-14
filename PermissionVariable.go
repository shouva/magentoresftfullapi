package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// PermissionVariable :
type PermissionVariable struct {
	VariableId   uint16 `gorm:"column:variable_id;primary_key" form:"variable_id;primary_key" json:"variable_id;primary_key"`
	VariableName string `gorm:"column:variable_name;primary_key" form:"variable_name;primary_key" json:"variable_name;primary_key"`
	IsAllowed    uint16 `gorm:"column:is_allowed" form:"is_allowed" json:"is_allowed"`
}

// TableName :
func (*PermissionVariable) TableName() string {
	return "permission_variable"
}

// handler create
func initRoutersPermissionVariable(r *gin.Engine, permissionvariable string) {
	route := r.Group("permissionvariable")
	route.GET("/", getPermissionVariables)
	route.GET("/:id", getPermissionVariable)
	route.POST("/", createPermissionVariable)
	route.PUT("/:id", updatePermissionVariable)
	route.DELETE("/:id", deletePermissionVariable)
}

func getPermissionVariables(c *gin.Context) {
	var permissionVariables []PermissionVariable
	if err := g.Find(&permissionVariables).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, permissionVariables)
	}
}

func getPermissionVariable(c *gin.Context) {
	id := c.Params.ByName("id")
	var permissionVariable PermissionVariable
	if err := g.Where("id = ?", id).First(&permissionVariable).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, permissionVariable)
	}
}

func createPermissionVariable(c *gin.Context) {
	var permissionVariable PermissionVariable
	c.BindJSON(&permissionVariable)
	g.Create(&permissionVariable)
	c.JSON(200, permissionVariable)
}

func updatePermissionVariable(c *gin.Context) {
	var permissionVariable PermissionVariable
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&permissionVariable).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&permissionVariable)
	g.Save(&permissionVariable)
	c.JSON(200, permissionVariable)
}
func deletePermissionVariable(c *gin.Context) {
	id := c.Params.ByName("id")
	var permissionVariable PermissionVariable
	d := g.Where("id = ?", id).Delete(&permissionVariable)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
