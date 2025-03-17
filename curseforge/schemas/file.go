package schemas

import (
	"MCMPDL/curseforge/schemas/enums/file_release_type"
	"MCMPDL/curseforge/schemas/enums/file_status"
	"time"
)

type File struct {
	Id                   int                               `json:"id"`
	GameId               int                               `json:"gameId"`
	ModId                int                               `json:"modId"`
	IsAvailable          bool                              `json:"isAvailable"`
	DisplayName          string                            `json:"displayName"`
	FileName             string                            `json:"fileName"`
	ReleaseType          file_release_type.FileReleaseType `json:"releaseType"`
	FileStatus           file_status.FileStatus            `json:"fileStatus"`
	Hashes               []FileHash                        `json:"hashes"`
	FileDate             time.Time                         `json:"fileDate"`
	FileLength           int                               `json:"fileLength"`
	DownloadCount        int                               `json:"downloadCount"`
	FileSizeOnDisk       int                               `json:"fileSizeOnDisk"`
	DownloadUrl          string                            `json:"downloadUrl"`
	GameVersions         []string                          `json:"gameVersions"`
	SortableGameVersions []SortableGameVersion             `json:"sortableGameVersions"`
	Dependencies         []FileDependency                  `json:"dependencies"`
	ExposeAsAlternative  bool                              `json:"exposeAsAlternative"`
	ParentProjectFileId  int                               `json:"parentProjectFileId"`
	AlternateFileId      int                               `json:"alternateFileId"`
	IsServerPack         bool                              `json:"isServerPack"`
	ServerPackFileId     int                               `json:"serverPackFileId"`
	IsEarlyAccessContent bool                              `json:"isEarlyAccessContent"`
	EarlyAccessEndDate   time.Time                         `json:"earlyAccessEndDate"`
	FileFingerprint      int                               `json:"fileFingerprint"`
	Modules              []FileModule                      `json:"modules"`
}
