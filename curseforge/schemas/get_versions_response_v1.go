package schemas

type GetVersionsResponseV1 struct {
	Data []GameVersionsByType `json:"data"`
}
