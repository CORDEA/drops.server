package user

import (
	"time"
)

type body struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
