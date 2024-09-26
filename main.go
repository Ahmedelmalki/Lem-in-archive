package main

import (
	"fmt"
	"os"
	"text/template/parse"
)

// var colony = make(map[string]parse.Room)

func main() {
	colony := parse.Parse(os.Args[1])
	// c := map[string]lemin.Room{
	// 	"0": {0, 3, 0.0, []lemin.Link{"2"}},
	// 	"1": {8, 3, 0.0, []lemin.Link{"3"}},
	// 	"2": {2, 5, 0.0, []lemin.Link{"0", "3"}},
	// 	"3": {4, 0, 0.0, []lemin.Link{"2", "1"}},
	// }
	// if CompareMaps(c, colony) {
	// 	fmt.Println("Maps are equal")
	// }
	fmt.Println(colony)
	// fmt.Println(c)
}
