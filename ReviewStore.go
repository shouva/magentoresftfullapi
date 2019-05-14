package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ReviewStore :
type ReviewStore struct {
	ReviewId uint64 `gorm:"column:review_id;primary_key" form:"review_id;primary_key" json:"review_id;primary_key"`
	StoreId  uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
}

// TableName :
func (*ReviewStore) TableName() string {
	return "review_store"
}

// handler create
func initRoutersReviewStore(r *gin.Engine, reviewstore string) {
	route := r.Group("reviewstore")
	route.GET("/", getReviewStores)
	route.GET("/:id", getReviewStore)
	route.POST("/", createReviewStore)
	route.PUT("/:id", updateReviewStore)
	route.DELETE("/:id", deleteReviewStore)
}

func getReviewStores(c *gin.Context) {
	var reviewStores []ReviewStore
	if err := g.Find(&reviewStores).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviewStores)
	}
}

func getReviewStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var reviewStore ReviewStore
	if err := g.Where("id = ?", id).First(&reviewStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviewStore)
	}
}

func createReviewStore(c *gin.Context) {
	var reviewStore ReviewStore
	c.BindJSON(&reviewStore)
	g.Create(&reviewStore)
	c.JSON(200, reviewStore)
}

func updateReviewStore(c *gin.Context) {
	var reviewStore ReviewStore
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reviewStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reviewStore)
	g.Save(&reviewStore)
	c.JSON(200, reviewStore)
}
func deleteReviewStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var reviewStore ReviewStore
	d := g.Where("id = ?", id).Delete(&reviewStore)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
