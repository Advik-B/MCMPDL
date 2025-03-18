package curseforge

import (
	"fmt"
	"os"
)

func Main() {
	api, err := NewAPI(os.Getenv("CURSEFORGE_API_KEY"), true, false)
	if err != nil {
		panic(err)
	}
	//fmt.Println(api) // NOTE: NEVER DO THIS UNLESS YOU WANT YOUR API KEY TO BE LEAKED
	service := NewService(api)
	//games, _ := service.GetGames()
	//for _, game := range games.Data {
	//	fmt.Printf("%d: "+game.Name+"\n", game.Id)
	//}
	//
	game, x := service.GetGame(432)
	fmt.Println(game, x)
}
