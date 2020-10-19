package items

import (
	"github.com/CORDEA/drops.server/etsy"
	"github.com/gin-gonic/gin"
)

type Route struct {
	client *etsy.Client
}

func NewRoute(client *etsy.Client) *Route {
	return &Route{client}
}

func (r *Route) Apply(g *gin.Engine) {
	items := g.Group("/items")

	items.GET("/", r.getItems)
}
