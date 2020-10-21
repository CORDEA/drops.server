package user

import (
	"github.com/CORDEA/drops.server/etsy/model"
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
	Item          model.Item `json:"item"`
	NumberOfItems int        `json:"number_of_items"`
}

func (r *Route) getCartItems(ctx *gin.Context) {
	earrings, err := r.getEarrings.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, CartItems{
		Items: []CartItem{
			{
				earrings.Items[0],
				1,
			},
			{
				earrings.Items[2],
				4,
			},
			{
				earrings.Items[3],
				2,
			},
			{
				earrings.Items[5],
				1,
			},
		},
	})
}

func (r *Route) addCartItem(ctx *gin.Context) {
	var b addCartItemBody
	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusOK, "")
}
