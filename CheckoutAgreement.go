package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CheckoutAgreement :
type CheckoutAgreement struct {
	AgreementId   uint16 `gorm:"column:agreement_id;primary_key" form:"agreement_id;primary_key" json:"agreement_id;primary_key"`
	Name          string `gorm:"column:name" form:"name" json:"name"`
	Content       string `gorm:"column:content" form:"content" json:"content"`
	ContentHeight string `gorm:"column:content_height" form:"content_height" json:"content_height"`
	CheckboxText  string `gorm:"column:checkbox_text" form:"checkbox_text" json:"checkbox_text"`
	IsActive      uint16 `gorm:"column:is_active" form:"is_active" json:"is_active"`
	IsHtml        uint16 `gorm:"column:is_html" form:"is_html" json:"is_html"`
}

// TableName :
func (*CheckoutAgreement) TableName() string {
	return "checkout_agreement"
}

// handler create
func initRoutersCheckoutAgreement(r *gin.Engine, checkoutagreement string) {
	route := r.Group("checkoutagreement")
	route.GET("/", getCheckoutAgreements)
	route.GET("/:id", getCheckoutAgreement)
	route.POST("/", createCheckoutAgreement)
	route.PUT("/:id", updateCheckoutAgreement)
	route.DELETE("/:id", deleteCheckoutAgreement)
}

func getCheckoutAgreements(c *gin.Context) {
	var checkoutAgreements []CheckoutAgreement
	if err := g.Find(&checkoutAgreements).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, checkoutAgreements)
	}
}

func getCheckoutAgreement(c *gin.Context) {
	id := c.Params.ByName("id")
	var checkoutAgreement CheckoutAgreement
	if err := g.Where("id = ?", id).First(&checkoutAgreement).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, checkoutAgreement)
	}
}

func createCheckoutAgreement(c *gin.Context) {
	var checkoutAgreement CheckoutAgreement
	c.BindJSON(&checkoutAgreement)
	g.Create(&checkoutAgreement)
	c.JSON(200, checkoutAgreement)
}

func updateCheckoutAgreement(c *gin.Context) {
	var checkoutAgreement CheckoutAgreement
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&checkoutAgreement).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&checkoutAgreement)
	g.Save(&checkoutAgreement)
	c.JSON(200, checkoutAgreement)
}
func deleteCheckoutAgreement(c *gin.Context) {
	id := c.Params.ByName("id")
	var checkoutAgreement CheckoutAgreement
	d := g.Where("id = ?", id).Delete(&checkoutAgreement)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
