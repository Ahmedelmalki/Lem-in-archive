package brain

import (
	"fmt"
	"strings"

	"lemin/parse"
)

type Colony map[string]parse.Room

var ants int

// calculatePaths finds the starting and finishing rooms and updates the fullness of each room based on the paths.
func Lemin(colo parse.Colony, p [][]string) [][]string {
	c := Colony(colo.Rooms)
	s, f := colo.Strat, colo.Finish
	ants = colo.Ants
	// Update the fullness of the rooms based on the paths
	InitialiseRoomFullness(c, p)

	// Set the initial fullness of the starting room
	updateRoomFullnessValue(c[s], 400.0)

	// fmt.Println(c, p, f)

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
			i, next := Chose_next_move(c, r, p, x[d-1], d, f)

			// Move the ant to the next room
			if next != "" {
				moveAnt(c, r, &n, &x, i, d)
			}
		}
		if n[ants-1] == f {
			break
		}
	}

	return x
}

// InitialiseRoomFullness initialises the fullness of the rooms based on the paths.
func InitialiseRoomFullness(c Colony, p [][]string) {
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

// initializeAntPositions sets the initial positions of the ants in the Colony.
func initializeAntPositions(start string) ([]string, [][]string) {
	n := make([]string, ants)
	x := make([][]string, ants)
	for i := range n {
		n[i] = start
		x[i] = []string{fmt.Sprintf("L%d-%s", i+1, start)}
	}
	return n, x
}

// chose_next_move finds the minimum fullness value among the possible next rooms for an ant.
func Chose_next_move(c Colony, r parse.Room, p [][]string, currentPath []string, antIndex int, f string) (float64, string) {
	i := float64(len(p[0]) * 100.0)
	t := ""
	for link := range r.Links {
		r1 := c[link]
		if i >= r1.Fullness && check(currentPath, antIndex, link) && !r1.Empty || link == f {
			i = r1.Fullness
			t = link
		}
	}
	return i, t
}

// moveAnt attempts to move an ant to the next room based on fullness constraints.
func moveAnt(c Colony, r parse.Room, n *[]string, x *[][]string, minFullness float64, d int) {
	w := -1
	for link := range r.Links {
		w++
		r1 := c[link]

		if r1.Fullness == minFullness && check((*x)[d-1], d, link) && !r1.Empty { // Dereference x to access the slice
			free(&c, (*n)[d-1])
			r1.Fullness += 100.0
			r1.Empty = true
			c[link] = r1
			(*n)[d-1] = link
			(*x)[d-1] = append((*x)[d-1], fmt.Sprintf("L%d-%s", d, (*n)[d-1])) // Dereference x to append to it
			break
		}
		if w == len(r.Links)-1 {
			(*x)[d-1] = append((*x)[d-1], "")
		}
	}
}

// free marks a room as free and decreases its fullness.
func free(c *Colony, p string) {
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
	for i := 1; i < len(result[len(result)-1]); i++ {
		x := []string{}
		for _, v := range result {
			if v[i] != "" {
				x = append(x, v[i])
			}
		}
		fmt.Println(strings.Join(x, " "))
	}
}
