package tests

import (
	"testing"

	"lemin/brain"
	"lemin/parse"
)

func TestInitialiseRoomFullness(t *testing.T) {
	c := parse.Parse("../parse/test.txt")                                                                             // Setup the test data
	p := [][]string{{"0", "2", "3", "1"}, {"0", "2", "4", "1"}, {"0", "6", "4", "1"}, {"0", "6", "4", "2", "3", "1"}} // Example input for the UpdateRoomFullness function

	brain.InitialiseRoomFullness(c, p)

	// Define the expected result
	expect := map[string]float64{
		"1": 100.0,
		"3": 200.0,
		"2": 300.0,
		"6": 300.0,
		"4": 200.0,
	}
	for i, j := range expect {
		actual := c[i]

		// Check if the actual result matches the expected result
		if j != actual.Fullness {
			t.Errorf("\n\nExpected %f, got %f\nin room :%s with c:%f", j, actual.Fullness, i, actual.Fullness)
		}
	}
}
