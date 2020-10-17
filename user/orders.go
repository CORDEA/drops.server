package user

import "drops/items"

type Status int

const (
	CONFIRMED = iota + 1
	SHIPPED
	DELIVERED
	CANCELED
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
