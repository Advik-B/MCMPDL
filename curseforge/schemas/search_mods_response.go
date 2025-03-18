package schemas

type SearchModsResponse struct {
	Data       []Mod      `json:"data"`
	Pagination Pagination `json:"pagination"`
}
