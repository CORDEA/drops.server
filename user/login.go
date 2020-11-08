package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type body struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type User struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	City        string    `json:"city"`
	Rank        uint8     `json:"rank"`
	CreatedAt   time.Time `json:"created_at"`
	Birthday    time.Time `json:"birthday"`
}

func (r *Route) login(ctx *gin.Context) {
	var b body
	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, User{
		Id:          int64(id.ID()),
		Name:        "name",
		Description: "description",
		ImageUrl:    "",
		City:        "Tokyo",
		Rank:        1,
		CreatedAt:   time.Time{},
		Birthday:    time.Time{},
	})
}

func (r *Route) signUp(ctx *gin.Context) {
	var b body
	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, User{
		Id:          int64(id.ID()),
		Name:        "name",
		Description: "description",
		ImageUrl:    "",
		City:        "Tokyo",
		Rank:        1,
		CreatedAt:   time.Time{},
		Birthday:    time.Time{},
	})
}
