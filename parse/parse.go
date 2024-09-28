package parse

import (
	"bytes"
	"fmt"
	"log"
	"os"
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
	Links    []Link
}
type Link string

func Parse(file_name string) map[string]Room {
	file, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	colony := make(map[string]Room)
	data := bytes.Split(file, []byte("\n"))
	var s, f bool
	for _, line := range data {
		var r Room
		if string(line) == start {
			s = true
			continue
		}
		if string(line) == end {
			f = true
			continue
		}
		// line = line[:len(line)-1] // Remove trailing newline
		// remove comments
		// fmt.Println(line, string(line), "trimer")
		comments(&line)
		if len(line) == 0 {
			continue
		}
		// fmt.Println(line)
		name := ""
		n, err := fmt.Sscanf(string(line), "%s %d %d", &name, &r.X, &r.Y)
		// fmt.Println(n, err, line, string(line), "lwla")
		if err == nil && n == 3 {
			switch {
			case s:
				r.Fullness = 1.0
				s = false
				break
			case f:
				r.Fullness = 2.0
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
		// fmt.Println(line, string(line))
		name, n2 = l[0], l[1]
		n = len(l)
		// fmt.Println(n, line, "  tania", name, "azer", n2)
		if n != 2 {
			continue
		}

		// Ensure both rooms exist in the map
		if r, exists := colony[name]; exists {
			r.Links = append(r.Links, Link(n2))
			colony[name] = r // Update room with new link
		}
		if r, exists := colony[n2]; exists {
			r.Links = append(r.Links, Link(name))
			colony[n2] = r // Update room with new link
		}
	}
	return colony
}

func comments(line *[]byte) {
	for i, j := range *line {
		if j == '#' {
			*line = (*line)[:i]
			break
		}
	}
}
