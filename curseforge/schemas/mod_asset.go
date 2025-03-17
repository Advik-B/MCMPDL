package schemas

type ModAsset struct {
	Id           int    `json:"id"`
	ModId        int    `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	Url          string `json:"url"`
}
