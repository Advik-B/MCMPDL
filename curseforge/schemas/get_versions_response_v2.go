package schemas

type GetVersionsResponseV2 struct {
	Data []GameVersionsByType2 `json:"data"`
}
