package items

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	ImageUrls          []string `json:"image_urls"`
	CreatedAt          float32  `json:"created_at"`
	EndedAt            float32  `json:"ended_at"`
	UpdatedAt          float32  `json:"updated_at"`
	Price              string   `json:"price"`
	CurrencyCode       string   `json:"currency_code"`
	Tags               []string `json:"tags"`
	Materials          []string `json:"materials"`
	Views              int      `json:"views"`
	ItemWeight         int      `json:"item_weight"`
	ItemWeightUnit     string   `json:"item_weight_unit"`
	ItemLength         int      `json:"item_length"`
	ItemWidth          int      `json:"item_width"`
	ItemHeight         int      `json:"item_height"`
	ItemDimensionsUnit string   `json:"item_dimensions_unit"`
	Recipient          string   `json:"recipient"`
}

func (r *Route) getItems(ctx *gin.Context) {
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

func (r *Route) getItemDetails(ctx *gin.Context) {
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
