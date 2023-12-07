package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/mattboll/aoc/lib"
)

func main() {
	fmt.Println("--2023 day 07 solution--")
	start := time.Now()
	part1()
	fmt.Println(time.Since(start))

	start = time.Now()
	part2()
	fmt.Println(time.Since(start))
}

type Hand struct {
	Cards []rune
	Bid   int
	Value int
}

var joker = false

func part1() {
	solve()
}

func solve() {
	strLines := lib.ReadFile()
	hands := []Hand{}

	for _, line := range strLines {
		f := strings.Fields(line)
		b, _ := strconv.Atoi(f[1])
		h := Hand{
			Cards: []rune(f[0]),
			Bid:   b,
		}
		hands = append(hands, h)
	}
	hands = orderHands(hands)
	s := 0
	for i := 1; i <= len(hands); i++ {
		s = s + i*hands[i-1].Bid
	}
	fmt.Printf("%d\n", s)
}

func getHandValue(s []rune) int {
	five := 0
	four := 0
	full := 0
	three := 0
	two := 0
	x := []bool{false, false, false, false, false}
	for i := 0; i < len(s); i++ {
		if x[i] {
			continue
		}
		x[i] = true
		n := 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				n++
				x[j] = true
			}
		}
		if n == 5 {
			five++
		} else if n == 4 {
			four++
		} else if n == 3 {
			if two > 0 {
				full++
				two--
			} else {
				three++
			}
		} else if n == 2 {
			if three > 0 {
				full++
				three--
			} else {
				two++
			}
		}
	}
	return two + three*3 + full*4 + four*5 + five*6
}

func getJokerHandValue(s []rune) int {
	five := 0
	four := 0
	full := 0
	three := 0
	two := 0
	x := []bool{false, false, false, false, false}
	nbJokers := 0
	for i := 0; i < len(s); i++ {
		if x[i] {
			continue
		}
		x[i] = true
		if s[i] == 'J' {
			nbJokers++
			continue
		}
		n := 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				n++
				x[j] = true
			}
		}
		if n == 5 {
			five++
		} else if n == 4 {
			four++
		} else if n == 3 {
			if two > 0 {
				full++
				two--
			} else {
				three++
			}
		} else if n == 2 {
			if three > 0 {
				full++
				three--
			} else {
				two++
			}
		}
	}
	if five > 0 {
		return 6
	}
	if four > 0 {
		if nbJokers > 0 {
			return 6
		}
		return 5
	}
	if full > 0 {
		return 4
	}

	if three > 0 {
		if nbJokers == 1 {
			return 5
		}
		if nbJokers == 2 {
			return 6
		}
		return 3
	}
	if two > 0 {
		if nbJokers == 3 {
			return 6
		}
		if nbJokers == 2 {
			return 5
		}
		if nbJokers == 1 {
			if two > 1 {
				return 4
			}
			return 3
		}
		return two
	}
	if nbJokers == 5 {
		return 6
	}
	if nbJokers == 4 {
		return 6
	}
	if nbJokers == 3 {
		return 5
	}
	if nbJokers == 2 {
		return 3
	}
	if nbJokers == 1 {
		return 1
	}

	return 0
}

func cardValue(s rune) int {
	switch s {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		if joker {
			return 1
		} else {
			return 11
		}
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	return 0
}

func orderHands(hands []Hand) []Hand {
	for i := 0; i < len(hands); i++ {
		if joker {
			hands[i].Value = getJokerHandValue(hands[i].Cards)
		} else {
			hands[i].Value = getHandValue(hands[i].Cards)
		}
	}

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].Value != hands[j].Value {
			return hands[i].Value < hands[j].Value
		}
		for k := 0; k < len(hands[i].Cards); k++ {
			a := cardValue(hands[i].Cards[k])
			b := cardValue(hands[j].Cards[k])
			if a != b {
				return a < b
			}
		}
		return false
	})
	return hands
}

func part2() {
	joker = true
	solve()
}
