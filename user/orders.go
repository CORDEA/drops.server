package user

import (
	"github.com/CORDEA/drops.server/etsy/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var orders []Order
	id, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	orders = append(orders, Order{
		Id:           id.String(),
		Status:       Confirmed,
		Items:        earrings.Items[0:3],
		IsCancelable: true,
	})
	id, err = uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	orders = append(orders, Order{
		Id:           id.String(),
		Status:       Delivered,
		Items:        earrings.Items[4:5],
		IsCancelable: false,
	})
	ctx.JSON(
		http.StatusOK,
		Orders{orders},
	)
}

func (r *Route) cancelOrder(ctx *gin.Context) {
	ctx.String(http.StatusOK, "")
}
