package schemas

import (
	"MCMPDL/curseforge/schemas/enums/core_api_status"
	"MCMPDL/curseforge/schemas/enums/core_status"
	"time"
)

type Game struct {
	Id           int                           `json:"id"`
	Name         string                        `json:"name"`
	Slug         string                        `json:"slug"`
	DateModified time.Time                     `json:"dateModified"`
	Assets       GameAssets                    `json:"assets"`
	Status       core_status.CoreStatus        `json:"status"`
	ApiStatus    core_api_status.CoreApiStatus `json:"apiStatus"`
}
