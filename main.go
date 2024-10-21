package main

import (
	"os"

	"lemin/brain"
	"lemin/parse"
	"lemin/pathing"
)

func main() {
	colony := parse.Parse(os.Args[1])
	paths := pathing.FindAllPath(colony)
	// fmt.Println("Shortest path:", paths)
	result := brain.Lemin(colony, paths)
	// fmt.Println("\nFinal room fullness:", result)
	brain.DisplayResult(result)
}
