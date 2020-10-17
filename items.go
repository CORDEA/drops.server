package main

import (
	itm "drops/items"
	"github.com/gin-gonic/gin"
)

func addItemRoutes(g *gin.Engine) {
	items := g.Group("/items")

	items.GET("/", itm.GetItems)
	items.GET("/:id", itm.GetItemDetails)
}
