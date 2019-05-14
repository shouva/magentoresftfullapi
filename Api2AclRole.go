package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Api2AclRole :
type Api2AclRole struct {
	EntityId  uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	RoleName  string     `gorm:"column:role_name" form:"role_name" json:"role_name"`
}

// TableName :
func (*Api2AclRole) TableName() string {
	return "api2_acl_role"
}

// handler create
func initRoutersApi2AclRole(r *gin.Engine, api2aclrole string) {
	route := r.Group("api2aclrole")
	route.GET("/", getApi2AclRoles)
	route.GET("/:id", getApi2AclRole)
	route.POST("/", createApi2AclRole)
	route.PUT("/:id", updateApi2AclRole)
	route.DELETE("/:id", deleteApi2AclRole)
}

func getApi2AclRoles(c *gin.Context) {
	var api2AclRoles []Api2AclRole
	if err := g.Find(&api2AclRoles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, api2AclRoles)
	}
}

func getApi2AclRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var api2AclRole Api2AclRole
	if err := g.Where("id = ?", id).First(&api2AclRole).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, api2AclRole)
	}
}

func createApi2AclRole(c *gin.Context) {
	var api2AclRole Api2AclRole
	c.BindJSON(&api2AclRole)
	g.Create(&api2AclRole)
	c.JSON(200, api2AclRole)
}

func updateApi2AclRole(c *gin.Context) {
	var api2AclRole Api2AclRole
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&api2AclRole).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&api2AclRole)
	g.Save(&api2AclRole)
	c.JSON(200, api2AclRole)
}
func deleteApi2AclRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var api2AclRole Api2AclRole
	d := g.Where("id = ?", id).Delete(&api2AclRole)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
