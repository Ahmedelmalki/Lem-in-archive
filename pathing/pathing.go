package pathing

import (
	"sort"

	"lemin/parse"
)

func FindAllPath(colony parse.Colony) [][]string {
	visited := make(map[string]bool)
	paths := make([][]string, 0)
	backtrack(colony, colony.Strat, colony.Finish, visited, []string{colony.Strat}, &paths)
	// fmt.Println(paths)
	sort.SliceStable(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j]) // Compare lengths of inner slices
	})
	// fmt.Println(paths)
	return Wcov(paths)
	// return nil
}

func backtrack(colony parse.Colony, start, end string, visited map[string]bool, path []string, paths *[][]string) {
	if start == end {
		*paths = append(*paths, append([]string{}, path...))
		return
	}
	visited[start] = true
	for link := range colony.Rooms[start].Links {
		if !visited[link] {
			backtrack(colony, link, end, visited, append(path, link), paths)
		}
	}
	visited[start] = false
}

func Wcov(paths [][]string) [][]string {
	gr := [][]int{}
	for i := 0; i < len(paths); i++ {
		c0 := []string{}
		c1 := []int{}
		for i1, v := range paths {
			// fmt.Print("check ", i1, v)
			if i1 != i && Diffrent(&v, &c0) {
				// fmt.Print("pass ", i1, v)
				c1 = append(c1, i1)
				// res = append(res, v)
			}
		}
		// fmt.Println(c0)
		// fmt.Println("\n", i, c1, res)
		gr = append(gr, c1)
	}

	// fmt.Printf("\ngr: \t\t%v", gr)
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
	// fmt.Printf("\n\nres :\t\t%v\ngr: \t\t%v\n\n", res, gr)
	return res
}

func isIn[T comparable](val T, all []T) bool {
	for _, v := range all {
		// fmt.Println("checked in", val, all)
		if val == v {
			return true
		}
	}
	return false
}

// Check if paths are diffrent
// NOTE : it doesn't check the first and last elements
func Diffrent[T comparable](a1, a2 *[]T) bool {
	// fmt.Println("checkedazer          in", a1, a2)
	for _, v := range (*a1)[1 : len(*a1)-1] {
		if isIn(v, *a2) {
			return false
		}
	}
	// fmt.Println()
	// fmt.Println((*a1)[1 : len(*a1)-1])
	*a2 = append(*a2, (*a1)[1:len(*a1)-1]...)
	return true
}

// [[0 2 3 1] [0 2 4 1] [0 6 4 1] [0 6 4 2 3 1]]
// [[0 2 4 1] [0 2 3 1] [0 6 4 1] [0 6 4 2 3 1]]
