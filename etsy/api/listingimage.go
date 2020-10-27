package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ListingImages struct {
	Count   int            `json:"count"`
	Results []ListingImage `json:"results"`
}

type ListingImage struct {
	Id       int64  `json:"id"`
	Rank     int    `json:"rank"`
	ImageUrl string `json:"url_570xN"`
}

func (c *Client) FindImages(id int64) (*ListingImages, error) {
	u, err := url.Parse(BaseUrl + "/listings/" + strconv.FormatInt(id, 10) + "/images")
	if err != nil {
		return nil, err
	}
	values := u.Query()
	values.Add("api_key", c.key)
	u.RawQuery = values.Encode()
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
