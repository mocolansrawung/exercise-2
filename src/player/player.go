package player

import (
	"battleship-2/src/battleground"
	"strconv"
	"strings"
)

type Player struct {
	activeShips int
	numShips    int
	ground      *battleground.Batleground
	totalPoints int
	totalMissle int
}

func NewPlayer(nShip int, nMissile int, gridSize int) *Player {
	p := Player{
		activeShips: nShip,
		numShips:    nShip,
		ground:      battleground.NewBattleground(gridSize),
		totalPoints: 0,
		totalMissle: nMissile,
	}

	return &p
}

func (p Player) GetTotalPoints() int {
	return p.totalPoints
}

func (p Player) GetBatleground() battleground.Batleground {
	return *p.ground
}

func (p *Player) SetDestroyShip(x, y int) {
	p.ground.SetDeadShip(x, y)
	p.activeShips -= 1
}

func (p *Player) SetMissile(x, y int) {
	p.ground.SetMissile(x, y)
}

func (p *Player) SetShips(p1Ships []string) {
	for _, ship := range p1Ships {
		coordinate := strings.Split(ship, ",")
		x, _ := strconv.Atoi(coordinate[0])
		y, _ := strconv.Atoi(coordinate[1])
		p.ground.SetShip(x, y)
	}
}

func (p *Player) AttackPlayer(p2 *Player, p1Missile []string) {
	for _, missile := range p1Missile {
		coordinate := strings.Split(missile, ",")
		x, _ := strconv.Atoi(coordinate[0])
		y, _ := strconv.Atoi(coordinate[1])
		if p2.GetBatleground().IsShip(x, y) {
			p2.SetDestroyShip(x, y)
			p.totalPoints += 1
		} else {
			p2.SetMissile(x, y)
		}
	}
}
