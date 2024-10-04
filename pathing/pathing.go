package pathing

import "lemin/parse"

func FindAllPath(colony map[string]parse.Room, start, end string) [][]string {
	visited := make(map[string]bool)
	paths := make([][]string, 0)
	backtrack(colony, start, end, visited, []string{start}, &paths)
	for i := 0; i < len(paths)-1; i++ {
		if len(paths[i]) < len(paths[i+1]) {
			paths[i], paths[i+1] = paths[i+1], paths[i]
			i = 0
		}
	}
	return paths
}

func backtrack(colony map[string]parse.Room, start, end string, visited map[string]bool, path []string, paths *[][]string) {
	if start == end {
		*paths = append(*paths, append([]string{}, path...))
		return
	}
	visited[start] = true
	for _, link := range colony[start].Links {
		if !visited[string(link)] {
			backtrack(colony, string(link), end, visited, append(path, string(link)), paths)
		}
	}
	visited[start] = false
}
