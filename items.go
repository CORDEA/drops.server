package main

import "github.com/gin-gonic/gin"

func addItemRoutes(g *gin.Engine) {
	items := g.Group("/items")

	items.GET("/", func(context *gin.Context) {
	})
	items.GET("/:id", func(context *gin.Context) {
	})
}
