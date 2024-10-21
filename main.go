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
	result := brain.Lemin(colony.Ants, paths)
	brain.DisplayResult(result)
}
