package schemas

type GameVersionType struct {
	Id         int    `json:"id"`
	GameId     int    `json:"gameId"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	IsSyncable bool   `json:"isSyncable"`
	Status     int    `json:"status"`
}
