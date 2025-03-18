package schemas

type GameVersionsByType struct {
	Type     int      `json:"type"`
	Versions []string `json:"versions"`
}
