package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// OauthNonce :
type OauthNonce struct {
	Nonce     string `gorm:"column:nonce;primary_key" form:"nonce;primary_key" json:"nonce;primary_key"`
	Timestamp uint16 `gorm:"column:timestamp" form:"timestamp" json:"timestamp"`
}

// TableName :
func (*OauthNonce) TableName() string {
	return "oauth_nonce"
}

// handler create
func initRoutersOauthNonce(r *gin.Engine, oauthnonce string) {
	route := r.Group("oauthnonce")
	route.GET("/", getOauthNonces)
	route.GET("/:id", getOauthNonce)
	route.POST("/", createOauthNonce)
	route.PUT("/:id", updateOauthNonce)
	route.DELETE("/:id", deleteOauthNonce)
}

func getOauthNonces(c *gin.Context) {
	var oauthNonces []OauthNonce
	if err := g.Find(&oauthNonces).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, oauthNonces)
	}
}

func getOauthNonce(c *gin.Context) {
	id := c.Params.ByName("id")
	var oauthNonce OauthNonce
	if err := g.Where("id = ?", id).First(&oauthNonce).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, oauthNonce)
	}
}

func createOauthNonce(c *gin.Context) {
	var oauthNonce OauthNonce
	c.BindJSON(&oauthNonce)
	g.Create(&oauthNonce)
	c.JSON(200, oauthNonce)
}

func updateOauthNonce(c *gin.Context) {
	var oauthNonce OauthNonce
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&oauthNonce).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&oauthNonce)
	g.Save(&oauthNonce)
	c.JSON(200, oauthNonce)
}
func deleteOauthNonce(c *gin.Context) {
	id := c.Params.ByName("id")
	var oauthNonce OauthNonce
	d := g.Where("id = ?", id).Delete(&oauthNonce)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
