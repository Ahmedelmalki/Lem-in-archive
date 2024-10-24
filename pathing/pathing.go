package pathing

import (
	"container/list"
	"fmt"
	"lemin/static"
	"os"
	"sort"
)

// PathNode represents a node in our path with its history
type PathNode struct {
	current string
	path    []string
}

func FindAllPath(colony *static.Colony) [][]string {
	// Start the BFS process from the start room
	paths := BFS(colony)
	
	// Sort the paths based on length
	sort.SliceStable(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	// Handle the case where no paths are found
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format\n\nCouldn't find a path from start to finish")
		os.Exit(0)
	}
	return RemoveRepetition(paths)
}

// BFS implements breadth-first search to find all paths from start to end
func BFS(c *static.Colony) [][]string {
	// Initialize result slice to store valid paths
	paths := [][]string{}
	
	// Create a queue for BFS
	queue := list.New()
	
	// Create a map to track visited states (room + path history)
	// This prevents exploring the same state multiple times
	visited := make(map[string]bool)
	
	// Start with the initial room
	start := PathNode{
		current: c.Start,
		path:    []string{c.Start},
	}
	queue.PushBack(start)
	
	// Track the shortest path length we've found
	shortestLen := -1
	
	for queue.Len() > 0 {
		// Get the next node from the queue
		element := queue.Front()
		queue.Remove(element)
		node := element.Value.(PathNode)
		
		// Skip if we've gone beyond the shortest path length + threshold
		if shortestLen != -1 && len(node.path) > shortestLen+3 {
			continue
		}
		
		// If we've reached the end room
		if node.current == c.Finish {
			paths = append(paths, node.path)
			// Update shortest path length if this is the first path found
			if shortestLen == -1 {
				shortestLen = len(node.path)
			}
			continue
		}
		
		// Explore all connected rooms
		for _, next := range c.Rooms[node.current] {
			// Skip if we've already visited this room in this path
			if static.IsIn(next, node.path) {
				continue
			}
			
			// Create new path by appending the next room
			newPath := make([]string, len(node.path))
			copy(newPath, node.path)
			newPath = append(newPath, next)
			
			// Create state key for visited check
			stateKey := fmt.Sprintf("%s-%v", next, newPath)
			if !visited[stateKey] {
				visited[stateKey] = true
				queue.PushBack(PathNode{
					current: next,
					path:    newPath,
				})
			}
		}
	}
	
	return paths
}

// RemoveRepetition removes repeated paths from the input slice of paths
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
			return i < j
		})
		gr = append(gr, c1)
	}
	
	// Sort groups by size and path length
	sort.Slice(gr, func(i, j int) bool {
		if len(gr[i]) != len(gr[j]) {
			return len(gr[i]) > len(gr[j])
		}
		for v := range gr[i] {
			if len(paths[gr[i][v]]) != len(paths[gr[j][v]]) {
				return len(paths[gr[i][v]]) < len(paths[gr[j][v]])
			}
		}
		return false
	})
	
	// Return the best group of compatible paths
	res := [][]string{}
	for _, i := range gr[0] {
		res = append(res, paths[i])
	}
	return res
}