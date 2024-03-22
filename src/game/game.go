package game

import (
	"battleship-2/src/player"
	"fmt"
)

type Battleship struct {
	player1    *player.Player
	player2    *player.Player
	gridSize   int
	numShip    int
	numMissile int
}

func NewGame(nShp int, gridSize int, nMissile int, p1ShipPos []string, p2ShipPos []string) *Battleship {
	b := Battleship{
		player1:    player.NewPlayer(nShp, nMissile, gridSize),
		player2:    player.NewPlayer(nShp, nMissile, gridSize),
		gridSize:   gridSize,
		numShip:    nShp,
		numMissile: nMissile,
	}

	b.player1.SetShips(p1ShipPos)
	b.player2.SetShips(p2ShipPos)
	return &b
}

func (b *Battleship) ProceedGame(p1MissilePos []string, p2MissilePos []string) {
	b.player1.AttackPlayer(b.player2, p1MissilePos)
	b.player2.AttackPlayer(b.player1, p2MissilePos)
}

func (b Battleship) Result() {
	fmt.Println("Player 1")
	b.player1.GetBatleground().PrintSurface()
	fmt.Println()
	fmt.Println("Player 2")
	b.player2.GetBatleground().PrintSurface()

	p1TotalPoint := b.player1.GetTotalPoints()
	p2TotalPoint := b.player2.GetTotalPoints()
	fmt.Printf("\nP1: %d\n", p1TotalPoint)
	fmt.Printf("P2: %d\n", p2TotalPoint)

	if p1TotalPoint > p2TotalPoint {
		fmt.Println("Player 1 wins")
	} else if p1TotalPoint < p2TotalPoint {
		fmt.Println("Player 2 wins")
	} else {
		fmt.Println("It is draw")
	}
}
