package parse

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	start = "##start"
	end   = "##end"
)

type Room struct {
	X, Y     int
	Fullness float64
	Empty    bool
	Links    map[string]Link
}

type Link struct {
	P bool
	R Room
}

type Colony struct {
	Rooms         map[string]Room
	Strat, Finish string
	Ants          int
}

func Parse(file_name string) Colony {
	file, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	colony := make(map[string]Room)
	data := bytes.Split(file, []byte("\n"))
	var s, f bool
	var s1, f1 string
	ants := 0
	for _, line := range data {
		var r Room
		r.Links = make(map[string]Link)
		if string(line) == start {
			s = true
			continue
		}
		if string(line) == end {
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
		}
		name := ""
		n, err := fmt.Sscanf(string(line), "%s %d %d", &name, &r.X, &r.Y)
		if err == nil && n == 3 {
			switch {
			case s:
				r.Fullness = 1.0
				s1 = name
				s = false
				break
			case f:
				r.Fullness = 2.0
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
		n = len(l)
		if n != 2 {
			continue
		}

		// Ensure both rooms exist in the map
		if r, exists := colony[name]; exists {
			r.Links[n2] = Link{P: false, R: (colony[n2])}
			colony[name] = r // Update room with new link
		}
		if r, exists := colony[n2]; exists {
			r.Links[name] = Link{P: false, R: (colony[name])}
			colony[n2] = r // Update room with new link
		}
	}
	fmt.Println("azer", ants)
	return Colony{colony, s1, f1, ants}
}

func comments(line *[]byte) {
	for i, j := range *line {
		if j == '#' {
			*line = (*line)[:i]
			break
		}
	}
}
