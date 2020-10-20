package items

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
	items := g.Group("/items")

	items.GET("/", r.getItems)
}
