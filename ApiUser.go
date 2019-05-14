package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ApiUser :
type ApiUser struct {
	UserId        uint16     `gorm:"column:user_id;primary_key" form:"user_id;primary_key" json:"user_id;primary_key"`
	Firstname     string     `gorm:"column:firstname" form:"firstname" json:"firstname"`
	Lastname      string     `gorm:"column:lastname" form:"lastname" json:"lastname"`
	Email         string     `gorm:"column:email" form:"email" json:"email"`
	Username      string     `gorm:"column:username" form:"username" json:"username"`
	ApiKey        string     `gorm:"column:api_key" form:"api_key" json:"api_key"`
	Created       *time.Time `gorm:"column:created" form:"created" json:"created"`
	Modified      *time.Time `gorm:"column:modified" form:"modified" json:"modified"`
	Lognum        uint16     `gorm:"column:lognum" form:"lognum" json:"lognum"`
	ReloadAclFlag uint16     `gorm:"column:reload_acl_flag" form:"reload_acl_flag" json:"reload_acl_flag"`
	IsActive      uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
}

// TableName :
func (*ApiUser) TableName() string {
	return "api_user"
}

// handler create
func initRoutersApiUser(r *gin.Engine, apiuser string) {
	route := r.Group("apiuser")
	route.GET("/", getApiUsers)
	route.GET("/:id", getApiUser)
	route.POST("/", createApiUser)
	route.PUT("/:id", updateApiUser)
	route.DELETE("/:id", deleteApiUser)
}

func getApiUsers(c *gin.Context) {
	var apiUsers []ApiUser
	if err := g.Find(&apiUsers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, apiUsers)
	}
}

func getApiUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var apiUser ApiUser
	if err := g.Where("id = ?", id).First(&apiUser).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, apiUser)
	}
}

func createApiUser(c *gin.Context) {
	var apiUser ApiUser
	c.BindJSON(&apiUser)
	g.Create(&apiUser)
	c.JSON(200, apiUser)
}

func updateApiUser(c *gin.Context) {
	var apiUser ApiUser
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&apiUser).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&apiUser)
	g.Save(&apiUser)
	c.JSON(200, apiUser)
}
func deleteApiUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var apiUser ApiUser
	d := g.Where("id = ?", id).Delete(&apiUser)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
