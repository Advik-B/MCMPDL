package service

import (
	"MCMPDL/curseforge/schemas"
	"encoding/json"
)

func (s *Service) GetCategories() ([]schemas.Category, error) {
	var categoriesResponse struct {
		Data []schemas.Category `json:"data"`
	}
	res, err := s.api.Fetch("/v1/categories")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &categoriesResponse)
	if err != nil {
		return nil, err
	}
	return categoriesResponse.Data, nil
}
