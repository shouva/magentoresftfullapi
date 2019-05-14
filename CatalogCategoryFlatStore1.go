package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryFlatStore1 :
type CatalogCategoryFlatStore1 struct {
	EntityId                uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ParentId                uint16     `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	CreatedAt               *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt               *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	Path                    string     `gorm:"column:path" form:"path" json:"path"`
	Position                uint16     `gorm:"column:position" form:"position" json:"position"`
	Level                   uint16     `gorm:"column:level" form:"level" json:"level"`
	ChildrenCount           uint16     `gorm:"column:children_count" form:"children_count" json:"children_count"`
	StoreId                 uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	AllChildren             string     `gorm:"column:all_children" form:"all_children" json:"all_children"`
	AvailableSortBy         string     `gorm:"column:available_sort_by" form:"available_sort_by" json:"available_sort_by"`
	Children                string     `gorm:"column:children" form:"children" json:"children"`
	CustomApplyToProducts   uint16     `gorm:"column:custom_apply_to_products" form:"custom_apply_to_products" json:"custom_apply_to_products"`
	CustomDesign            string     `gorm:"column:custom_design" form:"custom_design" json:"custom_design"`
	CustomDesignFrom        *time.Time `gorm:"column:custom_design_from" form:"custom_design_from" json:"custom_design_from"`
	CustomDesignTo          *time.Time `gorm:"column:custom_design_to" form:"custom_design_to" json:"custom_design_to"`
	CustomLayoutUpdate      string     `gorm:"column:custom_layout_update" form:"custom_layout_update" json:"custom_layout_update"`
	CustomUseParentSettings uint16     `gorm:"column:custom_use_parent_settings" form:"custom_use_parent_settings" json:"custom_use_parent_settings"`
	DefaultSortBy           string     `gorm:"column:default_sort_by" form:"default_sort_by" json:"default_sort_by"`
	Description             string     `gorm:"column:description" form:"description" json:"description"`
	DisplayMode             string     `gorm:"column:display_mode" form:"display_mode" json:"display_mode"`
	FilterPriceRange        float64    `gorm:"column:filter_price_range" form:"filter_price_range" json:"filter_price_range"`
	Image                   string     `gorm:"column:image" form:"image" json:"image"`
	IncludeInMenu           uint16     `gorm:"column:include_in_menu" form:"include_in_menu" json:"include_in_menu"`
	IsActive                uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
	IsAnchor                uint16     `gorm:"column:is_anchor" form:"is_anchor" json:"is_anchor"`
	LandingPage             uint16     `gorm:"column:landing_page" form:"landing_page" json:"landing_page"`
	MetaDescription         string     `gorm:"column:meta_description" form:"meta_description" json:"meta_description"`
	MetaKeywords            string     `gorm:"column:meta_keywords" form:"meta_keywords" json:"meta_keywords"`
	MetaTitle               string     `gorm:"column:meta_title" form:"meta_title" json:"meta_title"`
	Name                    string     `gorm:"column:name" form:"name" json:"name"`
	PageLayout              string     `gorm:"column:page_layout" form:"page_layout" json:"page_layout"`
	PathInStore             string     `gorm:"column:path_in_store" form:"path_in_store" json:"path_in_store"`
	Thumbnail               string     `gorm:"column:thumbnail" form:"thumbnail" json:"thumbnail"`
	UrlKey                  string     `gorm:"column:url_key" form:"url_key" json:"url_key"`
	UrlPath                 string     `gorm:"column:url_path" form:"url_path" json:"url_path"`
}

// TableName :
func (*CatalogCategoryFlatStore1) TableName() string {
	return "catalog_category_flat_store_1"
}

// handler create
func initRoutersCatalogCategoryFlatStore1(r *gin.Engine, catalogcategoryflatstore1 string) {
	route := r.Group("catalogcategoryflatstore1")
	route.GET("/", getCatalogCategoryFlatStore1s)
	route.GET("/:id", getCatalogCategoryFlatStore1)
	route.POST("/", createCatalogCategoryFlatStore1)
	route.PUT("/:id", updateCatalogCategoryFlatStore1)
	route.DELETE("/:id", deleteCatalogCategoryFlatStore1)
}

func getCatalogCategoryFlatStore1s(c *gin.Context) {
	var catalogCategoryFlatStore1s []CatalogCategoryFlatStore1
	if err := g.Find(&catalogCategoryFlatStore1s).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryFlatStore1s)
	}
}

func getCatalogCategoryFlatStore1(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryFlatStore1 CatalogCategoryFlatStore1
	if err := g.Where("id = ?", id).First(&catalogCategoryFlatStore1).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryFlatStore1)
	}
}

func createCatalogCategoryFlatStore1(c *gin.Context) {
	var catalogCategoryFlatStore1 CatalogCategoryFlatStore1
	c.BindJSON(&catalogCategoryFlatStore1)
	g.Create(&catalogCategoryFlatStore1)
	c.JSON(200, catalogCategoryFlatStore1)
}

func updateCatalogCategoryFlatStore1(c *gin.Context) {
	var catalogCategoryFlatStore1 CatalogCategoryFlatStore1
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryFlatStore1).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryFlatStore1)
	g.Save(&catalogCategoryFlatStore1)
	c.JSON(200, catalogCategoryFlatStore1)
}
func deleteCatalogCategoryFlatStore1(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryFlatStore1 CatalogCategoryFlatStore1
	d := g.Where("id = ?", id).Delete(&catalogCategoryFlatStore1)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
