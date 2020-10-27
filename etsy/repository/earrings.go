package repository

import (
	"github.com/CORDEA/drops.server/etsy/api"
	"github.com/patrickmn/go-cache"
)

const EarringsCacheKey = "earrings"

type EarringsRepository struct {
	client *api.Client
	cache  *cache.Cache
}

func NewEarringsRepository(client *api.Client, cache *cache.Cache) *EarringsRepository {
	return &EarringsRepository{client, cache}
}

func (r *EarringsRepository) FindAll() (*api.Listings, error) {
	if c, found := r.cache.Get(EarringsCacheKey); found {
		return c.(*api.Listings), nil
	}
	res, err := r.client.FindActiveEarrings()
	if err == nil {
		r.cache.Set(EarringsCacheKey, res, cache.DefaultExpiration)
	} else {
		r.cache.Delete(EarringsCacheKey)
	}
	return res, err
}

func (r *EarringsRepository) FindImages(id int64) (*api.ListingImages, error) {
	return r.client.FindImages(id)
}
