package user

import (
	"github.com/CORDEA/drops.server/etsy/usecase"
	"github.com/gin-gonic/gin"
)

type Route struct {
	getEarrings *usecase.GetEarrings
}

func NewRoute(getEarrings *usecase.GetEarrings) *Route {
	return &Route{getEarrings}
}

func (r *Route) Apply(g *gin.Engine) {
	user := g.Group("/user")

	user.GET("/", r.getUser)
	user.POST("/login", r.login)
	user.POST("/new", r.signUp)
	user.GET("/achievements", r.getAchievements)
	user.GET("/cart", r.getCartItems)
	user.POST("/cart", r.addCartItem)
	user.GET("/orders", r.getOrders)
	user.PATCH("/orders/cancel/:id", r.cancelOrder)
}
