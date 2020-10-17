package user

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(g *gin.Engine) {
	user := g.Group("/user")

	user.POST("/login", login)
	user.POST("/new", signUp)
	user.GET("/achievements", getAchievements)
	user.GET("/cart", getCartItems)
	user.POST("/cart", addCartItem)
	user.GET("/orders", getOrders)
	user.PATCH("/orders/cancel/:id", cancelOrder)
}
