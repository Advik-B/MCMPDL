package service

import (
	"MCMPDL/curseforge"
)

type Service struct {
	api *curseforge.API
}

func New(api *curseforge.API) *Service {
	return &Service{
		api: api,
	}
}
