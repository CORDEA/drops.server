package repository

import (
	"github.com/CORDEA/drops.server/etsy/api"
)

type EarringsRepository struct {
	client *api.Client
}

func (r *EarringsRepository) FindAll() (*api.Listings, error) {
	return r.client.FindActiveEarrings()
}

func (r *EarringsRepository) FindImages(id string) (*api.ListingImages, error) {
	return r.client.FindImages(id)
}
