package model

import (
	"github.com/CORDEA/drops.server/etsy/api"
)

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	ImageUrls          []string `json:"image_urls"`
	CreatedAt          float32  `json:"created_at"`
	EndedAt            float32  `json:"ended_at"`
	UpdatedAt          float32  `json:"updated_at"`
	Price              string   `json:"price"`
	CurrencyCode       string   `json:"currency_code"`
	Tags               []string `json:"tags"`
	Materials          []string `json:"materials"`
	Views              int      `json:"views"`
	ItemWeight         int      `json:"item_weight"`
	ItemWeightUnit     string   `json:"item_weight_unit"`
	ItemLength         int      `json:"item_length"`
	ItemWidth          int      `json:"item_width"`
	ItemHeight         int      `json:"item_height"`
	ItemDimensionsUnit string   `json:"item_dimensions_unit"`
	Recipient          string   `json:"recipient"`
}

func NewItems(items []Item) Items {
	return Items{items}
}

func NewItem(listing api.Listing, images *api.ListingImages) Item {
	var img []string
	for _, result := range images.Results {
		img = append(img, result.ImageUrl)
	}
	return Item{
		Id:                 listing.Id,
		Name:               listing.Title,
		ImageUrls:          img,
		CreatedAt:          listing.CreatedAt,
		EndedAt:            listing.EndedAt,
		UpdatedAt:          listing.UpdatedAt,
		Price:              listing.Price,
		CurrencyCode:       listing.CurrencyCode,
		Tags:               listing.Tags,
		Materials:          listing.Materials,
		Views:              listing.Views,
		ItemWeight:         listing.ItemWeight,
		ItemWeightUnit:     listing.ItemWeightUnit,
		ItemLength:         listing.ItemLength,
		ItemWidth:          listing.ItemWidth,
		ItemHeight:         listing.ItemHeight,
		ItemDimensionsUnit: listing.ItemDimensionsUnit,
		Recipient:          listing.Recipient,
	}
}
