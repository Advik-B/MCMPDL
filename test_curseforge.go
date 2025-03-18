package main

import (
	"MCMPDL/curseforge"
	"MCMPDL/curseforge/service"
	"fmt"
	"os"
)

func Main() {
	api, err := curseforge.NewAPI(os.Getenv("CURSEFORGE_API_KEY"), true, false)
	if err != nil {
		panic(err)
	}
	//fmt.Println(api) // NOTE: NEVER DO THIS UNLESS YOU WANT YOUR API KEY TO BE LEAKED
	service_ := service.New(api)
	//games, _ := service_.GetGames()
	//for _, game := range games.Data {
	//	fmt.Printf("%d: "+game.Name+"\n", game.Id)
	//}
	//
	game, x := service_.GetGame(432)
	fmt.Println(game, x)
}
