package user

import (
	"github.com/CORDEA/drops.server/items"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Status int

const (
	Confirmed = iota + 1
	Shipped
	Delivered
	Canceled
)

func (s Status) String() string {
	return [...]string{
		"confirmed",
		"shipped",
		"delivered",
		"canceled",
	}[s]
}

type Orders struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	Id           string       `json:"id"`
	Status       Status       `json:"status"`
	Items        []items.Item `json:"items"`
	IsCancelable bool         `json:"is_cancelable"`
}

func getOrders(ctx *gin.Context) {
	// TODO
	ctx.JSON(
		http.StatusOK,
		Orders{[]Order{
			{
				Id:     "id",
				Status: Confirmed,
				Items: []items.Item{
					{
						Id:       "id",
						Name:     "name",
						Price:    100,
						ImageUrl: "url",
					},
				},
				IsCancelable: true,
			},
		}},
	)
}

func cancelOrder(ctx *gin.Context) {
	ctx.String(http.StatusOK, "")
}
