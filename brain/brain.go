package brain

import (
	"fmt"
	"strings"

	"lemin/parse"
)

const ants = 5

// calculatePaths finds the starting and finishing rooms and updates the fullness of each room based on the paths.
func Lemin(c map[string]parse.Room, p [][]string) [][]string {
	s, f := findStartAndFinish(c)

	// Update the fullness of the rooms based on the paths
	InitialiseRoomFullness(c, p)

	// Set the initial fullness of the starting room
	updateRoomFullnessValue(c[s], 400.0)

	fmt.Println(c, p, f)

	// Initialize the starting positions for each ant
	n, x := initializeAntPositions(s)
	for {
		for d := 1; d <= ants; d++ {
			free(&c, f) // Mark the finish room as free

			if n[d-1] == f {
				x[d-1] = append(x[d-1], "")
				continue
			}

			r := c[n[d-1]]
			// Fix the call by passing f as the last argument
			i, _ := Chose_next_move(c, r, p, x[d-1], d, f)

			// Move the ant to the next room
			moveAnt(c, r, &n, &x, i, d)
		}
		if n[ants-1] == f {
			break
		}
	}

	return x
}

// findStartAndFinish identifies the starting and finishing rooms in the colony.
func findStartAndFinish(c map[string]parse.Room) (string, string) {
	var s, f string
	for l, r := range c {
		if r.Fullness == 1.0 {
			s = l
			r.Fullness = 0.0
			c[l] = r
		} else if r.Fullness == 2.0 {
			f = l
			r.Fullness = 0.0
			c[l] = r
		}
	}
	return s, f
}

// InitialiseRoomFullness initialises the fullness of the rooms based on the paths.
func InitialiseRoomFullness(c map[string]parse.Room, p [][]string) {
	for i := len(p) - 1; i >= 0; i-- {
		for k, l := range p[i] {
			r := c[l]
			r.Fullness = float64(len(p[i][k:])) * 100.0
			c[l] = r
		}
	}
}

// updateRoomFullnessValue sets the fullness of a room to the specified value.
func updateRoomFullnessValue(r parse.Room, fullness float64) parse.Room {
	r.Fullness = fullness
	return r
}

// initializeAntPositions sets the initial positions of the ants in the colony.
func initializeAntPositions(start string) ([ants]string, [][]string) {
	n := [ants]string{}
	x := make([][]string, ants)
	for i := range n {
		n[i] = start
		x[i] = []string{fmt.Sprintf("L%d-%s", i+1, start)}
	}
	return n, x
}

// chose_next_move finds the minimum fullness value among the possible next rooms for an ant.
func Chose_next_move(c map[string]parse.Room, r parse.Room, p [][]string, currentPath []string, antIndex int, f string) (float64, string) {
	i := float64(len(p[0]) * 100.0)
	t := ""
	for _, link := range r.Links {
		r1 := c[string(link)]
		if i >= r1.Fullness && check(currentPath, antIndex, string(link)) && !r1.Empty || string(link) == f {
			i = r1.Fullness
			t = string(link)
		}
	}
	return i, t
}

// moveAnt attempts to move an ant to the next room based on fullness constraints.
func moveAnt(c map[string]parse.Room, r parse.Room, n *[ants]string, x *[][]string, minFullness float64, d int) {
	for w, link := range r.Links {
		r1 := c[string(link)]

		if r1.Fullness == minFullness && check((*x)[d-1], d, string(link)) && !r1.Empty { // Dereference x to access the slice
			free(&c, n[d-1])
			r1.Fullness += 100.0
			r1.Empty = true
			c[string(link)] = r1
			n[d-1] = string(link)
			(*x)[d-1] = append((*x)[d-1], fmt.Sprintf("L%d-%s", d, n[d-1])) // Dereference x to append to it
			break
		}
		if w == len(r.Links)-1 {
			(*x)[d-1] = append((*x)[d-1], "")
		}
	}
}

// free marks a room as free and decreases its fullness.
func free(c *map[string]parse.Room, p string) {
	r := (*c)[p]
	r.Empty = false
	r.Fullness -= 100.0
	(*c)[p] = r
}

// check verifies if an ant can move to a specific room.
func check(a []string, d int, b string) bool {
	for _, c := range a {
		if c == fmt.Sprintf("L%d-%s", d, b) {
			return false
		}
	}
	return true
}

// DisplayResult outputs the results of the path calculation.
func DisplayResult(result [][]string) {
	fmt.Println("\nTotal weight:", result)
	for i := 1; i < len(result[0]); i++ {
		x := []string{}
		for _, v := range result {
			if v[i] != "" {
				x = append(x, v[i])
			}
		}
		fmt.Println(strings.Join(x, " "))
	}
}
