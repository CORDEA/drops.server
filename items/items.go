package items

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Route) getItems(ctx *gin.Context) {
	earrings, err := r.getEarrings.Execute()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, earrings)
}
