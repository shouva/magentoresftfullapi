package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerGroup :
type CustomerGroup struct {
	CustomerGroupId   uint16 `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	CustomerGroupCode string `gorm:"column:customer_group_code" form:"customer_group_code" json:"customer_group_code"`
	TaxClassId        uint16 `gorm:"column:tax_class_id" form:"tax_class_id" json:"tax_class_id"`
}

// TableName :
func (*CustomerGroup) TableName() string {
	return "customer_group"
}

// handler create
func initRoutersCustomerGroup(r *gin.Engine, customergroup string) {
	route := r.Group("customergroup")
	route.GET("/", getCustomerGroups)
	route.GET("/:id", getCustomerGroup)
	route.POST("/", createCustomerGroup)
	route.PUT("/:id", updateCustomerGroup)
	route.DELETE("/:id", deleteCustomerGroup)
}

func getCustomerGroups(c *gin.Context) {
	var customerGroups []CustomerGroup
	if err := g.Find(&customerGroups).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerGroups)
	}
}

func getCustomerGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerGroup CustomerGroup
	if err := g.Where("id = ?", id).First(&customerGroup).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerGroup)
	}
}

func createCustomerGroup(c *gin.Context) {
	var customerGroup CustomerGroup
	c.BindJSON(&customerGroup)
	g.Create(&customerGroup)
	c.JSON(200, customerGroup)
}

func updateCustomerGroup(c *gin.Context) {
	var customerGroup CustomerGroup
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerGroup).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerGroup)
	g.Save(&customerGroup)
	c.JSON(200, customerGroup)
}
func deleteCustomerGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerGroup CustomerGroup
	d := g.Where("id = ?", id).Delete(&customerGroup)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
