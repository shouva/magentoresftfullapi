package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CoreEmailTemplate :
type CoreEmailTemplate struct {
	TemplateId            uint16     `gorm:"column:template_id;primary_key" form:"template_id;primary_key" json:"template_id;primary_key"`
	TemplateCode          string     `gorm:"column:template_code" form:"template_code" json:"template_code"`
	TemplateText          string     `gorm:"column:template_text" form:"template_text" json:"template_text"`
	TemplateStyles        string     `gorm:"column:template_styles" form:"template_styles" json:"template_styles"`
	TemplateType          uint16     `gorm:"column:template_type" form:"template_type" json:"template_type"`
	TemplateSubject       string     `gorm:"column:template_subject" form:"template_subject" json:"template_subject"`
	TemplateSenderName    string     `gorm:"column:template_sender_name" form:"template_sender_name" json:"template_sender_name"`
	TemplateSenderEmail   string     `gorm:"column:template_sender_email" form:"template_sender_email" json:"template_sender_email"`
	AddedAt               *time.Time `gorm:"column:added_at" form:"added_at" json:"added_at"`
	ModifiedAt            *time.Time `gorm:"column:modified_at" form:"modified_at" json:"modified_at"`
	OrigTemplateCode      string     `gorm:"column:orig_template_code" form:"orig_template_code" json:"orig_template_code"`
	OrigTemplateVariables string     `gorm:"column:orig_template_variables" form:"orig_template_variables" json:"orig_template_variables"`
}

// TableName :
func (*CoreEmailTemplate) TableName() string {
	return "core_email_template"
}

// handler create
func initRoutersCoreEmailTemplate(r *gin.Engine, coreemailtemplate string) {
	route := r.Group("coreemailtemplate")
	route.GET("/", getCoreEmailTemplates)
	route.GET("/:id", getCoreEmailTemplate)
	route.POST("/", createCoreEmailTemplate)
	route.PUT("/:id", updateCoreEmailTemplate)
	route.DELETE("/:id", deleteCoreEmailTemplate)
}

func getCoreEmailTemplates(c *gin.Context) {
	var coreEmailTemplates []CoreEmailTemplate
	if err := g.Find(&coreEmailTemplates).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreEmailTemplates)
	}
}

func getCoreEmailTemplate(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreEmailTemplate CoreEmailTemplate
	if err := g.Where("id = ?", id).First(&coreEmailTemplate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreEmailTemplate)
	}
}

func createCoreEmailTemplate(c *gin.Context) {
	var coreEmailTemplate CoreEmailTemplate
	c.BindJSON(&coreEmailTemplate)
	g.Create(&coreEmailTemplate)
	c.JSON(200, coreEmailTemplate)
}

func updateCoreEmailTemplate(c *gin.Context) {
	var coreEmailTemplate CoreEmailTemplate
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreEmailTemplate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreEmailTemplate)
	g.Save(&coreEmailTemplate)
	c.JSON(200, coreEmailTemplate)
}
func deleteCoreEmailTemplate(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreEmailTemplate CoreEmailTemplate
	d := g.Where("id = ?", id).Delete(&coreEmailTemplate)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
