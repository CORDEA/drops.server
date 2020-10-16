package user

import "drops/items"

type CartItems struct {
	Items []CartItem `json:"items"`
}

type CartItem struct {
	Item          items.Item `json:"item"`
	NumberOfItems int        `json:"number_of_items"`
}
