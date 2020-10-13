package main

import "github.com/gin-gonic/gin"

func addUserRoutes(g *gin.Engine) {
	_ = g.Group("/user")
}
