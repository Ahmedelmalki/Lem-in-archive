package path_finder

import "lemin/parse"

func findAllPath(colony map[string]parse.Room, start, end string) [][]string {
	visited := make(map[string]bool)
	paths := make([][]string, 0)
	backtrack(colony, start, end, visited, []string{start}, &paths)
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
