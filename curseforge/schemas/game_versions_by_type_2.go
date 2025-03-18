package schemas

type GameVersionsByType2 struct {
	Type     int           `json:"type"`
	Versions []GameVersion `json:"versions"`
}
