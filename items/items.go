package items

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageUrl string `json:"image_url"`
}

func getItems(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		Items{[]Item{
			{
				Id:       "id",
				Name:     "name",
				Price:    100,
				ImageUrl: "url",
			},
		}},
	)
}

func getItemDetails(ctx *gin.Context) {
	// TODO
	ctx.JSON(
		http.StatusOK,
		Items{[]Item{
			{
				Id:       "id",
				Name:     "name",
				Price:    100,
				ImageUrl: "url",
			},
		}},
	)
}
