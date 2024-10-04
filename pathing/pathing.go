package pathing

import (
	"lemin/parse"
)

func FindAllPath(colony parse.Colony) [][]string {
	visited := make(map[string]bool)
	paths := make([][]string, 0)
	backtrack(colony, colony.Strat, colony.Finish, visited, []string{colony.Strat}, &paths)
	for i := 0; i < len(paths)-1; i++ {
		if len(paths[i]) < len(paths[i+1]) {
			paths[i], paths[i+1] = paths[i+1], paths[i]
			i = 0
		}
	}
	// fmt.Println(paths)
	return paths
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
