package user

import (
	"github.com/CORDEA/drops.server/etsy/model"
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
	Items        []model.Item `json:"items"`
	IsCancelable bool         `json:"is_cancelable"`
}

func (r *Route) getOrders(ctx *gin.Context) {
	earrings, err := r.getEarrings.Execute()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(
		http.StatusOK,
		Orders{[]Order{
			{
				Id:           "id1",
				Status:       Confirmed,
				Items:        earrings.Items[0:3],
				IsCancelable: true,
			},
			{
				Id:           "id2",
				Status:       Delivered,
				Items:        earrings.Items[4:5],
				IsCancelable: false,
			},
		}},
	)
}

func (r *Route) cancelOrder(ctx *gin.Context) {
	ctx.String(http.StatusOK, "")
}
