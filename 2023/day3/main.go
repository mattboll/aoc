package main

import (
	"fmt"
	"strconv"
	"time"
	"unicode"

	"github.com/mattboll/aoc/lib"
)

func main() {
	fmt.Println("--2023 day 03 solution--")
	start := time.Now()
	part1()
	fmt.Println(time.Since(start))

	start = time.Now()
	part2()
	fmt.Println(time.Since(start))
}

func part1() {
	strLines := lib.ReadFile()
	s := 0
	a := make([][]rune, len(strLines))

	for i, line := range strLines {
		a[i] = []rune(line)
	}

	for i, line := range a {
		number := ""
		for j, r := range line {
			if unicode.IsDigit(r) {
				number += string(r)
			} else {
				if number != "" {
					n, _ := strconv.Atoi(number)
					if nearSymbol(i, j, number, a) {
						s += n
					}
				}
				number = ""
			}
		}

		if number != "" {
			n, _ := strconv.Atoi(number)
			if nearSymbol(i, len(a[i]), number, a) {
				s += n
			}
		}
	}

	fmt.Printf("%d\n", s)
}
func nearSymbol(i int, j int, number string, a [][]rune) bool {
	if i > 0 {
		for y := lib.Max(j-len(number)-1, 0); y <= lib.Min(j, len(a[i])-1); y++ {
			r := a[i-1][y]
			if r != '.' && r != '0' && r != '1' && r != '2' && r != '3' && r != '4' && r != '5' && r != '6' && r != '7' && r != '8' && r != '9' {
				return true
			}
		}
	}
	if j-len(number) > 0 {
		r := a[i][j-len(number)-1]
		if r != '.' && r != '0' && r != '1' && r != '2' && r != '3' && r != '4' && r != '5' && r != '6' && r != '7' && r != '8' && r != '9' {
			return true
		}
	}
	if j+1 < len(a[i]) {
		r := a[i][j]
		if r != '.' && r != '0' && r != '1' && r != '2' && r != '3' && r != '4' && r != '5' && r != '6' && r != '7' && r != '8' && r != '9' {
			return true
		}
	}
	if i < len(a)-2 {
		for y := lib.Max(j-len(number)-1, 0); y <= lib.Min(j, len(a[i])-1); y++ {
			r := a[i+1][y]
			if r != '.' && r != '0' && r != '1' && r != '2' && r != '3' && r != '4' && r != '5' && r != '6' && r != '7' && r != '8' && r != '9' {
				return true
			}
		}
	}
	return false
}

// ------------------

// P1 P2 P3
// P4 *  P5
// P6 P7 P8
func getAdjacentNumbers(i int, j int, a [][]rune) []int {
	r := []int{}
	p2 := false
	p3 := false
	p7 := false
	p8 := false

	if i > 0 {
		s := ""
		// test P1
		if j > 1 {
			for x := j - 1; x >= 0; x-- {
				if unicode.IsDigit(a[i-1][x]) {
					s = string(a[i-1][x]) + s
				} else {
					break
				}
			}
			for x := j; x < len(a[i-1]); x++ {
				if unicode.IsDigit(a[i-1][x]) {
					if x == j {
						p2 = true
					}
					if x == j+1 {
						p3 = true
					}
					s = s + string(a[i-1][x])
				} else {
					break
				}
			}

			if s != "" {
				n, _ := strconv.Atoi(s)
				r = append(r, n)
			}
		}

		if !p2 {
			s := ""
			for x := j; x < len(a[i-1]); x++ {
				if unicode.IsDigit(a[i-1][x]) {
					if x == j+1 {
						p3 = true
					}
					s = s + string(a[i-1][x])
				} else {
					break
				}
			}

			if s != "" {
				n, _ := strconv.Atoi(s)
				r = append(r, n)
			}

		}
		if !p3 {
			s := ""
			for x := j + 1; x < len(a[i-1]); x++ {
				if unicode.IsDigit(a[i-1][x]) {
					s = s + string(a[i-1][x])
				} else {
					break
				}
			}

			if s != "" {
				n, _ := strconv.Atoi(s)
				r = append(r, n)
			}
		}
	}
	// test P4
	s := ""
	for x := j - 1; x >= 0; x-- {
		if unicode.IsDigit(a[i][x]) {
			s = string(a[i][x]) + s
		} else {
			break
		}
	}
	if s != "" {
		n, _ := strconv.Atoi(s)
		r = append(r, n)
	}
	// test P5
	s = ""
	for x := j + 1; x < len(a[i]); x++ {
		if unicode.IsDigit(a[i][x]) {
			s = s + string(a[i][x])
		} else {
			break
		}
	}
	if s != "" {
		n, _ := strconv.Atoi(s)
		r = append(r, n)
	}

	if i < len(a)-1 {
		s := ""
		// test P6
		if j > 1 {
			for x := j - 1; x >= 0; x-- {
				if unicode.IsDigit(a[i+1][x]) {
					s = string(a[i+1][x]) + s
				} else {
					break
				}
			}
			for x := j; x < len(a[i+1]); x++ {
				if unicode.IsDigit(a[i+1][x]) {
					if x == j {
						p7 = true
					}
					if x == j+1 {
						p8 = true
					}
					s = s + string(a[i+1][x])
				} else {
					break
				}
			}

			if s != "" {
				n, _ := strconv.Atoi(s)
				r = append(r, n)
			}
		}

		// test P7
		if !p7 {
			s := ""
			for x := j; x < len(a[i+1]); x++ {
				if unicode.IsDigit(a[i+1][x]) {
					if x == j+1 {
						p3 = true
					}
					s = s + string(a[i+1][x])
				} else {
					break
				}
			}

			if s != "" {
				n, _ := strconv.Atoi(s)
				r = append(r, n)
			}

		}
		// test P8
		if !p8 {
			s := ""
			for x := j + 1; x < len(a[i+1]); x++ {
				if unicode.IsDigit(a[i+1][x]) {
					s = s + string(a[i+1][x])
				} else {
					break
				}
			}

			if s != "" {
				n, _ := strconv.Atoi(s)
				r = append(r, n)
			}
		}
	}

	return r
}

func part2() {
	strLines := lib.ReadFile()
	s := 0
	a := make([][]rune, len(strLines))
	for i, line := range strLines {
		a[i] = []rune(line)
	}

	for i, line := range a {
		for j, r := range line {
			if r == '*' {
				nums := getAdjacentNumbers(i, j, a)
				if len(nums) == 1 {
					fmt.Println(nums[0])
				}
				if len(nums) == 2 {
					fmt.Println(nums[0], " ", nums[1])
					s += nums[0] * nums[1]
				}
			}
		}
	}

	fmt.Printf("%d\n", s)
}

// part1: 554003

// part2: 87263515
