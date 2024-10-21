package parse

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"lemin/pathing"
)

const (
	start = "##start"
	end   = "##end"
)

type Colony struct {
	Rooms         map[string][]string
	Strat, Finish string
	Ants          int
}

func Parse(file_name string) *Colony {
	file, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	colony := make(map[string][]string)
	data := bytes.Split(file, []byte("\n"))
	var s, f bool
	var s1, f1 string
	ants := 0
	for _, line := range data {
		var r []string
		if string(line) == start {
			s = true
			continue
		}
		if string(line) == end {
			colony[name] = r // Update room with new link== end {
			f = true
			continue
		}
		// remove comments
		comments(&line)
		if len(line) == 0 {
			continue
		}
		if ants == 0 {
			ants, err = strconv.Atoi(string(line))
			if err != nil {
				log.Fatal(err)
			}
			if ants <= 0 {
				fmt.Println("ERROR: invalid data format\n\nneed one or more ants ")
				os.Exit(0)
			}
		}
		name := ""
		n, err := fmt.Sscanf(string(line), "%s %d %d", &name, &r.X, &r.Y)
		if err == nil && n == 3 {
			switch {
			case s:
				s1 = name
				s = false
				break
			case f:
				f1 = name
				f = false
				break
			}
			colony[name] = r
			continue
		}
		// Attempt to parse links
		n2 := ""
		l := strings.Split(string(line), "-")
		if len(l) != 2 {
			continue
		}
		name, n2 = l[0], l[1]

		// Ensure both rooms exist in the map
		if r, exists := colony[name]; exists && !pathing.IsIn(n2, r) {
			colony[name] = append(r, n2) // Update room with new link
		}
		if r, exists := colony[n2]; exists && !pathing.IsIn(name, r) {
			colony[n2] = append(r, name) // Update room with new link
		}
	}
	return &Colony{colony, s1, f1, ants}
}

func comments(line *[]byte) {
	for i, j := range *line {
		if j == '#' {
			*line = (*line)[:i]
			break
		}
	}
}
