package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesFlatOrderAddres :
type SalesFlatOrderAddres struct {
	EntityId          uint16 `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ParentId          uint16 `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	CustomerAddressId uint16 `gorm:"column:customer_address_id" form:"customer_address_id" json:"customer_address_id"`
	QuoteAddressId    uint16 `gorm:"column:quote_address_id" form:"quote_address_id" json:"quote_address_id"`
	RegionId          uint16 `gorm:"column:region_id" form:"region_id" json:"region_id"`
	CustomerId        uint16 `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	Fax               string `gorm:"column:fax" form:"fax" json:"fax"`
	Region            string `gorm:"column:region" form:"region" json:"region"`
	Postcode          string `gorm:"column:postcode" form:"postcode" json:"postcode"`
	Lastname          string `gorm:"column:lastname" form:"lastname" json:"lastname"`
	Street            string `gorm:"column:street" form:"street" json:"street"`
	City              string `gorm:"column:city" form:"city" json:"city"`
	Email             string `gorm:"column:email" form:"email" json:"email"`
	Telephone         string `gorm:"column:telephone" form:"telephone" json:"telephone"`
	CountryId         string `gorm:"column:country_id" form:"country_id" json:"country_id"`
	Firstname         string `gorm:"column:firstname" form:"firstname" json:"firstname"`
	AddressType       string `gorm:"column:address_type" form:"address_type" json:"address_type"`
	Prefix            string `gorm:"column:prefix" form:"prefix" json:"prefix"`
	Middlename        string `gorm:"column:middlename" form:"middlename" json:"middlename"`
	Suffix            string `gorm:"column:suffix" form:"suffix" json:"suffix"`
	Company           string `gorm:"column:company" form:"company" json:"company"`
	VatId             string `gorm:"column:vat_id" form:"vat_id" json:"vat_id"`
	VatIsValid        uint16 `gorm:"column:vat_is_valid" form:"vat_is_valid" json:"vat_is_valid"`
	VatRequestId      string `gorm:"column:vat_request_id" form:"vat_request_id" json:"vat_request_id"`
	VatRequestDate    string `gorm:"column:vat_request_date" form:"vat_request_date" json:"vat_request_date"`
	VatRequestSuccess uint16 `gorm:"column:vat_request_success" form:"vat_request_success" json:"vat_request_success"`
}

// TableName :
func (*SalesFlatOrderAddres) TableName() string {
	return "sales_flat_order_address"
}

// handler create
func initRoutersSalesFlatOrderAddres(r *gin.Engine, salesflatorderaddres string) {
	route := r.Group("salesflatorderaddres")
	route.GET("/", getSalesFlatOrderAddress)
	route.GET("/:id", getSalesFlatOrderAddres)
	route.POST("/", createSalesFlatOrderAddres)
	route.PUT("/:id", updateSalesFlatOrderAddres)
	route.DELETE("/:id", deleteSalesFlatOrderAddres)
}

func getSalesFlatOrderAddress(c *gin.Context) {
	var salesFlatOrderAddress []SalesFlatOrderAddres
	if err := g.Find(&salesFlatOrderAddress).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrderAddress)
	}
}

func getSalesFlatOrderAddres(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrderAddres SalesFlatOrderAddres
	if err := g.Where("id = ?", id).First(&salesFlatOrderAddres).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrderAddres)
	}
}

func createSalesFlatOrderAddres(c *gin.Context) {
	var salesFlatOrderAddres SalesFlatOrderAddres
	c.BindJSON(&salesFlatOrderAddres)
	g.Create(&salesFlatOrderAddres)
	c.JSON(200, salesFlatOrderAddres)
}

func updateSalesFlatOrderAddres(c *gin.Context) {
	var salesFlatOrderAddres SalesFlatOrderAddres
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatOrderAddres).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatOrderAddres)
	g.Save(&salesFlatOrderAddres)
	c.JSON(200, salesFlatOrderAddres)
}
func deleteSalesFlatOrderAddres(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrderAddres SalesFlatOrderAddres
	d := g.Where("id = ?", id).Delete(&salesFlatOrderAddres)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
