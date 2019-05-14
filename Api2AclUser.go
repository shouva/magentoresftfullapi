package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Api2AclUser :
type Api2AclUser struct {
	AdminId uint16 `gorm:"column:admin_id;primary_key" form:"admin_id;primary_key" json:"admin_id;primary_key"`
	RoleId  uint16 `gorm:"column:role_id" form:"role_id" json:"role_id"`
}

// TableName :
func (*Api2AclUser) TableName() string {
	return "api2_acl_user"
}

// handler create
func initRoutersApi2AclUser(r *gin.Engine, api2acluser string) {
	route := r.Group("api2acluser")
	route.GET("/", getApi2AclUsers)
	route.GET("/:id", getApi2AclUser)
	route.POST("/", createApi2AclUser)
	route.PUT("/:id", updateApi2AclUser)
	route.DELETE("/:id", deleteApi2AclUser)
}

func getApi2AclUsers(c *gin.Context) {
	var api2AclUsers []Api2AclUser
	if err := g.Find(&api2AclUsers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, api2AclUsers)
	}
}

func getApi2AclUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var api2AclUser Api2AclUser
	if err := g.Where("id = ?", id).First(&api2AclUser).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, api2AclUser)
	}
}

func createApi2AclUser(c *gin.Context) {
	var api2AclUser Api2AclUser
	c.BindJSON(&api2AclUser)
	g.Create(&api2AclUser)
	c.JSON(200, api2AclUser)
}

func updateApi2AclUser(c *gin.Context) {
	var api2AclUser Api2AclUser
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&api2AclUser).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&api2AclUser)
	g.Save(&api2AclUser)
	c.JSON(200, api2AclUser)
}
func deleteApi2AclUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var api2AclUser Api2AclUser
	d := g.Where("id = ?", id).Delete(&api2AclUser)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
