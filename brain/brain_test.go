package brain_test

import (
	"testing"

	"lemin/brain"
	"lemin/parse"
)

var c parse.Colony

func init() {
	c = parse.Parse("../parse/test.txt")
}

func TestChose_next_move(t *testing.T) {
	// TODO: add tests for Chose_next_move
	a = brain.Colony(
		u := make(map[string]parse.Room)
		u["0"]= parse.Room{X: 0, Y: 3, Fullness: 0.0, Empty: false, Links: map[string]parse.Link{"2": {P: false, R: u["2"]}, "3": {P: false, R: u["3"]}, "1": {P: false, R: u["1"]}}}
		u["1"]= parse.Room{X: 8, Y: 3, Fullness: 0.0, Empty: false, Links: map[string]parse.Link{"3": {P: false, R: u["3"]}, "2": {P: false, R: u["2"]}, "4": {P: false, R: u["4"]}, "5": {P: false, R: u["5"]}}}
		u["2"]= parse.Room{X: 2, Y: 5, Fullness: 0.0, Empty: false, Links: map[string]parse.Link{"0": {P: false, R: u["0"]}, "3": {P: false, R: u["3"]}, "4": {P: false, R: u["4"]}, "5": {P: false, R: u["5"]}}}
		u["3"]= parse.Room{X: 4, Y: 0, Fullness: 0.0, Empty: false, Links: map[string]parse.Link{"2": {P: false, R: u["2"]}, "1": {P: false, R: u["1"]}, "4": {P: false, R: u["4"]}, "5": {P: false, R: u["5"]}}}
		u["4"]= parse.Room{X: 6, Y: 0, Fullness: 0.0, Empty: false, Links: map[string]parse.Link{"2": {P: false, R: u["2"]}, "3": {P: false, R: u["3"]}, "5": {P: false, R: u["5"]}, "6": {P: false, R: u["6"]}}}
		u["5"]= parse.Room{X: 10, Y: 0, Fullness: 0.0, Empty: false, Links: map[string]parse.Link{"1": {P: false, R: u["1"]}, "2": {P: false, R: u["2"]}, "3": {P: false, R: u["3"]}, "4": {P: false, R: u["4"]}, "6": {P: false, R: u["6"]}}}
	)

	x, false, y := brain.Chose_next_move(a, b, c, d, e, f)
	t.SkipNow() // Placeholder for actual tests
}

func TestMoveAnt(t *testing.T) {
	// TODO: add tests for moveAnt
	t.SkipNow() // Placeholder for actual tests
}

func TestLemin(t *testing.T) {
	// TODO: add tests for Lemin
	t.SkipNow() // Placeholder for actual tests
}

func TestInitialiseRoomFullness(t *testing.T) {
	// Example input for the InitialiseRoomFullness function
	p := [][]string{
		{"0", "2", "3", "1"},
		{"0", "2", "4", "1"},
		{"0", "6", "4", "1"},
		{"0", "6", "4", "2", "3", "1"},
	}

	// Initialize rooms (assuming c.Rooms is a slice of Room structs)
	brain.InitialiseRoomFullness(c.Rooms, p)

	// Define the expected result
	expect := map[string]float64{
		"1": 100.0,
		"2": 300.0,
		"3": 200.0,
		"4": 200.0,
		"6": 300.0,
	}

	// Check if the actual result matches the expected result
	for roomID, expectedFullness := range expect {
		actual, exists := c.Rooms[roomID]
		if !exists {
			t.Errorf("Room %s does not exist", roomID)
			continue
		}

		// Check if the actual fullness matches the expected fullness
		if expectedFullness != actual.Fullness {
			t.Errorf("\n\nExpected %f, got %f\nin room: %s", expectedFullness, actual.Fullness, roomID)
		}
	}
}
