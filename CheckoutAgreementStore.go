package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CheckoutAgreementStore :
type CheckoutAgreementStore struct {
	AgreementId uint16 `gorm:"column:agreement_id;primary_key" form:"agreement_id;primary_key" json:"agreement_id;primary_key"`
	StoreId     uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
}

// TableName :
func (*CheckoutAgreementStore) TableName() string {
	return "checkout_agreement_store"
}

// handler create
func initRoutersCheckoutAgreementStore(r *gin.Engine, checkoutagreementstore string) {
	route := r.Group("checkoutagreementstore")
	route.GET("/", getCheckoutAgreementStores)
	route.GET("/:id", getCheckoutAgreementStore)
	route.POST("/", createCheckoutAgreementStore)
	route.PUT("/:id", updateCheckoutAgreementStore)
	route.DELETE("/:id", deleteCheckoutAgreementStore)
}

func getCheckoutAgreementStores(c *gin.Context) {
	var checkoutAgreementStores []CheckoutAgreementStore
	if err := g.Find(&checkoutAgreementStores).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, checkoutAgreementStores)
	}
}

func getCheckoutAgreementStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var checkoutAgreementStore CheckoutAgreementStore
	if err := g.Where("id = ?", id).First(&checkoutAgreementStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, checkoutAgreementStore)
	}
}

func createCheckoutAgreementStore(c *gin.Context) {
	var checkoutAgreementStore CheckoutAgreementStore
	c.BindJSON(&checkoutAgreementStore)
	g.Create(&checkoutAgreementStore)
	c.JSON(200, checkoutAgreementStore)
}

func updateCheckoutAgreementStore(c *gin.Context) {
	var checkoutAgreementStore CheckoutAgreementStore
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&checkoutAgreementStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&checkoutAgreementStore)
	g.Save(&checkoutAgreementStore)
	c.JSON(200, checkoutAgreementStore)
}
func deleteCheckoutAgreementStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var checkoutAgreementStore CheckoutAgreementStore
	d := g.Where("id = ?", id).Delete(&checkoutAgreementStore)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
