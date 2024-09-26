package main

import (
	"fmt"
	"os"

	"lemin/parse"
)

// var colony = make(map[string]parse.Room)

func main() {
	colony := parse.Parse(os.Args[1])
	paths := path_finder.findAllPath(colony, "0", "1")
	fmt.Println("Shortest path:", paths)
	fmt.Println(colony)
}
