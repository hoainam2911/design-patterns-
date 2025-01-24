package main

import (
	"fmt"
)

type Dress interface {
	getColor() string
}

type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var dressFactorySingleInstance = &DressFactory{
	dressMap: make(map[string]Dress),
}

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	switch dressType {
	case TerroristDressType:
		d.dressMap[dressType] = newTerroristDress()
	case CounterTerroristDressType:
		d.dressMap[dressType] = newCounterTerroristDress()
	default:
		return nil, fmt.Errorf("wrong dress type passed")
	}

	return d.dressMap[dressType], nil
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) setLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 0),
		counterTerrorists: make([]*Player, 0),
	}
}

func (g *Game) addTerrorist() {
	player := newPlayer("T", TerroristDressType)
	g.terrorists = append(g.terrorists, player)
}

func (g *Game) addCounterTerrorist() {
	player := newPlayer("CT", CounterTerroristDressType)
	g.counterTerrorists = append(g.counterTerrorists, player)
}

func main() {
	game := newGame()
	game.addTerrorist()
	game.addTerrorist()
	game.addTerrorist()
	game.addCounterTerrorist()
	game.addCounterTerrorist()

	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressType: %s, Color: %s\n", dressType, dress.getColor())
	}
}