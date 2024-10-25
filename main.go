package main

import (
	"fmt"
	"os"

	"lemin/brain"
	"lemin/parse"
	"lemin/pathing"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: lemin <file>")
		os.Exit(0)
	}
	colony := parse.Parse(os.Args[1])
	paths := pathing.FindAllPaths(colony)
	brain.Lemin(colony.Ants, paths)
}
