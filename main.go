package main

import (
	"fmt"
	"os"

	"lemin/brain"
	"lemin/parse"
	"lemin/pathing"
)

func main() {
	colony := parse.Parse(os.Args[1])
	paths := pathing.FindAllPath(colony, "0", "1")
	fmt.Println("Shortest path:", paths)
	fmt.Println(colony)

	result := brain.Lemin(colony, paths)
	brain.DisplayResult(result)
}
