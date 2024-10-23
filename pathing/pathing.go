package pathing

import (
	"fmt"
	"lemin/static"
	"os"
	"sort"
)

func FindAllPath(colony *static.Colony) [][]string {

	// Start the backtracking process from the start room
	paths := Backtrack(colony)
	// Sort the paths based on length
	sort.SliceStable(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j]) // Compare lengths of inner slices
	})

	// Handle the case where no paths are found
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format\n\nCouldn't find a path from start to fisish")
		os.Exit(0)
	}
	return RemoveRepetition(paths)
}

// Backtrack is a function that explores all possible paths from the starting room to the end room.
func Backtrack(c *static.Colony) [][]string {
	// NOTE : append is necessary because of how slices are handled in golang
	// paths stores the valid paths
	// stack stores the diffrent branches to allow back tracking

	// Initialize the list of paths and the stack to hold the current path
	paths := [][]string{}
	stack := [][]string{{c.Start}}
	for len(stack) > 0 {

		path := append([]string{}, stack[len(stack)-1]...) // current path being checked
		stack = stack[:len(stack)-1]                       // Remove the curren path from the stack
		current := path[len(path)-1]

		if current == c.Finish {
			paths = append(paths, append([]string{}, path...))
			continue
		}
		// Add all the valid neighboring rooms to the stack.
		// Valid rooms are those that are'nt in the path
		for _, link := range (*c).Rooms[current] {
			if !static.IsIn(link, path) {
				stack = append(stack, append(append([]string{}, path...), link)) // Push the new path onto the stack

			}
		}
	}
	return paths
}

// RemoveRepetition removes repeated paths from the input slice of paths.
func RemoveRepetition(paths [][]string) [][]string {
	// gr stores the indexes of the paths that are compatible with each other
	gr := [][]int{}
	for i := 0; i < len(paths); i++ {
		c0 := []string{}
		c1 := []int{i}
		static.Diffrent(&paths[i], &c0)
		for i1, v := range paths {
			if i1 != i && static.Diffrent(&v, &c0) {
				c1 = append(c1, i1)
			}
		}
		sort.SliceStable(c1, func(i, j int) bool {
			return i < j // Compare lengths of inner slices
		})
		gr = append(gr, c1)
	}
	fmt.Println(gr, paths)
	sort.Slice(gr, func(i, j int) bool {
		if len(gr[i]) != len(gr[j]) {
			return len(gr[i]) > len(gr[j])
		}
		// If lengths are the same, compare by the lenght of the path
		for v := range gr[i] {
			if len(paths[gr[i][v]]) != len(paths[gr[j][v]]) {
				return len(paths[gr[i][v]]) < len(paths[gr[j][v]])
			}
		}
		return false
	})
	res := [][]string{}
	for _, i := range gr[0] {
		res = append(res, paths[i])
	}
	return res
}
