package schemas

import (
	"time"
)

type Category struct {
	Id               int       `json:"id"`
	GameId           int       `json:"gameId"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	Url              string    `json:"url"`
	IconUrl          string    `json:"iconUrl"`
	DateModified     time.Time `json:"dateModified"`
	IsClass          bool      `json:"isClass"`
	ClassId          int       `json:"classId"`
	ParentCategoryId int       `json:"parentCategoryId"`
	DisplayIndex     int       `json:"displayIndex"`
}
