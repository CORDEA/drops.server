package main

import (
	"github.com/CORDEA/drops.server/etsy/api"
	"github.com/CORDEA/drops.server/etsy/repository"
	"github.com/CORDEA/drops.server/etsy/usecase"
	"github.com/CORDEA/drops.server/items"
	"github.com/CORDEA/drops.server/user"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"os"
	"time"
)

var router = gin.Default()

func main() {
	key := os.Getenv("ETSY_API_KEY")
	client := api.NewClient(key)
	c := cache.New(1*time.Hour, 30*time.Minute)
	earringsRepo := repository.NewEarringsRepository(client, c)
	getEarrings := usecase.NewGetEarrings(earringsRepo)

	user.NewRoute(getEarrings).Apply(router)
	items.NewRoute(getEarrings).Apply(router)
	router.Run()
}
