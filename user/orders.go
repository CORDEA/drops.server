package user

import (
	"github.com/CORDEA/drops.server/etsy/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Status int

const (
	Confirmed Status = iota + 1
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
	Status       string       `json:"status"`
	Items        []model.Item `json:"items"`
	IsCancelable bool         `json:"is_cancelable"`
	OrderedAt    *time.Time   `json:"ordered_at"`
	ShippedAt    *time.Time   `json:"shipped_at"`
	DeliveredAt  *time.Time   `json:"delivered_at"`
	CanceledAt   *time.Time   `json:"canceled_at"`
}

func (r *Route) getOrders(ctx *gin.Context) {
	earrings, err := r.getEarrings.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var orders []Order
	id, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	orders = append(orders, Order{
		Id:           id.String(),
		Status:       Confirmed.String(),
		Items:        earrings.Items[0:3],
		IsCancelable: true,
		OrderedAt:    &now,
		ShippedAt:    nil,
		DeliveredAt:  nil,
		CanceledAt:   nil,
	})
	id, err = uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	orderedAt := now.Add(-5 * 24 * time.Hour)
	shippedAt := now.Add(-2 * 24 * time.Hour)
	orders = append(orders, Order{
		Id:           id.String(),
		Status:       Delivered.String(),
		Items:        earrings.Items[4:5],
		IsCancelable: false,
		OrderedAt:    &orderedAt,
		ShippedAt:    &shippedAt,
		DeliveredAt:  &now,
		CanceledAt:   nil,
	})
	ctx.JSON(
		http.StatusOK,
		Orders{orders},
	)
}

func (r *Route) cancelOrder(ctx *gin.Context) {
	ctx.String(http.StatusOK, "")
}
