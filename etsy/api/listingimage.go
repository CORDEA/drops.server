package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ListingImages struct {
	Count   int            `json:"count"`
	Results []ListingImage `json:"results"`
}

type ListingImage struct {
	Id       string `json:"id"`
	Rank     string `json:"rank"`
	ImageUrl string `json:"url_570xN"`
}

func (c *Client) FindImages(id string) (*ListingImages, error) {
	u, err := url.Parse(BaseUrl + "/listings/" + id + "/images")
	if err != nil {
		return nil, err
	}
	values := u.Query()
	values.Add("api_key", c.key)
	response, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	var images ListingImages
	if err := json.Unmarshal(body, &images); err != nil {
		return nil, err
	}
	return &images, nil
}
