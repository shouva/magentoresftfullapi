package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// PollStore :
type PollStore struct {
	PollId  uint16 `gorm:"column:poll_id;primary_key" form:"poll_id;primary_key" json:"poll_id;primary_key"`
	StoreId uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
}

// TableName :
func (*PollStore) TableName() string {
	return "poll_store"
}

// handler create
func initRoutersPollStore(r *gin.Engine, pollstore string) {
	route := r.Group("pollstore")
	route.GET("/", getPollStores)
	route.GET("/:id", getPollStore)
	route.POST("/", createPollStore)
	route.PUT("/:id", updatePollStore)
	route.DELETE("/:id", deletePollStore)
}

func getPollStores(c *gin.Context) {
	var pollStores []PollStore
	if err := g.Find(&pollStores).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pollStores)
	}
}

func getPollStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var pollStore PollStore
	if err := g.Where("id = ?", id).First(&pollStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pollStore)
	}
}

func createPollStore(c *gin.Context) {
	var pollStore PollStore
	c.BindJSON(&pollStore)
	g.Create(&pollStore)
	c.JSON(200, pollStore)
}

func updatePollStore(c *gin.Context) {
	var pollStore PollStore
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&pollStore).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&pollStore)
	g.Save(&pollStore)
	c.JSON(200, pollStore)
}
func deletePollStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var pollStore PollStore
	d := g.Where("id = ?", id).Delete(&pollStore)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
