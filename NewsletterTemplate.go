package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// NewsletterTemplate :
type NewsletterTemplate struct {
	TemplateId               uint16     `gorm:"column:template_id;primary_key" form:"template_id;primary_key" json:"template_id;primary_key"`
	TemplateCode             string     `gorm:"column:template_code" form:"template_code" json:"template_code"`
	TemplateText             string     `gorm:"column:template_text" form:"template_text" json:"template_text"`
	TemplateTextPreprocessed string     `gorm:"column:template_text_preprocessed" form:"template_text_preprocessed" json:"template_text_preprocessed"`
	TemplateStyles           string     `gorm:"column:template_styles" form:"template_styles" json:"template_styles"`
	TemplateType             uint16     `gorm:"column:template_type" form:"template_type" json:"template_type"`
	TemplateSubject          string     `gorm:"column:template_subject" form:"template_subject" json:"template_subject"`
	TemplateSenderName       string     `gorm:"column:template_sender_name" form:"template_sender_name" json:"template_sender_name"`
	TemplateSenderEmail      string     `gorm:"column:template_sender_email" form:"template_sender_email" json:"template_sender_email"`
	TemplateActual           uint16     `gorm:"column:template_actual" form:"template_actual" json:"template_actual"`
	AddedAt                  *time.Time `gorm:"column:added_at" form:"added_at" json:"added_at"`
	ModifiedAt               *time.Time `gorm:"column:modified_at" form:"modified_at" json:"modified_at"`
}

// TableName :
func (*NewsletterTemplate) TableName() string {
	return "newsletter_template"
}

// handler create
func initRoutersNewsletterTemplate(r *gin.Engine, newslettertemplate string) {
	route := r.Group("newslettertemplate")
	route.GET("/", getNewsletterTemplates)
	route.GET("/:id", getNewsletterTemplate)
	route.POST("/", createNewsletterTemplate)
	route.PUT("/:id", updateNewsletterTemplate)
	route.DELETE("/:id", deleteNewsletterTemplate)
}

func getNewsletterTemplates(c *gin.Context) {
	var newsletterTemplates []NewsletterTemplate
	if err := g.Find(&newsletterTemplates).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterTemplates)
	}
}

func getNewsletterTemplate(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterTemplate NewsletterTemplate
	if err := g.Where("id = ?", id).First(&newsletterTemplate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterTemplate)
	}
}

func createNewsletterTemplate(c *gin.Context) {
	var newsletterTemplate NewsletterTemplate
	c.BindJSON(&newsletterTemplate)
	g.Create(&newsletterTemplate)
	c.JSON(200, newsletterTemplate)
}

func updateNewsletterTemplate(c *gin.Context) {
	var newsletterTemplate NewsletterTemplate
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&newsletterTemplate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&newsletterTemplate)
	g.Save(&newsletterTemplate)
	c.JSON(200, newsletterTemplate)
}
func deleteNewsletterTemplate(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterTemplate NewsletterTemplate
	d := g.Where("id = ?", id).Delete(&newsletterTemplate)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
