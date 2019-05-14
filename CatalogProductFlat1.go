package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CatalogProductFlat1 :
type CatalogProductFlat1 struct {
	EntityId                   uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	AttributeSetId             uint16     `gorm:"column:attribute_set_id" form:"attribute_set_id" json:"attribute_set_id"`
	TypeId                     string     `gorm:"column:type_id" form:"type_id" json:"type_id"`
	Cost                       float64    `gorm:"column:cost" form:"cost" json:"cost"`
	CreatedAt                  *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	GiftMessageAvailable       uint16     `gorm:"column:gift_message_available" form:"gift_message_available" json:"gift_message_available"`
	HasOptions                 uint16     `gorm:"column:has_options" form:"has_options" json:"has_options"`
	ImageLabel                 string     `gorm:"column:image_label" form:"image_label" json:"image_label"`
	IsRecurring                uint16     `gorm:"column:is_recurring" form:"is_recurring" json:"is_recurring"`
	LinksExist                 uint16     `gorm:"column:links_exist" form:"links_exist" json:"links_exist"`
	LinksPurchasedSeparately   uint16     `gorm:"column:links_purchased_separately" form:"links_purchased_separately" json:"links_purchased_separately"`
	LinksTitle                 string     `gorm:"column:links_title" form:"links_title" json:"links_title"`
	Msrp                       float64    `gorm:"column:msrp" form:"msrp" json:"msrp"`
	MsrpDisplayActualPriceType string     `gorm:"column:msrp_display_actual_price_type" form:"msrp_display_actual_price_type" json:"msrp_display_actual_price_type"`
	MsrpEnabled                uint16     `gorm:"column:msrp_enabled" form:"msrp_enabled" json:"msrp_enabled"`
	Name                       string     `gorm:"column:name" form:"name" json:"name"`
	NewsFromDate               *time.Time `gorm:"column:news_from_date" form:"news_from_date" json:"news_from_date"`
	NewsToDate                 *time.Time `gorm:"column:news_to_date" form:"news_to_date" json:"news_to_date"`
	Price                      float64    `gorm:"column:price" form:"price" json:"price"`
	PriceType                  uint16     `gorm:"column:price_type" form:"price_type" json:"price_type"`
	PriceView                  uint16     `gorm:"column:price_view" form:"price_view" json:"price_view"`
	RecurringProfile           string     `gorm:"column:recurring_profile" form:"recurring_profile" json:"recurring_profile"`
	RequiredOptions            uint16     `gorm:"column:required_options" form:"required_options" json:"required_options"`
	ShipmentType               uint16     `gorm:"column:shipment_type" form:"shipment_type" json:"shipment_type"`
	ShortDescription           string     `gorm:"column:short_description" form:"short_description" json:"short_description"`
	Sku                        string     `gorm:"column:sku" form:"sku" json:"sku"`
	SkuType                    uint16     `gorm:"column:sku_type" form:"sku_type" json:"sku_type"`
	SmallImage                 string     `gorm:"column:small_image" form:"small_image" json:"small_image"`
	SmallImageLabel            string     `gorm:"column:small_image_label" form:"small_image_label" json:"small_image_label"`
	SpecialFromDate            *time.Time `gorm:"column:special_from_date" form:"special_from_date" json:"special_from_date"`
	SpecialPrice               float64    `gorm:"column:special_price" form:"special_price" json:"special_price"`
	SpecialToDate              *time.Time `gorm:"column:special_to_date" form:"special_to_date" json:"special_to_date"`
	TaxClassId                 uint16     `gorm:"column:tax_class_id" form:"tax_class_id" json:"tax_class_id"`
	Thumbnail                  string     `gorm:"column:thumbnail" form:"thumbnail" json:"thumbnail"`
	ThumbnailLabel             string     `gorm:"column:thumbnail_label" form:"thumbnail_label" json:"thumbnail_label"`
	UpdatedAt                  *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	UrlKey                     string     `gorm:"column:url_key" form:"url_key" json:"url_key"`
	UrlPath                    string     `gorm:"column:url_path" form:"url_path" json:"url_path"`
	Visibility                 uint16     `gorm:"column:visibility" form:"visibility" json:"visibility"`
	Weight                     float64    `gorm:"column:weight" form:"weight" json:"weight"`
	WeightType                 uint16     `gorm:"column:weight_type" form:"weight_type" json:"weight_type"`
}

// TableName :
func (*CatalogProductFlat1) TableName() string {
	return "catalog_product_flat_1"
}

// handler create
func initRoutersCatalogProductFlat1(r *gin.Engine, catalogproductflat1 string) {
	route := r.Group("catalogproductflat1")
	route.GET("/", getCatalogProductFlat1s)
	route.GET("/:id", getCatalogProductFlat1)
	route.POST("/", createCatalogProductFlat1)
	route.PUT("/:id", updateCatalogProductFlat1)
	route.DELETE("/:id", deleteCatalogProductFlat1)
}

func getCatalogProductFlat1s(c *gin.Context) {
	var catalogProductFlat1s []CatalogProductFlat1
	if err := g.Find(&catalogProductFlat1s).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductFlat1s)
	}
}

func getCatalogProductFlat1(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductFlat1 CatalogProductFlat1
	if err := g.Where("id = ?", id).First(&catalogProductFlat1).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductFlat1)
	}
}

func createCatalogProductFlat1(c *gin.Context) {
	var catalogProductFlat1 CatalogProductFlat1
	c.BindJSON(&catalogProductFlat1)
	g.Create(&catalogProductFlat1)
	c.JSON(200, catalogProductFlat1)
}

func updateCatalogProductFlat1(c *gin.Context) {
	var catalogProductFlat1 CatalogProductFlat1
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductFlat1).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductFlat1)
	g.Save(&catalogProductFlat1)
	c.JSON(200, catalogProductFlat1)
}
func deleteCatalogProductFlat1(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductFlat1 CatalogProductFlat1
	d := g.Where("id = ?", id).Delete(&catalogProductFlat1)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
