package main

import (
	usr "drops/user"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(g *gin.Engine) {
	user := g.Group("/user")

	user.POST("/login", usr.Handle)
	user.POST("/new", func(context *gin.Context) {
	})
	user.GET("/achievements", func(context *gin.Context) {
	})
	user.GET("/cart", func(context *gin.Context) {
	})
	user.POST("/cart", func(context *gin.Context) {
	})
	user.GET("/orders", func(context *gin.Context) {
	})
	user.PATCH("/orders/cancel/:id", func(context *gin.Context) {
	})
}
