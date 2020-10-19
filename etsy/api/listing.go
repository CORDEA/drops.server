package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const TaxonomyId = "1203"

type Listings struct {
	Count   int       `json:"count"`
	Results []Listing `json:"results"`
}

type Listing struct {
	Id                 string   `json:"listing_id"`
	State              string   `json:"state"`
	UserId             string   `json:"user_id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	CreatedAt          float32  `json:"creation_tsz"`
	EndedAt            float32  `json:"ending_tsz"`
	UpdatedAt          float32  `json:"last_modified_tsz"`
	Price              string   `json:"price"`
	CurrencyCode       string   `json:"currency_code"`
	Quantity           int      `json:"quantity"`
	Tags               []string `json:"tags"`
	TaxonomyId         int      `json:"taxonomy_id"`
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

func (c *Client) FindActiveEarrings() (*Listings, error) {
	u, err := url.Parse(BaseUrl + "/listings/active")
	if err != nil {
		return nil, err
	}
	values := u.Query()
	values.Add("api_key", c.key)
	values.Add("taxonomy_id", TaxonomyId)
	response, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	var list Listings
	if err := json.Unmarshal(body, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
