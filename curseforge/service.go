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

type GameResponse struct {
	Data schemas.Game `json:"data"`
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
	gameResponse := GameResponse{}
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
