package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type body struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

func (r *Route) login(ctx *gin.Context) {
	var b body
	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusOK, "")
}

func (r *Route) signUp(ctx *gin.Context) {
	var b body
	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusOK, "")
}
