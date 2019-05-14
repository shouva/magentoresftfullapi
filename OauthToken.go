package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// OauthToken :
type OauthToken struct {
	EntityId    uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ConsumerId  uint16     `gorm:"column:consumer_id" form:"consumer_id" json:"consumer_id"`
	AdminId     uint16     `gorm:"column:admin_id" form:"admin_id" json:"admin_id"`
	CustomerId  uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	Type        string     `gorm:"column:type" form:"type" json:"type"`
	Token       string     `gorm:"column:token" form:"token" json:"token"`
	Secret      string     `gorm:"column:secret" form:"secret" json:"secret"`
	Verifier    string     `gorm:"column:verifier" form:"verifier" json:"verifier"`
	CallbackUrl string     `gorm:"column:callback_url" form:"callback_url" json:"callback_url"`
	Revoked     uint16     `gorm:"column:revoked" form:"revoked" json:"revoked"`
	Authorized  uint16     `gorm:"column:authorized" form:"authorized" json:"authorized"`
	CreatedAt   *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}

// TableName :
func (*OauthToken) TableName() string {
	return "oauth_token"
}

// handler create
func initRoutersOauthToken(r *gin.Engine, oauthtoken string) {
	route := r.Group("oauthtoken")
	route.GET("/", getOauthTokens)
	route.GET("/:id", getOauthToken)
	route.POST("/", createOauthToken)
	route.PUT("/:id", updateOauthToken)
	route.DELETE("/:id", deleteOauthToken)
}

func getOauthTokens(c *gin.Context) {
	var oauthTokens []OauthToken
	if err := g.Find(&oauthTokens).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, oauthTokens)
	}
}

func getOauthToken(c *gin.Context) {
	id := c.Params.ByName("id")
	var oauthToken OauthToken
	if err := g.Where("id = ?", id).First(&oauthToken).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, oauthToken)
	}
}

func createOauthToken(c *gin.Context) {
	var oauthToken OauthToken
	c.BindJSON(&oauthToken)
	g.Create(&oauthToken)
	c.JSON(200, oauthToken)
}

func updateOauthToken(c *gin.Context) {
	var oauthToken OauthToken
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&oauthToken).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&oauthToken)
	g.Save(&oauthToken)
	c.JSON(200, oauthToken)
}
func deleteOauthToken(c *gin.Context) {
	id := c.Params.ByName("id")
	var oauthToken OauthToken
	d := g.Where("id = ?", id).Delete(&oauthToken)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
