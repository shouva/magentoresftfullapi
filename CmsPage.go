package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CmsPage :
type CmsPage struct {
	PageId                uint16     `gorm:"column:page_id;primary_key" form:"page_id;primary_key" json:"page_id;primary_key"`
	Title                 string     `gorm:"column:title" form:"title" json:"title"`
	RootTemplate          string     `gorm:"column:root_template" form:"root_template" json:"root_template"`
	MetaKeywords          string     `gorm:"column:meta_keywords" form:"meta_keywords" json:"meta_keywords"`
	MetaDescription       string     `gorm:"column:meta_description" form:"meta_description" json:"meta_description"`
	Identifier            string     `gorm:"column:identifier" form:"identifier" json:"identifier"`
	ContentHeading        string     `gorm:"column:content_heading" form:"content_heading" json:"content_heading"`
	Content               string     `gorm:"column:content" form:"content" json:"content"`
	CreationTime          *time.Time `gorm:"column:creation_time" form:"creation_time" json:"creation_time"`
	UpdateTime            *time.Time `gorm:"column:update_time" form:"update_time" json:"update_time"`
	IsActive              uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
	SortOrder             uint16     `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	LayoutUpdateXml       string     `gorm:"column:layout_update_xml" form:"layout_update_xml" json:"layout_update_xml"`
	CustomTheme           string     `gorm:"column:custom_theme" form:"custom_theme" json:"custom_theme"`
	CustomRootTemplate    string     `gorm:"column:custom_root_template" form:"custom_root_template" json:"custom_root_template"`
	CustomLayoutUpdateXml string     `gorm:"column:custom_layout_update_xml" form:"custom_layout_update_xml" json:"custom_layout_update_xml"`
	CustomThemeFrom       *time.Time `gorm:"column:custom_theme_from" form:"custom_theme_from" json:"custom_theme_from"`
	CustomThemeTo         *time.Time `gorm:"column:custom_theme_to" form:"custom_theme_to" json:"custom_theme_to"`
}

// TableName :
func (*CmsPage) TableName() string {
	return "cms_page"
}

// handler create
func initRoutersCmsPage(r *gin.Engine, cmspage string) {
	route := r.Group("cmspage")
	route.GET("/", getCmsPages)
	route.GET("/:id", getCmsPage)
	route.POST("/", createCmsPage)
	route.PUT("/:id", updateCmsPage)
	route.DELETE("/:id", deleteCmsPage)
}

func getCmsPages(c *gin.Context) {
	var cmsPages []CmsPage
	if err := g.Find(&cmsPages).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cmsPages)
	}
}

func getCmsPage(c *gin.Context) {
	id := c.Params.ByName("id")
	var cmsPage CmsPage
	if err := g.Where("id = ?", id).First(&cmsPage).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cmsPage)
	}
}

func createCmsPage(c *gin.Context) {
	var cmsPage CmsPage
	c.BindJSON(&cmsPage)
	g.Create(&cmsPage)
	c.JSON(200, cmsPage)
}

func updateCmsPage(c *gin.Context) {
	var cmsPage CmsPage
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&cmsPage).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cmsPage)
	g.Save(&cmsPage)
	c.JSON(200, cmsPage)
}
func deleteCmsPage(c *gin.Context) {
	id := c.Params.ByName("id")
	var cmsPage CmsPage
	d := g.Where("id = ?", id).Delete(&cmsPage)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
