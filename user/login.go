package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type body struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func Handle(ctx *gin.Context) {
	var b body
	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO
	ctx.JSON(http.StatusOK, User{Id: "id", Name: "name", CreatedAt: time.Time{}})
}
