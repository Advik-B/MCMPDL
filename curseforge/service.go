package curseforge

import (
	"MCMPDL/curseforge/schemas"
	"encoding/json"
)

type Service struct {
	api *API
}

func NewService(api *API) *Service {
	return &Service{
		api: api,
	}
}

func (s *Service) GetGames() (schemas.GetGamesResponse, error) {
	games := schemas.GetGamesResponse{}
	res, err := s.api.Fetch("/v1/games")
	if err != nil {
		return games, err
	}
	err = json.Unmarshal(res, &games)
	if err != nil {
		return games, err
	}
	return games, nil
}

func (s *Service) GetGame(gameId int) (schemas.Game, error) {
	game := schemas.Game{}
	res, err := s.api.Fetch("/v1/games/" + string(gameId))
	if err != nil {
		return game, err
	}
	err = json.Unmarshal(res, &game)
	if err != nil {
		return game, err
	}
	return game, nil
}
