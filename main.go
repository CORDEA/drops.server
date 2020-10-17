package main

import (
	"drops/etsy"
	"drops/items"
	"drops/user"
	"github.com/gin-gonic/gin"
	"os"
)

var router = gin.Default()

func main() {
	key := os.Getenv("ETSY_API_KEY")
	client := etsy.NewClient(key)
	user.AddRoutes(router)
	items.NewRoute(client).Apply(router)
	router.Run()
}
