package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// AdminUser :
type AdminUser struct {
	UserId           uint16     `gorm:"column:user_id;primary_key" form:"user_id;primary_key" json:"user_id;primary_key"`
	Firstname        string     `gorm:"column:firstname" form:"firstname" json:"firstname"`
	Lastname         string     `gorm:"column:lastname" form:"lastname" json:"lastname"`
	Email            string     `gorm:"column:email" form:"email" json:"email"`
	Username         string     `gorm:"column:username" form:"username" json:"username"`
	Password         string     `gorm:"column:password" form:"password" json:"password"`
	Created          *time.Time `gorm:"column:created" form:"created" json:"created"`
	Modified         *time.Time `gorm:"column:modified" form:"modified" json:"modified"`
	Logdate          *time.Time `gorm:"column:logdate" form:"logdate" json:"logdate"`
	Lognum           uint16     `gorm:"column:lognum" form:"lognum" json:"lognum"`
	ReloadAclFlag    uint16     `gorm:"column:reload_acl_flag" form:"reload_acl_flag" json:"reload_acl_flag"`
	IsActive         uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
	Extra            string     `gorm:"column:extra" form:"extra" json:"extra"`
	RpToken          string     `gorm:"column:rp_token" form:"rp_token" json:"rp_token"`
	RpTokenCreatedAt *time.Time `gorm:"column:rp_token_created_at" form:"rp_token_created_at" json:"rp_token_created_at"`
}

// TableName :
func (*AdminUser) TableName() string {
	return "admin_user"
}

// handler create
func initRoutersAdminUser(r *gin.Engine, adminuser string) {
	route := r.Group("adminuser")
	route.GET("/", getAdminUsers)
	route.GET("/:id", getAdminUser)
	route.POST("/", createAdminUser)
	route.PUT("/:id", updateAdminUser)
	route.DELETE("/:id", deleteAdminUser)
}

func getAdminUsers(c *gin.Context) {
	var adminUsers []AdminUser
	if err := g.Find(&adminUsers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, adminUsers)
	}
}

func getAdminUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var adminUser AdminUser
	if err := g.Where("id = ?", id).First(&adminUser).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, adminUser)
	}
}

func createAdminUser(c *gin.Context) {
	var adminUser AdminUser
	c.BindJSON(&adminUser)
	g.Create(&adminUser)
	c.JSON(200, adminUser)
}

func updateAdminUser(c *gin.Context) {
	var adminUser AdminUser
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&adminUser).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&adminUser)
	g.Save(&adminUser)
	c.JSON(200, adminUser)
}
func deleteAdminUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var adminUser AdminUser
	d := g.Where("id = ?", id).Delete(&adminUser)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
