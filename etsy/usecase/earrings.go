package usecase

import (
	"github.com/CORDEA/drops.server/etsy/model"
	"github.com/CORDEA/drops.server/etsy/repository"
)

type GetEarrings struct {
	repository *repository.EarringsRepository
}

func NewGetEarrings(repository *repository.EarringsRepository) *GetEarrings {
	return &GetEarrings{repository}
}

func (g *GetEarrings) Execute() (*model.Items, error) {
	earrings, err := g.repository.FindAll()
	if err != nil {
		return nil, err
	}
	var s []model.Item
	for _, result := range earrings.Results {
		images, err := g.repository.FindImages(result.Id)
		if err != nil {
			return nil, err
		}
		s = append(s, model.NewItem(result, images))
	}
	items := model.NewItems(s)
	return &items, nil
}
