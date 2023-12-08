package main

import (
	"fmt"
	"regexp"

	"github.com/mattboll/aoc/lib"
)

func main() {
	lib.RunAndMeasure(part1)
	lib.RunAndMeasure(part2)
}

type Move struct {
	Libelle string
	Left    string
	Right   string
}

func findPos(s string, m []Move) int {
	for i, v := range m {
		if v.Libelle == s {
			return i
		}
	}
	return 0
}

func part1() {
	strLines := lib.ReadFile()
	instructions := strLines[0]
	re := regexp.MustCompile("[A-Z]+")
	mappings := []Move{}
	s := 0

	for i := 2; i < len(strLines); i++ {
		f := re.FindAllString(strLines[i], -1)
		m := Move{
			Libelle: f[0],
			Left:    f[1],
			Right:   f[2],
		}
		mappings = append(mappings, m)
	}

	finish := false
	// p := findPos("AAA", mappings)
	p := 189
	for !finish {
		for _, r := range instructions {
			if mappings[p].Libelle[2] == 'Z' {
				finish = true
				break
			}
			if r == 'L' {
				p = findPos(mappings[p].Left, mappings)
			} else if r == 'R' {
				p = findPos(mappings[p].Right, mappings)
			}
			s++
		}
	}

	fmt.Println(s)
}

func findPosList(s rune, m []Move) []int {
	pos := []int{}
	for i, v := range m {
		if rune(v.Libelle[2]) == s {
			pos = append(pos, i)
		}
	}
	return pos
}

func exists(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part2() {
	strLines := lib.ReadFile()
	instructions := strLines[0]
	re := regexp.MustCompile("[A-Z]+")
	mappings := []Move{}

	for i := 2; i < len(strLines); i++ {
		f := re.FindAllString(strLines[i], -1)
		m := Move{
			Libelle: f[0],
			Left:    f[1],
			Right:   f[2],
		}
		mappings = append(mappings, m)
	}

	p := findPosList('A', mappings)
	ps := []int{}
	for i := 0; i < len(p); i++ {
		s := 0
		finish := false
		for !finish {
			for _, r := range instructions {
				if mappings[p[i]].Libelle[2] == 'Z' {
					finish = true
					break
				}
				if r == 'L' {
					p[i] = findPos(mappings[p[i]].Left, mappings)
				} else if r == 'R' {
					p[i] = findPos(mappings[p[i]].Right, mappings)
				}
				s++
			}
		}
		ps = append(ps, s)
	}
	fmt.Println(ps)
	d := LCM(ps[0], ps[1], ps[2:]...)
	fmt.Println(d)

}
