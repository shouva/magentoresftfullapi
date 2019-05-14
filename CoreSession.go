package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreSession :
type CoreSession struct {
	SessionId      string `gorm:"column:session_id;primary_key" form:"session_id;primary_key" json:"session_id;primary_key"`
	SessionExpires uint16 `gorm:"column:session_expires" form:"session_expires" json:"session_expires"`
	SessionData    []byte `gorm:"column:session_data" form:"session_data" json:"session_data"`
}

// TableName :
func (*CoreSession) TableName() string {
	return "core_session"
}

// handler create
func initRoutersCoreSession(r *gin.Engine, coresession string) {
	route := r.Group("coresession")
	route.GET("/", getCoreSessions)
	route.GET("/:id", getCoreSession)
	route.POST("/", createCoreSession)
	route.PUT("/:id", updateCoreSession)
	route.DELETE("/:id", deleteCoreSession)
}

func getCoreSessions(c *gin.Context) {
	var coreSessions []CoreSession
	if err := g.Find(&coreSessions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreSessions)
	}
}

func getCoreSession(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreSession CoreSession
	if err := g.Where("id = ?", id).First(&coreSession).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreSession)
	}
}

func createCoreSession(c *gin.Context) {
	var coreSession CoreSession
	c.BindJSON(&coreSession)
	g.Create(&coreSession)
	c.JSON(200, coreSession)
}

func updateCoreSession(c *gin.Context) {
	var coreSession CoreSession
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreSession).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreSession)
	g.Save(&coreSession)
	c.JSON(200, coreSession)
}
func deleteCoreSession(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreSession CoreSession
	d := g.Where("id = ?", id).Delete(&coreSession)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
