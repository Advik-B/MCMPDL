package curseforge

import (
	"MCMPDL/curseforge/schemas"
	"encoding/json"
	"strconv"
)

type Service struct {
	api *API
}

func NewService(api *API) *Service {
	return &Service{
		api: api,
	}
}

func (s *Service) GetGames() ([]schemas.Game, error) {
	var gamesResponse struct {
		Data []schemas.Game `json:"data"`
	}
	res, err := s.api.Fetch("/v1/games")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &gamesResponse)
	if err != nil {
		return nil, err
	}
	return gamesResponse.Data, nil
}

func (s *Service) GetGame(gameId int) (schemas.Game, error) {
	gameResponse := struct {
		Data schemas.Game `json:"data"`
	}{}
	res, err := s.api.Fetch("/v1/games/" + strconv.Itoa(gameId))
	if err != nil {
		return gameResponse.Data, err
	}
	err = json.Unmarshal(res, &gameResponse)
	if err != nil {
		return gameResponse.Data, err
	}
	return gameResponse.Data, nil
}

func (s *Service) GetVersions(gameId int) (schemas.GetVersionsResponseV1, error) {
	versionsResponse := schemas.GetVersionsResponseV1{}
	res, err := s.api.Fetch("/v1/games/" + strconv.Itoa(gameId) + "/versions")
	if err != nil {
		return versionsResponse, err
	}
	err = json.Unmarshal(res, &versionsResponse)
	if err != nil {
		return versionsResponse, err
	}
	return versionsResponse, nil
}

func (s *Service) GetVersionTypes(gameId int) (schemas.GetVersionTypesResponse, error) {
	versionTypesResponse := schemas.GetVersionTypesResponse{}
	res, err := s.api.Fetch("/v1/games/" + strconv.Itoa(gameId) + "/version-types")
	if err != nil {
		return versionTypesResponse, err
	}
	err = json.Unmarshal(res, &versionTypesResponse)
	if err != nil {
		return versionTypesResponse, err
	}
	return versionTypesResponse, nil
}

func (s *Service) GetVersionsV2(gameId int) (schemas.GetVersionsResponseV2, error) {
	versionsResponse := schemas.GetVersionsResponseV2{}
	res, err := s.api.Fetch("/v2/games/" + strconv.Itoa(gameId) + "/versions")
	if err != nil {
		return versionsResponse, err
	}
	err = json.Unmarshal(res, &versionsResponse)
	if err != nil {
		return versionsResponse, err
	}
	return versionsResponse, nil
}

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
