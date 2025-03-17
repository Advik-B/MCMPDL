package schemas

import (
	"MCMPDL/curseforge/schemas/enums/mod_status"
	"time"
)

type Mod struct {
	Id                            int                  `json:"id"`
	GameId                        int                  `json:"gameId"`
	Name                          string               `json:"name"`
	Slug                          string               `json:"slug"`
	Links                         ModLinks             `json:"links"`
	Summary                       string               `json:"summary"`
	Status                        mod_status.ModStatus `json:"status"`
	DownloadCount                 int                  `json:"downloadCount"`
	IsFeatured                    bool                 `json:"isFeatured"`
	PrimaryCategoryId             int                  `json:"primaryCategoryId"`
	Categories                    []Category           `json:"categories"`
	ClassId                       int                  `json:"classId"`
	Authors                       []ModAuthor          `json:"authors"`
	Logo                          ModAsset             `json:"logo"`
	Screenshots                   []ModAsset           `json:"screenshots"`
	MainFileId                    int                  `json:"mainFileId"`
	LatestFiles                   []File               `json:"latestFiles"`
	LatestFilesIndexes            []FileIndex          `json:"latestFilesIndexes"`
	LatestEarlyAccessFilesIndexes []FileIndex          `json:"latestEarlyAccessFilesIndexes"`
	DateCreated                   time.Time            `json:"dateCreated"`
	DateModified                  time.Time            `json:"dateModified"`
	DateReleased                  time.Time            `json:"dateReleased"`
	AllowModDistribution          bool                 `json:"allowModDistribution"`
	GamePopularityRank            int                  `json:"gamePopularityRank"`
	IsAvailable                   bool                 `json:"isAvailable"`
	ThumbsUpCount                 int                  `json:"thumbsUpCount"`
	Rating                        int                  `json:"rating"`
}
