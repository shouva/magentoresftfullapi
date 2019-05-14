package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavEntityStore :
type EavEntityStore struct {
	EntityStoreId   uint16 `gorm:"column:entity_store_id;primary_key" form:"entity_store_id;primary_key" json:"entity_store_id;primary_key"`
	EntityTypeId    uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	StoreId         uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	IncrementPrefix string `gorm:"column:increment_prefix" form:"increment_prefix" json:"increment_prefix"`
	IncrementLastId string `gorm:"column:increment_last_id" form:"increment_last_id" json:"increment_last_id"`
}

// TableName :
func (*EavEntityStore) TableName() string {
	return "eav_entity_store"
}

// handler create
func initRoutersEavEntityStore(r *gin.Engine, eaventitystore string) {
	route := r.Group("eaventitystore")
	route.GET("/", getEavEntityStores)
	route.GET("/:id", getEavEntityStore)
	route.POST("/", createEavEntityStore)
	route.PUT("/:id", updateEavEntityStore)
	route.DELETE("/:id", deleteEavEntityStore)
}

func getEavEntityStores(c *gin.Context) {
	var eavEntityStores []EavEntityStore
	if err := g.Find(&eavEntityStores).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityStores)
	}
}

func getEavEntityStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityStore EavEntityStore
	if err := g.Where("id = ?", id).First(&eavEntityStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityStore)
	}
}

func createEavEntityStore(c *gin.Context) {
	var eavEntityStore EavEntityStore
	c.BindJSON(&eavEntityStore)
	g.Create(&eavEntityStore)
	c.JSON(200, eavEntityStore)
}

func updateEavEntityStore(c *gin.Context) {
	var eavEntityStore EavEntityStore
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavEntityStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavEntityStore)
	g.Save(&eavEntityStore)
	c.JSON(200, eavEntityStore)
}
func deleteEavEntityStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityStore EavEntityStore
	d := g.Where("id = ?", id).Delete(&eavEntityStore)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
