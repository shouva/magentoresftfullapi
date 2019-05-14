package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CustomerAddressEntity :
type CustomerAddressEntity struct {
	EntityId       uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	EntityTypeId   uint16     `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeSetId uint16     `gorm:"column:attribute_set_id" form:"attribute_set_id" json:"attribute_set_id"`
	IncrementId    string     `gorm:"column:increment_id" form:"increment_id" json:"increment_id"`
	ParentId       uint16     `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	CreatedAt      *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	IsActive       uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
}

// TableName :
func (*CustomerAddressEntity) TableName() string {
	return "customer_address_entity"
}

// handler create
func initRoutersCustomerAddressEntity(r *gin.Engine, customeraddressentity string) {
	route := r.Group("customeraddressentity")
	route.GET("/", getCustomerAddressEntitys)
	route.GET("/:id", getCustomerAddressEntity)
	route.POST("/", createCustomerAddressEntity)
	route.PUT("/:id", updateCustomerAddressEntity)
	route.DELETE("/:id", deleteCustomerAddressEntity)
}

func getCustomerAddressEntitys(c *gin.Context) {
	var customerAddressEntitys []CustomerAddressEntity
	if err := g.Find(&customerAddressEntitys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntitys)
	}
}

func getCustomerAddressEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntity CustomerAddressEntity
	if err := g.Where("id = ?", id).First(&customerAddressEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntity)
	}
}

func createCustomerAddressEntity(c *gin.Context) {
	var customerAddressEntity CustomerAddressEntity
	c.BindJSON(&customerAddressEntity)
	g.Create(&customerAddressEntity)
	c.JSON(200, customerAddressEntity)
}

func updateCustomerAddressEntity(c *gin.Context) {
	var customerAddressEntity CustomerAddressEntity
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerAddressEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerAddressEntity)
	g.Save(&customerAddressEntity)
	c.JSON(200, customerAddressEntity)
}
func deleteCustomerAddressEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntity CustomerAddressEntity
	d := g.Where("id = ?", id).Delete(&customerAddressEntity)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
