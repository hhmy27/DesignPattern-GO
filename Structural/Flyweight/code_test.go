package Flyweight

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrorismDressType)
	game.addCounterTerrorist(CounterTerrorismDressType)
	game.addCounterTerrorist(CounterTerrorismDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
	// output:
	// DressColorType: tDress
	// DressColor: red
	// DressColorType: ctDress
	// DressColor: green
}
