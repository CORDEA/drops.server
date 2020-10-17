package main

import (
	"drops/items"
	"drops/user"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	user.AddRoutes(router)
	items.AddRoutes(router)
	router.Run()
}
