package etsy

type ListingImages struct {
	Count   int            `json:"count"`
	Results []ListingImage `json:"results"`
}

type ListingImage struct {
	Id       string `json:"id"`
	ImageUrl string `json:"url_570xN"`
}
