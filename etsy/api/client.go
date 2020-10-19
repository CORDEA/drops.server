package api

const BaseUrl = "https://openapi.etsy.com/v2"

type Client struct {
	key string
}

func NewClient(key string) *Client {
	return &Client{key: key}
}
