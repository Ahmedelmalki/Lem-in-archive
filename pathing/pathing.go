package pathing

import (
	"fmt"
	"lemin/static"
	"os"
	"sort"
)

func FindAllPath(colony *static.Colony) [][]string {
	visited := make(map[string]bool)
	paths := make([][]string, 0)

	// Start the backtracking process from the start room
	Backtrack(colony, colony.Start, visited, []string{colony.Start}, &paths)
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

// Backtrack is a recursive function that explores all possible paths from the current room to the end room.
func Backtrack(c *static.Colony, current string, visited map[string]bool, path []string, paths *[][]string) {
	if current == c.Finish {
		*paths = append(*paths, append([]string{}, path...))
		return
	}
	visited[current] = true
	for _, link := range (*c).Rooms[current] {
		if !visited[link] {
			Backtrack(c, link, visited, append(path, link), paths) // Add the link to the path before making the recursive call
		}
	}
	visited[current] = false
}

// RemoveRepetition removes repeated paths from the input slice of paths.
func RemoveRepetition(paths [][]string) [][]string {
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
		gr = append(gr, c1)
	}
	sort.Slice(gr, func(i, j int) bool {
		if len(gr[i]) != len(gr[j]) {
			return len(gr[i]) > len(gr[j])
		}
		// If lengths are the same, compare by the first element
		for v := range gr[i] {
			if gr[i][v] != gr[j][v] {
				return gr[i][v] < gr[j][v]
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
