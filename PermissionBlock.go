package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// PermissionBlock :
type PermissionBlock struct {
	BlockId   uint16 `gorm:"column:block_id;primary_key" form:"block_id;primary_key" json:"block_id;primary_key"`
	BlockName string `gorm:"column:block_name" form:"block_name" json:"block_name"`
	IsAllowed uint16 `gorm:"column:is_allowed" form:"is_allowed" json:"is_allowed"`
}

// TableName :
func (*PermissionBlock) TableName() string {
	return "permission_block"
}

// handler create
func initRoutersPermissionBlock(r *gin.Engine, permissionblock string) {
	route := r.Group("permissionblock")
	route.GET("/", getPermissionBlocks)
	route.GET("/:id", getPermissionBlock)
	route.POST("/", createPermissionBlock)
	route.PUT("/:id", updatePermissionBlock)
	route.DELETE("/:id", deletePermissionBlock)
}

func getPermissionBlocks(c *gin.Context) {
	var permissionBlocks []PermissionBlock
	if err := g.Find(&permissionBlocks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, permissionBlocks)
	}
}

func getPermissionBlock(c *gin.Context) {
	id := c.Params.ByName("id")
	var permissionBlock PermissionBlock
	if err := g.Where("id = ?", id).First(&permissionBlock).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, permissionBlock)
	}
}

func createPermissionBlock(c *gin.Context) {
	var permissionBlock PermissionBlock
	c.BindJSON(&permissionBlock)
	g.Create(&permissionBlock)
	c.JSON(200, permissionBlock)
}

func updatePermissionBlock(c *gin.Context) {
	var permissionBlock PermissionBlock
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&permissionBlock).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&permissionBlock)
	g.Save(&permissionBlock)
	c.JSON(200, permissionBlock)
}
func deletePermissionBlock(c *gin.Context) {
	id := c.Params.ByName("id")
	var permissionBlock PermissionBlock
	d := g.Where("id = ?", id).Delete(&permissionBlock)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
