package service

import (
	"MCMPDL/curseforge/schemas"
)

func (s *Service) SearchMods(gameID int) (schemas.SearchModsResponse, error) {
	// TODO: Implement this
	// https://docs.curseforge.com/rest-api/#search-mods
	return schemas.SearchModsResponse{}, nil
}
