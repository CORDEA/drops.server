package repository

import "github.com/CORDEA/drops.server/etsy"

type EarringsRepository struct {
	client *etsy.Client
}

func (r *EarringsRepository) FindAll() (*etsy.Listings, error) {
	return r.client.FindActiveEarrings()
}

func (r *EarringsRepository) FindImages(id string) (*etsy.ListingImages, error) {
	return r.client.FindImages(id)
}
