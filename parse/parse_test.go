package parse_test

import (
	"fmt"
	"testing"

	"lemin/parse"
)

func TestParse(t *testing.T) {
	colony := parse.Parse("test.txt")
	c := map[string]parse.Room{
		"0": {0, 3, 0.0, []parse.Link{"2"}},
		"1": {8, 3, 0.0, []parse.Link{"3"}},
		"2": {2, 5, 0.0, []parse.Link{"0", "3"}},
		"3": {4, 0, 0.0, []parse.Link{"2", "1"}},
	}
	if CompareMaps(c, colony) {
		fmt.Println("Maps are equal")
	} else {
		t.Fail()
	}
	fmt.Println(colony)
	fmt.Println(c)
}

// CompareMaps compares two maps of rooms.

func CompareMaps(map1, map2 map[string]parse.Room) bool {
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
		case value1.X != value2.X:
			return false
		case value1.Y != value2.Y:
			return false
		case value1.Fullness != value2.Fullness:
			return false
		case len(value1.Links) != len(value2.Links):
			return false
		default:
			for i := range value1.Links {
				if value1.Links[i] != value2.Links[i] {
					return false
				}
			}
		}
	}

	return true
}
