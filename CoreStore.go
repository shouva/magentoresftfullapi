package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreStore :
type CoreStore struct {
	StoreId   uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Code      string `gorm:"column:code" form:"code" json:"code"`
	WebsiteId uint16 `gorm:"column:website_id" form:"website_id" json:"website_id"`
	GroupId   uint16 `gorm:"column:group_id" form:"group_id" json:"group_id"`
	Name      string `gorm:"column:name" form:"name" json:"name"`
	SortOrder uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	IsActive  uint16 `gorm:"column:is_active" form:"is_active" json:"is_active"`
}

// TableName :
func (*CoreStore) TableName() string {
	return "core_store"
}

// handler create
func initRoutersCoreStore(r *gin.Engine, corestore string) {
	route := r.Group("corestore")
	route.GET("/", getCoreStores)
	route.GET("/:id", getCoreStore)
	route.POST("/", createCoreStore)
	route.PUT("/:id", updateCoreStore)
	route.DELETE("/:id", deleteCoreStore)
}

func getCoreStores(c *gin.Context) {
	var coreStores []CoreStore
	if err := g.Find(&coreStores).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreStores)
	}
}

func getCoreStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreStore CoreStore
	if err := g.Where("id = ?", id).First(&coreStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreStore)
	}
}

func createCoreStore(c *gin.Context) {
	var coreStore CoreStore
	c.BindJSON(&coreStore)
	g.Create(&coreStore)
	c.JSON(200, coreStore)
}

func updateCoreStore(c *gin.Context) {
	var coreStore CoreStore
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreStore)
	g.Save(&coreStore)
	c.JSON(200, coreStore)
}
func deleteCoreStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreStore CoreStore
	d := g.Where("id = ?", id).Delete(&coreStore)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
