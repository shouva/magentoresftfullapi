package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Api2AclAttribute :
type Api2AclAttribute struct {
	EntityId          uint16 `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	UserType          string `gorm:"column:user_type" form:"user_type" json:"user_type"`
	ResourceId        string `gorm:"column:resource_id" form:"resource_id" json:"resource_id"`
	Operation         string `gorm:"column:operation" form:"operation" json:"operation"`
	AllowedAttributes string `gorm:"column:allowed_attributes" form:"allowed_attributes" json:"allowed_attributes"`
}

// TableName :
func (*Api2AclAttribute) TableName() string {
	return "api2_acl_attribute"
}

// handler create
func initRoutersApi2AclAttribute(r *gin.Engine, api2aclattribute string) {
	route := r.Group("api2aclattribute")
	route.GET("/", getApi2AclAttributes)
	route.GET("/:id", getApi2AclAttribute)
	route.POST("/", createApi2AclAttribute)
	route.PUT("/:id", updateApi2AclAttribute)
	route.DELETE("/:id", deleteApi2AclAttribute)
}

func getApi2AclAttributes(c *gin.Context) {
	var api2AclAttributes []Api2AclAttribute
	if err := g.Find(&api2AclAttributes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, api2AclAttributes)
	}
}

func getApi2AclAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var api2AclAttribute Api2AclAttribute
	if err := g.Where("id = ?", id).First(&api2AclAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, api2AclAttribute)
	}
}

func createApi2AclAttribute(c *gin.Context) {
	var api2AclAttribute Api2AclAttribute
	c.BindJSON(&api2AclAttribute)
	g.Create(&api2AclAttribute)
	c.JSON(200, api2AclAttribute)
}

func updateApi2AclAttribute(c *gin.Context) {
	var api2AclAttribute Api2AclAttribute
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&api2AclAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&api2AclAttribute)
	g.Save(&api2AclAttribute)
	c.JSON(200, api2AclAttribute)
}
func deleteApi2AclAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var api2AclAttribute Api2AclAttribute
	d := g.Where("id = ?", id).Delete(&api2AclAttribute)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
