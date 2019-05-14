package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// EavEntityDatetime :
type EavEntityDatetime struct {
	ValueId      uint16     `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16     `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16     `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16     `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        *time.Time `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*EavEntityDatetime) TableName() string {
	return "eav_entity_datetime"
}

// handler create
func initRoutersEavEntityDatetime(r *gin.Engine, eaventitydatetime string) {
	route := r.Group("eaventitydatetime")
	route.GET("/", getEavEntityDatetimes)
	route.GET("/:id", getEavEntityDatetime)
	route.POST("/", createEavEntityDatetime)
	route.PUT("/:id", updateEavEntityDatetime)
	route.DELETE("/:id", deleteEavEntityDatetime)
}

func getEavEntityDatetimes(c *gin.Context) {
	var eavEntityDatetimes []EavEntityDatetime
	if err := g.Find(&eavEntityDatetimes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityDatetimes)
	}
}

func getEavEntityDatetime(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityDatetime EavEntityDatetime
	if err := g.Where("id = ?", id).First(&eavEntityDatetime).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityDatetime)
	}
}

func createEavEntityDatetime(c *gin.Context) {
	var eavEntityDatetime EavEntityDatetime
	c.BindJSON(&eavEntityDatetime)
	g.Create(&eavEntityDatetime)
	c.JSON(200, eavEntityDatetime)
}

func updateEavEntityDatetime(c *gin.Context) {
	var eavEntityDatetime EavEntityDatetime
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavEntityDatetime).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavEntityDatetime)
	g.Save(&eavEntityDatetime)
	c.JSON(200, eavEntityDatetime)
}
func deleteEavEntityDatetime(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityDatetime EavEntityDatetime
	d := g.Where("id = ?", id).Delete(&eavEntityDatetime)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
