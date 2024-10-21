package parse

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"lemin/static"
)

const (
	start = "##start"
	end   = "##end"
)

func Parse(file_name string) *static.Colony {
	file, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	colony := make(map[string][]string)
	data := strings.Split(string(file), "\n")
	var s, f bool
	var s1, f1 string
	ants := 0
	for ln, line := range data {
		if line == start {
			s = true
			continue
		}
		if line == end {
			f = true
			continue
		}
		// remove comments
		comments(&line)
		if len(line) == 0 {
			continue
		}
		if ants == 0 {
			ants, err = strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
		if ants <= 0 && ln != 0 {
			fmt.Printf("ERROR: invalid data format\nProgram needs to start with value of 1 or more.\n PROBLEM AT LINE %d : %s", ln, line)
			os.Exit(1)
		}
		name := ""
		i := 0
		n, err := fmt.Sscanf(line, "%s %d %d", &name, &i, &i)
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
			colony[name] = []string{}
			continue
		}
		// Attempt to parse links
		n2 := ""
		l := strings.Split(line, "-")
		if len(l) != 2 {
			continue
		}
		name, n2 = l[0], l[1]

		// Ensure both rooms exist in the map and create links
		if r, exists := colony[name]; exists && !static.IsIn(n2, r) {
			colony[name] = append(r, n2) // Update room with new link
		} else if !exists {
			fmt.Printf("ERROR: invalid data format\n\n PROBLEM AT LINE %d  room '%s' Doesn't exist : %s", ln, name, line)
			os.Exit(1)
		}
		if r, exists := colony[n2]; exists && !static.IsIn(name, r) {
			colony[n2] = append(r, name) // Update room with new link
		} else if !exists {
			fmt.Printf("ERROR: invalid data format\n\n PROBLEM AT LINE %d room '%s' Doesn't exist : %s", ln, n2, line)
			os.Exit(1)
		}
	}
	return &static.Colony{Rooms: colony, Start: s1, Finish: f1, Ants: ants}
}

func comments(line *string) {
	for i, j := range *line {
		if j == '#' {
			*line = (*line)[:i]
			break
		}
	}
}
