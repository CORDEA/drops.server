package main

import "github.com/gin-gonic/gin"

var router = gin.Default()

func main() {
	addUserRoutes(router)
	addItemRoutes(router)
	router.Run()
}
