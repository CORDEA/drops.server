package main

import (
	"github.com/CORDEA/drops.server/etsy/api"
	"github.com/CORDEA/drops.server/etsy/repository"
	"github.com/CORDEA/drops.server/etsy/usecase"
	"github.com/CORDEA/drops.server/items"
	"github.com/CORDEA/drops.server/user"
	"github.com/gin-gonic/gin"
	"os"
)

var router = gin.Default()

func main() {
	key := os.Getenv("ETSY_API_KEY")
	client := api.NewClient(key)
	earringsRepo := repository.NewEarringsRepository(client)
	getEarrings := usecase.NewGetEarrings(earringsRepo)

	user.NewRoute(getEarrings).Apply(router)
	items.NewRoute(getEarrings).Apply(router)
	router.Run()
}
