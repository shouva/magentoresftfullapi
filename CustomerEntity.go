package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CustomerEntity :
type CustomerEntity struct {
	EntityId               uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	EntityTypeId           uint16     `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeSetId         uint16     `gorm:"column:attribute_set_id" form:"attribute_set_id" json:"attribute_set_id"`
	WebsiteId              uint16     `gorm:"column:website_id" form:"website_id" json:"website_id"`
	Email                  string     `gorm:"column:email" form:"email" json:"email"`
	GroupId                uint16     `gorm:"column:group_id" form:"group_id" json:"group_id"`
	IncrementId            string     `gorm:"column:increment_id" form:"increment_id" json:"increment_id"`
	StoreId                uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	CreatedAt              *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt              *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	IsActive               uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
	DisableAutoGroupChange uint16     `gorm:"column:disable_auto_group_change" form:"disable_auto_group_change" json:"disable_auto_group_change"`
}

// TableName :
func (*CustomerEntity) TableName() string {
	return "customer_entity"
}

// handler create
func initRoutersCustomerEntity(r *gin.Engine, customerentity string) {
	route := r.Group("customerentity")
	route.GET("/", getCustomerEntitys)
	route.GET("/:id", getCustomerEntity)
	route.POST("/", createCustomerEntity)
	route.PUT("/:id", updateCustomerEntity)
	route.DELETE("/:id", deleteCustomerEntity)
}

func getCustomerEntitys(c *gin.Context) {
	var customerEntitys []CustomerEntity
	if err := g.Find(&customerEntitys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntitys)
	}
}

func getCustomerEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntity CustomerEntity
	if err := g.Where("id = ?", id).First(&customerEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntity)
	}
}

func createCustomerEntity(c *gin.Context) {
	var customerEntity CustomerEntity
	c.BindJSON(&customerEntity)
	g.Create(&customerEntity)
	c.JSON(200, customerEntity)
}

func updateCustomerEntity(c *gin.Context) {
	var customerEntity CustomerEntity
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerEntity)
	g.Save(&customerEntity)
	c.JSON(200, customerEntity)
}
func deleteCustomerEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntity CustomerEntity
	d := g.Where("id = ?", id).Delete(&customerEntity)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
