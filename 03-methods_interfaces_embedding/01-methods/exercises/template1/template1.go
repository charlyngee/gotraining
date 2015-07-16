// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Rj0QfwVPhX

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initalize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// Add imports.

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.
type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a batter.
func (pl player) average() float64 {
	return float64(pl.hits) / float64(pl.atBats)
}

// main is the entry point for the application.
func main() {
	// Create a slice of players and populate each player
	// with field values.
	players := []player{
		{"one", 5, 1},
		{"two", 10, 5},
		{"three", 15, 5},
		{"four", 20, 2},
	}

	// Display the batting average for each player in the slice.
	for _, pl := range players {
		fmt.Printf("%s\t%f\n", pl.name, pl.average())
	}
}
