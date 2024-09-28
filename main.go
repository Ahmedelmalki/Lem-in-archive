package main

import (
	"fmt"
	"os"

	"lemin/parse"
	"lemin/pathing"
)

// var colony = make(map[string]parse.Room)
// const n = 10
const ants = 5

func main() {
	colony := parse.Parse(os.Args[1])
	paths := pathing.FindAllPath(colony, "0", "1")
	fmt.Println("Shortest path:", paths)
	fmt.Println(colony)
	result := path(colony, paths)
	fmt.Println("\nTotal weight:", result)
	for i := 1; i < len(result[0]); i++ {
		fmt.Println()
		for _, v := range result {
			fmt.Print(v[i])
		}
	}
}

func path(c map[string]parse.Room, p [][]string) [][]string {
	for i := 0; i < len(p)-1; i++ {
		if len(p[i]) < len(p[i+1]) {
			p[i], p[i+1] = p[i+1], p[i]
			i = 0
		}
	}
	var s, f string
	for l, r := range c {
		if r.Fullness == 1.0 {
			s = l
			r.Fullness = 0.0
			c[l] = r
		} else if r.Fullness == 2.0 {
			f = l
			r.Fullness = 0.0
			c[l] = r
		}
	}

	for i := 0; i < len(p); i++ {
		for k, l := range p[i] {
			r := c[l]
			r.Fullness = float64((len(p[i][k:])) * 100.0)
			c[l] = r
		}
	}
	r := c[s]
	r.Fullness = 400.0
	c[s] = r
	for i := range c {
		fmt.Println(i, c[i])
	}
	fmt.Println(c, p, f)
	n := [ants]string{}
	x := make([][]string, ants)
	for i := range n {
		n[i] = s
		x[i] = []string{fmt.Sprintf("L%d-%s", i+1, s)}
	}
	for {
		for d := 1; d <= ants; d++ {
			free(&c, f)
			// fmt.Println(d, n[d-1], x)
			// if d == 5 {
			fmt.Println(n, x, "\n\n", c, p, r, d)
			// }
			if n[d-1] == f {
				x[d-1] = append(x[d-1], "")
				continue
			}
			r := c[n[d-1]]
			i := float64(len(p[0]) * 100.0)
			for _, link := range r.Links {
				r1 := c[string(link)]
				fmt.Println(i, r1)
				if i >= r1.Fullness && check(x[d-1], d, string(link)) && !r1.Empty || string(link) == f {
					i = r1.Fullness
				}
			}
			fmt.Println("test", d)
			for w, link := range r.Links {
				r1 := c[string(link)]

				if r1.Fullness == i && check(x[d-1], d, string(link)) && !r1.Empty {
					free(&c, n[d-1])
					// fmt.Println(d, n[d-1], x, link)
					r1.Fullness += 100.0
					r1.Empty = true
					c[string(link)] = r1
					n[d-1] = string(link)
					x[d-1] = append(x[d-1], fmt.Sprintf("L%d-%s", d, n[d-1]))
					break

				}
				if w == len(r.Links)-1 {
					x[d-1] = append(x[d-1], "")
				}
			}
			fmt.Println(d)
		}
		// fmt.Println(n, c, p)
		if n[ants-1] == f {
			break
		}
	}
	// fmt.Println(c, p)
	return x
}

func free(c *map[string]parse.Room, p string) {
	r := (*c)[p]
	r.Empty = false
	r.Fullness -= 100.0
	(*c)[p] = r
	// fmt.Println("\n\n\n", c)
}

func check(a []string, d int, b string) bool {
	fmt.Println("check               ", a, d, b)
	for _, c := range a {
		if c == fmt.Sprintf("L%d-%s", d, b) {
			fmt.Println("check          ---     ", a, d, b)
			return false
		}
	}
	return true
}
