package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreStoreGroup :
type CoreStoreGroup struct {
	GroupId        uint16 `gorm:"column:group_id;primary_key" form:"group_id;primary_key" json:"group_id;primary_key"`
	WebsiteId      uint16 `gorm:"column:website_id" form:"website_id" json:"website_id"`
	Name           string `gorm:"column:name" form:"name" json:"name"`
	RootCategoryId uint16 `gorm:"column:root_category_id" form:"root_category_id" json:"root_category_id"`
	DefaultStoreId uint16 `gorm:"column:default_store_id" form:"default_store_id" json:"default_store_id"`
}

// TableName :
func (*CoreStoreGroup) TableName() string {
	return "core_store_group"
}

// handler create
func initRoutersCoreStoreGroup(r *gin.Engine, corestoregroup string) {
	route := r.Group("corestoregroup")
	route.GET("/", getCoreStoreGroups)
	route.GET("/:id", getCoreStoreGroup)
	route.POST("/", createCoreStoreGroup)
	route.PUT("/:id", updateCoreStoreGroup)
	route.DELETE("/:id", deleteCoreStoreGroup)
}

func getCoreStoreGroups(c *gin.Context) {
	var coreStoreGroups []CoreStoreGroup
	if err := g.Find(&coreStoreGroups).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreStoreGroups)
	}
}

func getCoreStoreGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreStoreGroup CoreStoreGroup
	if err := g.Where("id = ?", id).First(&coreStoreGroup).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreStoreGroup)
	}
}

func createCoreStoreGroup(c *gin.Context) {
	var coreStoreGroup CoreStoreGroup
	c.BindJSON(&coreStoreGroup)
	g.Create(&coreStoreGroup)
	c.JSON(200, coreStoreGroup)
}

func updateCoreStoreGroup(c *gin.Context) {
	var coreStoreGroup CoreStoreGroup
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreStoreGroup).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreStoreGroup)
	g.Save(&coreStoreGroup)
	c.JSON(200, coreStoreGroup)
}
func deleteCoreStoreGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreStoreGroup CoreStoreGroup
	d := g.Where("id = ?", id).Delete(&coreStoreGroup)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
