package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mattboll/aoc/lib"
)

func main() {
	part1()
	part2()
}

func part1() {
	strLines := lib.ReadFile()
	s := 0
	for i, line := range strLines {
		ok := true
		games := strings.Split(line, ":")
		t := strings.Split(games[1], ";")
		for _, v := range t {
			c := strings.Split(string(v), ",")
			for _, vv := range c {
				vvv := strings.Fields(vv)
				n, _ := strconv.Atoi(vvv[0])
				switch vvv[1] {
				case "red":
					if n > 12 {
						ok = false
						break
					}
				case "green":
					if n > 13 {
						ok = false
						break
					}
				case "blue":
					if n > 14 {
						ok = false
						break
					}
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			s += i + 1
		}
	}
	fmt.Printf("%d\n", s)
}

func part2() {
	strLines := lib.ReadFile()
	s := 0
	for _, line := range strLines {
		games := strings.Split(line, ":")
		t := strings.Split(games[1], ";")
		r, g, b := 0, 0, 0
		for _, v := range t {
			c := strings.Split(v, ",")
			for _, vv := range c {
				vvv := strings.Fields(vv)
				n, _ := strconv.Atoi(vvv[0])
				switch vvv[1] {
				case "red":
					r = max(r, n)
				case "green":
					g = max(g, n)
				case "blue":
					b = max(b, n)
				}
			}
		}
		s += r * g * b
	}
	fmt.Println(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// part1:  3059
// part2:  65371
