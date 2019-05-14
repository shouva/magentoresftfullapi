package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogEavAttribute :
type CatalogEavAttribute struct {
	AttributeId               uint16 `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
	FrontendInputRenderer     string `gorm:"column:frontend_input_renderer" form:"frontend_input_renderer" json:"frontend_input_renderer"`
	IsGlobal                  uint16 `gorm:"column:is_global" form:"is_global" json:"is_global"`
	IsVisible                 uint16 `gorm:"column:is_visible" form:"is_visible" json:"is_visible"`
	IsSearchable              uint16 `gorm:"column:is_searchable" form:"is_searchable" json:"is_searchable"`
	IsFilterable              uint16 `gorm:"column:is_filterable" form:"is_filterable" json:"is_filterable"`
	IsComparable              uint16 `gorm:"column:is_comparable" form:"is_comparable" json:"is_comparable"`
	IsVisibleOnFront          uint16 `gorm:"column:is_visible_on_front" form:"is_visible_on_front" json:"is_visible_on_front"`
	IsHtmlAllowedOnFront      uint16 `gorm:"column:is_html_allowed_on_front" form:"is_html_allowed_on_front" json:"is_html_allowed_on_front"`
	IsUsedForPriceRules       uint16 `gorm:"column:is_used_for_price_rules" form:"is_used_for_price_rules" json:"is_used_for_price_rules"`
	IsFilterableInSearch      uint16 `gorm:"column:is_filterable_in_search" form:"is_filterable_in_search" json:"is_filterable_in_search"`
	UsedInProductListing      uint16 `gorm:"column:used_in_product_listing" form:"used_in_product_listing" json:"used_in_product_listing"`
	UsedForSortBy             uint16 `gorm:"column:used_for_sort_by" form:"used_for_sort_by" json:"used_for_sort_by"`
	IsConfigurable            uint16 `gorm:"column:is_configurable" form:"is_configurable" json:"is_configurable"`
	ApplyTo                   string `gorm:"column:apply_to" form:"apply_to" json:"apply_to"`
	IsVisibleInAdvancedSearch uint16 `gorm:"column:is_visible_in_advanced_search" form:"is_visible_in_advanced_search" json:"is_visible_in_advanced_search"`
	Position                  uint16 `gorm:"column:position" form:"position" json:"position"`
	IsWysiwygEnabled          uint16 `gorm:"column:is_wysiwyg_enabled" form:"is_wysiwyg_enabled" json:"is_wysiwyg_enabled"`
	IsUsedForPromoRules       uint16 `gorm:"column:is_used_for_promo_rules" form:"is_used_for_promo_rules" json:"is_used_for_promo_rules"`
}

// TableName :
func (*CatalogEavAttribute) TableName() string {
	return "catalog_eav_attribute"
}

// handler create
func initRoutersCatalogEavAttribute(r *gin.Engine, catalogeavattribute string) {
	route := r.Group("catalogeavattribute")
	route.GET("/", getCatalogEavAttributes)
	route.GET("/:id", getCatalogEavAttribute)
	route.POST("/", createCatalogEavAttribute)
	route.PUT("/:id", updateCatalogEavAttribute)
	route.DELETE("/:id", deleteCatalogEavAttribute)
}

func getCatalogEavAttributes(c *gin.Context) {
	var catalogEavAttributes []CatalogEavAttribute
	if err := g.Find(&catalogEavAttributes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogEavAttributes)
	}
}

func getCatalogEavAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogEavAttribute CatalogEavAttribute
	if err := g.Where("id = ?", id).First(&catalogEavAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogEavAttribute)
	}
}

func createCatalogEavAttribute(c *gin.Context) {
	var catalogEavAttribute CatalogEavAttribute
	c.BindJSON(&catalogEavAttribute)
	g.Create(&catalogEavAttribute)
	c.JSON(200, catalogEavAttribute)
}

func updateCatalogEavAttribute(c *gin.Context) {
	var catalogEavAttribute CatalogEavAttribute
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogEavAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogEavAttribute)
	g.Save(&catalogEavAttribute)
	c.JSON(200, catalogEavAttribute)
}
func deleteCatalogEavAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogEavAttribute CatalogEavAttribute
	d := g.Where("id = ?", id).Delete(&catalogEavAttribute)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
