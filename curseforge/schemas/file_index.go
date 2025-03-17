package schemas

import (
	"MCMPDL/curseforge/schemas/enums/file_release_type"
	"MCMPDL/curseforge/schemas/enums/mod_loader_type"
)

type FileIndex struct {
	GameVersion       string                            `json:"gameVersion"`
	FileId            int                               `json:"fileId"`
	Filename          string                            `json:"filename"`
	ReleaseType       file_release_type.FileReleaseType `json:"releaseType"`
	GameVersionTypeId int                               `json:"gameVersionTypeId"`
	ModLoader         mod_loader_type.ModLoaderType     `json:"modLoader"`
}
