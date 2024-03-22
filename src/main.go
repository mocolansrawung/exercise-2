package main

import (
	"battleship-2/src/config"
	battleship "battleship-2/src/game"
	"fmt"
)

func main() {
	// game configuration
	var filepath string
	fmt.Println("Enter file path: ")
	fmt.Scan(&filepath)

	configFile := config.ConfigReader{}
	configFile.ReadFile(filepath)

	game := battleship.NewGame(configFile.GetNumShip(), configFile.GetGridSize(), configFile.GetNumMissile(), configFile.GetPlayer1Ship(), configFile.GetPlayer2Ship())
	game.ProceedGame(configFile.GetPlayer1Missile(), configFile.GetPlayer2Missile())
	game.Result()
}
