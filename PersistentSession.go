package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// PersistentSession :
type PersistentSession struct {
	PersistentId uint16     `gorm:"column:persistent_id;primary_key" form:"persistent_id;primary_key" json:"persistent_id;primary_key"`
	Key          string     `gorm:"column:key" form:"key" json:"key"`
	CustomerId   uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	WebsiteId    uint16     `gorm:"column:website_id" form:"website_id" json:"website_id"`
	Info         string     `gorm:"column:info" form:"info" json:"info"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
}

// TableName :
func (*PersistentSession) TableName() string {
	return "persistent_session"
}

// handler create
func initRoutersPersistentSession(r *gin.Engine, persistentsession string) {
	route := r.Group("persistentsession")
	route.GET("/", getPersistentSessions)
	route.GET("/:id", getPersistentSession)
	route.POST("/", createPersistentSession)
	route.PUT("/:id", updatePersistentSession)
	route.DELETE("/:id", deletePersistentSession)
}

func getPersistentSessions(c *gin.Context) {
	var persistentSessions []PersistentSession
	if err := g.Find(&persistentSessions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, persistentSessions)
	}
}

func getPersistentSession(c *gin.Context) {
	id := c.Params.ByName("id")
	var persistentSession PersistentSession
	if err := g.Where("id = ?", id).First(&persistentSession).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, persistentSession)
	}
}

func createPersistentSession(c *gin.Context) {
	var persistentSession PersistentSession
	c.BindJSON(&persistentSession)
	g.Create(&persistentSession)
	c.JSON(200, persistentSession)
}

func updatePersistentSession(c *gin.Context) {
	var persistentSession PersistentSession
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&persistentSession).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&persistentSession)
	g.Save(&persistentSession)
	c.JSON(200, persistentSession)
}
func deletePersistentSession(c *gin.Context) {
	id := c.Params.ByName("id")
	var persistentSession PersistentSession
	d := g.Where("id = ?", id).Delete(&persistentSession)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
