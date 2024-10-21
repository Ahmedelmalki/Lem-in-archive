package parse_test

import (
	"fmt"
	"lemin/parse"
	"testing"
)

func TestParse(t *testing.T) {
	colony := parse.Parse("test.txt")
	c := map[string][]string{
		"0": {"1", "0"},
		"1": {"1", "0"},
		"2": {"1", "0", "1", "3"},
		"3": {"1", "0", "1", "1"},
	}
	if CompareMaps(c, colony.Rooms) {
		fmt.Println("Maps are equal")
	} else {
		t.Fail()
	}
}

// CompareMaps compares two maps of rooms.

func CompareMaps(map1, map2 map[string][]string) bool {
	// Check if lengths are the same
	if len(map1) != len(map2) {
		return false
	}

	// Compare key-value pairs
	for key, value1 := range map1 {
		value2, exists := map2[key]
		switch {
		case !exists:
			return false
		case len(value1) != len(value2):
			return false
		default:
			for i := range value1 {
				if value1[i] != value2[i] {
					return false
				}
			}
		}
	}

	return true
}
