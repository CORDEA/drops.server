package items

type Item struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageUrl string `json:"image_url"`
}
