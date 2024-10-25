package pathing

import (
	"container/list"
	"fmt"
	"os"
	"sort"

	"lemin/static"
)

// path reperesents a node in our path with its history
type PathNode struct {
	current string
	path    []string
}

// this func findes all valid paths in the colony and returns them
func FindAllPaths(colony *static.Colony) [][]string {
	paths := BFS(colony)
    // sort paths by length
    sort.SliceStable(paths, func(i,j int)bool{
        return len(paths[i]) < len(paths[j]) 
    })
    // Handle no paths case
    if len(paths) == 0 {
        fmt.Println("error: no path from start to finish")
        os.Exit(1)
    }
    return RemoveRepetition(paths)
}

// bfs performes a bla bla
func BFS(c *static.Colony) [][]string {
	var paths [][]string
	queue := list.New()
	visited := make(map[string]bool)

	// start the bfs with the start room
	start := PathNode{current: c.Start, path: []string{c.Start}}
	queue.PushBack(start)
	shortestLen := -1

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(PathNode)
		// stop if current path exceeds allowed length
		if shortestLen != -1 && len(node.path) > shortestLen+3 {
			continue
		}
		// if the end room is reached, add the path
		if node.current == c.Finish {
			paths = append(paths, node.path)
			if shortestLen == -1 {
				shortestLen = len(node.path)
			}
			continue
		}
		// Explore the next rooms
		for _, next := range c.Rooms[node.current] {
			if static.IsIn(next, node.path) {
				continue
			}
			// create the new path
			newPath := append(append([]string{}, node.path...), next)
			statekey := next + "-" + fmt.Sprint(newPath)
			if !visited[statekey] {
				visited[statekey] = true
				queue.PushBack(PathNode{current: next, path: newPath})
			}
		}
	}
	return paths
}

// RemoveRepetition removes redundant paths
func RemoveRepetition(paths [][]string) [][]string{
    groups := [][]int{}
    for i := range paths{
        differnces := []string{}
        group := []int{i}
        static.Diffrent(&paths[i], &differnces)

        for j, otherPath := range paths{
            if i != j && static.Diffrent(&otherPath, &differnces){
                group = append(group, j)
            }
        }
        sort.Ints(group)
        groups = append(groups, group)
    }
    // sort groups by size and path length
    sort.Slice(groups, func(i,j int)bool{
        if len(groups[i]) != len(groups[j]){
            return len(groups[i]) > len(groups[j])
        }
		return len(paths[groups[i][0]]) < len(paths[groups[j][0]])
    })

    // return the best group of paths
    result := [][]string{}
    for _,i := range groups[0]{
        result = append(result, paths[i])
    }
    return result
} 