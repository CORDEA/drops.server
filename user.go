package main

import (
	usr "drops/user"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(g *gin.Engine) {
	user := g.Group("/user")

	user.POST("/login", usr.Login)
	user.POST("/new", usr.SignUp)
	user.GET("/achievements", usr.GetAchievements)
	user.GET("/cart", usr.GetCartItems)
	user.POST("/cart", usr.AddCartItem)
	user.GET("/orders", usr.GetOrders)
	user.PATCH("/orders/cancel/:id", usr.CancelOrder)
}
