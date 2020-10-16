package user

import (
	"drops/items"
	"github.com/gin-gonic/gin"
	"net/http"
)

type addCartItemBody struct {
	Id string `json:"id"`
}

type CartItems struct {
	Items []CartItem `json:"items"`
}

type CartItem struct {
	Item          items.Item `json:"item"`
	NumberOfItems int        `json:"number_of_items"`
}

func GetCartItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, CartItems{
		Items: []CartItem{
			{
				items.Item{
					Id:       "id",
					Name:     "name",
					Price:    100,
					ImageUrl: "url",
				},
				1,
			},
		},
	})
}

func AddCartItem(ctx *gin.Context) {
	var b addCartItemBody
	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusOK, "")
}
