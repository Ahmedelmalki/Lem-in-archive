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
	a := brain.Colony(map[string]parse.Room{
		"0": {0, 3, 0.0, map[string]parse.Link{}},
        "1": {8, 3, 0.0, map[string]parse.Link{"3":true, "4":true}},
        "2": {2, 5, 0.0, map[string]parse.Link{}},
        "3": {4, 0, 0.0, map[string]parse.Link{}},
        "4": {6, 0, 0.0, map[string]parse.Link{}},
        "5": {10, 0, 0.0, map[string]parse.Link{}},
        "6": {14, 7, 0.0, map[string]parse.Link[]
	})

	x, y := brain.Chose_next_move(a, b, c, d, e, f)
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
