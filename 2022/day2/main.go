package main

import (
	"fmt"

	"github.com/mattboll/aoc/lib"
)

func main() {
	part1()
	part2()
}

func getWinner(m string, l string) int {
		if l == "X" {
			if m == "A" {
				return 3
			}
			if m == "B" {
				return 0
			}
			if m == "C" {
				return 6
			}
		}
		if l == "Y" {
			if m == "A" {
				return 6
			}
			if m == "B" {
				return 3
			}
			if m == "C" {
				return 0
			}
		}
		if l == "Z" {
			if m == "A" {
				return 0
			}
			if m == "B" {
				return 6
			}
			if m == "C" {
				return 3
			}
		}
		return 0
}

func getScore(m string, l string) int {
		if l == "X" {
			if m == "A" {
				return 3
			}
			if m == "B" {
				return 1
			}
			if m == "C" {
				return 2
			}
		}
		if l == "Y" {
			if m == "A" {
				return 3+1
			}
			if m == "B" {
				return 3+2
			}
			if m == "C" {
				return 3+3
			}
		}
		if l == "Z" {
			if m == "A" {
				return 6+2
			}
			if m == "B" {
				return 6+3
			}
			if m == "C" {
				return 6+1
			}
		}
		return 0
}

func part1() {
	strLines := lib.ReadFile()
	score := 0
	for _, line := range strLines {
		if line == "" {
			break
		}
		m := line[:1]
		l := line[2:]

		if l == "X" {
			score += 1
		}
		if l == "Y" {
			score += 2
		}
		if l == "Z" {
			score += 3
		}
		score += getWinner(m, l)
	}
		fmt.Printf("%d\n", score)

}

func part2() {
	strLines := lib.ReadFile()
	score := 0
	for _, line := range strLines {
		if line == "" {
			break
		}
		m := line[:1]
		l := line[2:]
		score += getScore(m, l)
	}
		fmt.Printf("%d\n", score)
}
