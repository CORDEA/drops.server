package items

import (
	"github.com/CORDEA/drops.server/etsy/api"
	"github.com/gin-gonic/gin"
)

type Route struct {
	client *api.Client
}

func NewRoute(client *api.Client) *Route {
	return &Route{client}
}

func (r *Route) Apply(g *gin.Engine) {
	items := g.Group("/items")

	items.GET("/", r.getItems)
}
