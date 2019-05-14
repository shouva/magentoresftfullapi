package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// PaypalCert :
type PaypalCert struct {
	CertId    uint16     `gorm:"column:cert_id;primary_key" form:"cert_id;primary_key" json:"cert_id;primary_key"`
	WebsiteId uint16     `gorm:"column:website_id" form:"website_id" json:"website_id"`
	Content   string     `gorm:"column:content" form:"content" json:"content"`
	UpdatedAt *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
}

// TableName :
func (*PaypalCert) TableName() string {
	return "paypal_cert"
}

// handler create
func initRoutersPaypalCert(r *gin.Engine, paypalcert string) {
	route := r.Group("paypalcert")
	route.GET("/", getPaypalCerts)
	route.GET("/:id", getPaypalCert)
	route.POST("/", createPaypalCert)
	route.PUT("/:id", updatePaypalCert)
	route.DELETE("/:id", deletePaypalCert)
}

func getPaypalCerts(c *gin.Context) {
	var paypalCerts []PaypalCert
	if err := g.Find(&paypalCerts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, paypalCerts)
	}
}

func getPaypalCert(c *gin.Context) {
	id := c.Params.ByName("id")
	var paypalCert PaypalCert
	if err := g.Where("id = ?", id).First(&paypalCert).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, paypalCert)
	}
}

func createPaypalCert(c *gin.Context) {
	var paypalCert PaypalCert
	c.BindJSON(&paypalCert)
	g.Create(&paypalCert)
	c.JSON(200, paypalCert)
}

func updatePaypalCert(c *gin.Context) {
	var paypalCert PaypalCert
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&paypalCert).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&paypalCert)
	g.Save(&paypalCert)
	c.JSON(200, paypalCert)
}
func deletePaypalCert(c *gin.Context) {
	id := c.Params.ByName("id")
	var paypalCert PaypalCert
	d := g.Where("id = ?", id).Delete(&paypalCert)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
