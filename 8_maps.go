package main

import (
	"fmt"
)

func main() {
	// Declaration may look like this
	// statePopulations := make(map[string]int)
	// Or with initialization
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	// slice is not a valid index, but an array is
	// it needs to be comparable
	m := map[[3]int]string{}
	fmt.Println(m)

	// Add an item to a map
	statePopulations["Georgia"] = 10310371
	// Remove map item
	delete(statePopulations, "Ohio")

	// Apparently order of a map is not guaranteed (it was order in my case though)
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Texas"])
	// Removed items are not clearly indicated
	fmt.Println(statePopulations["Ohio"])

	//To check if a map include an item
	pop, ok := statePopulations["Ohio"]
	// use the following to ignore value itself
	// _, ok := statePopulations["Ohio"]
	fmt.Println(pop, ok)

	// Lenght function works
	fmt.Println(len(statePopulations))

	// Look out, map works like a slice, it's not copied
	// Texas is deleted from the map, so both references do not include it in the listing
	sp := statePopulations
	delete(sp, "Texas")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}
