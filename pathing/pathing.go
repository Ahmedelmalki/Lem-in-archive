package pathing

import (
	"fmt"
	"os"
	"sort"

	"lemin/parse"
)

var (
	c   *parse.Colony
	end string
)

func FindAllPath(colony *parse.Colony) [][]string {
	visited := make(map[string]bool)
	paths := make([][]string, 0)
	c = colony
	end = colony.Finish

	// Start the backtracking process from the start room
	Backtrack(colony.Strat, visited, []string{colony.Strat}, &paths)

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
func Backtrack(current string, visited map[string]bool, path []string, paths *[][]string) {
	if current == end {
		*paths = append(*paths, append([]string{}, path...))
		return
	}
	visited[current] = true
	for link := range (*c).Rooms[current].Links {
		if !visited[link] {
			Backtrack(link, visited, append(path, link), paths) // Add the link to the path before making the recursive call
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
		Diffrent(&paths[i], &c0)
		for i1, v := range paths {
			if i1 != i && Diffrent(&v, &c0) {
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

// Check if an element is in a slice of any comparable type
func isIn[T comparable](val T, all []T) bool {
	for _, v := range all {
		if val == v {
			return true
		}
	}
	return false
}

// Check if paths are diffrent
// NOTE : it doesn't check the first and last elements
func Diffrent[T comparable](a1, a2 *[]T) bool {
	for _, v := range (*a1)[1 : len(*a1)-1] {
		if isIn(v, *a2) {
			return false
		}
	}
	*a2 = append(*a2, (*a1)[1:len(*a1)-1]...)
	return true
}
