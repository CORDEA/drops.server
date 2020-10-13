package main

import "github.com/gin-gonic/gin"

func addItemRoutes(g *gin.Engine) {
	_ = g.Group("/items")
}
