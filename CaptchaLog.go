package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CaptchaLog :
type CaptchaLog struct {
	Type      string     `gorm:"column:type;primary_key" form:"type;primary_key" json:"type;primary_key"`
	Value     string     `gorm:"column:value;primary_key" form:"value;primary_key" json:"value;primary_key"`
	Count     uint16     `gorm:"column:count" form:"count" json:"count"`
	UpdatedAt *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
}

// TableName :
func (*CaptchaLog) TableName() string {
	return "captcha_log"
}

// handler create
func initRoutersCaptchaLog(r *gin.Engine, captchalog string) {
	route := r.Group("captchalog")
	route.GET("/", getCaptchaLogs)
	route.GET("/:id", getCaptchaLog)
	route.POST("/", createCaptchaLog)
	route.PUT("/:id", updateCaptchaLog)
	route.DELETE("/:id", deleteCaptchaLog)
}

func getCaptchaLogs(c *gin.Context) {
	var captchaLogs []CaptchaLog
	if err := g.Find(&captchaLogs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, captchaLogs)
	}
}

func getCaptchaLog(c *gin.Context) {
	id := c.Params.ByName("id")
	var captchaLog CaptchaLog
	if err := g.Where("id = ?", id).First(&captchaLog).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, captchaLog)
	}
}

func createCaptchaLog(c *gin.Context) {
	var captchaLog CaptchaLog
	c.BindJSON(&captchaLog)
	g.Create(&captchaLog)
	c.JSON(200, captchaLog)
}

func updateCaptchaLog(c *gin.Context) {
	var captchaLog CaptchaLog
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&captchaLog).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&captchaLog)
	g.Save(&captchaLog)
	c.JSON(200, captchaLog)
}
func deleteCaptchaLog(c *gin.Context) {
	id := c.Params.ByName("id")
	var captchaLog CaptchaLog
	d := g.Where("id = ?", id).Delete(&captchaLog)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
