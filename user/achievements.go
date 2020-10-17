package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Achievement struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Level       uint    `json:"level"`
	Progress    float32 `json:"progress"`
}

func getAchievements(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Achievement{
		Name:        "name",
		Description: "description",
		Level:       1,
		Progress:    0.1,
	})
}
