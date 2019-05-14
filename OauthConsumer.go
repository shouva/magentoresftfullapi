package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// OauthConsumer :
type OauthConsumer struct {
	EntityId            uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CreatedAt           *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	Name                string     `gorm:"column:name" form:"name" json:"name"`
	Key                 string     `gorm:"column:key" form:"key" json:"key"`
	Secret              string     `gorm:"column:secret" form:"secret" json:"secret"`
	CallbackUrl         string     `gorm:"column:callback_url" form:"callback_url" json:"callback_url"`
	RejectedCallbackUrl string     `gorm:"column:rejected_callback_url" form:"rejected_callback_url" json:"rejected_callback_url"`
}

// TableName :
func (*OauthConsumer) TableName() string {
	return "oauth_consumer"
}

// handler create
func initRoutersOauthConsumer(r *gin.Engine, oauthconsumer string) {
	route := r.Group("oauthconsumer")
	route.GET("/", getOauthConsumers)
	route.GET("/:id", getOauthConsumer)
	route.POST("/", createOauthConsumer)
	route.PUT("/:id", updateOauthConsumer)
	route.DELETE("/:id", deleteOauthConsumer)
}

func getOauthConsumers(c *gin.Context) {
	var oauthConsumers []OauthConsumer
	if err := g.Find(&oauthConsumers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, oauthConsumers)
	}
}

func getOauthConsumer(c *gin.Context) {
	id := c.Params.ByName("id")
	var oauthConsumer OauthConsumer
	if err := g.Where("id = ?", id).First(&oauthConsumer).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, oauthConsumer)
	}
}

func createOauthConsumer(c *gin.Context) {
	var oauthConsumer OauthConsumer
	c.BindJSON(&oauthConsumer)
	g.Create(&oauthConsumer)
	c.JSON(200, oauthConsumer)
}

func updateOauthConsumer(c *gin.Context) {
	var oauthConsumer OauthConsumer
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&oauthConsumer).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&oauthConsumer)
	g.Save(&oauthConsumer)
	c.JSON(200, oauthConsumer)
}
func deleteOauthConsumer(c *gin.Context) {
	id := c.Params.ByName("id")
	var oauthConsumer OauthConsumer
	d := g.Where("id = ?", id).Delete(&oauthConsumer)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
