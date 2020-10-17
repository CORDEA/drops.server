package items

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(g *gin.Engine) {
	items := g.Group("/items")

	items.GET("/", getItems)
	items.GET("/:id", getItemDetails)
}
