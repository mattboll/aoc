package main

import (
	"fmt"
	"strings"

	"github.com/mattboll/aoc2021/lib"
)

func main() {
	strLines := lib.ReadFile()
	k := 0
	for _, line := range strLines {
		if line == "" {
			break
		}
		r := strings.SplitN(line, " | ", -1)
		l := strings.SplitN(r[0], " ", -1)
		output := strings.SplitN(r[1], " ", -1)

		m := make([]string, 10)

		for _, v := range l {
			s := len(v)
			if s == 2 {
				m[1] = v
			}
			if s == 3 {
				m[7] = v
			}
			if s == 4 {
				m[4] = v
			}
			if s == 7 {
				m[8] = v
			}
		}
		for _, v := range l {
			s := len(v)
			if s == 6 {
				if test6(m, v) {
					m[6] = v
				} else if test9(m, v) {
					m[9] = v
				} else if test0(m, v) {
					m[0] = v
				}
			}
		}

		for _, v := range l {
			s := len(v)
			if s == 5 {
				if test3(m, v) {
					m[3] = v
				} else if test5(m, v) {
					m[5] = v
				} else {
					m[2] = v
				}
			}
		}

		// barre du haut : 7 - 1
		// en taille 6 : 0 6 9
		//  6 : si pas les 2 barres du 1
		//  9 : toutes les barres du 4
		//  0 : si y a pas toutes les barres du 4
		// en taille 5 : 2 3 5
		//  3 si y a toutes les barres du 1
		//  5 toutes les barres sont dans le 6
		//  2 toutes les barres sont dans le 9

		k += getVal(m, output[3])
		k += 10 * getVal(m, output[2])
		k += 100 * getVal(m, output[1])
		k += 1000 * getVal(m, output[0])

	}
	fmt.Printf("k : %d \n", k)
}

func test6(m []string, v string) bool {
	for _, c := range m[1] {
		if !strings.Contains(v, string(c)) {
			return true
		}
	}
	return false
}
func test9(m []string, v string) bool {
	for _, c := range m[4] {
		if !strings.Contains(v, string(c)) {
			return false
		}
	}
	return true
}
func test0(m []string, v string) bool {
	for _, c := range m[4] {
		if !strings.Contains(v, string(c)) {
			return true
		}
	}
	return false
}
func test3(m []string, v string) bool {
	for _, c := range m[1] {
		if !strings.Contains(v, string(c)) {
			return false
		}
	}
	return true
}
func test5(m []string, v string) bool {
	for _, c := range v {
		if !strings.Contains(m[6], string(c)) {
			return false
		}
	}
	return true
}

func getVal(m []string, v string) int {
	for i := 0; i < 10; i++ {
		found := true
		if len(v) != len(m[i]) {
			continue
		}
		for _, c := range v {
			if !strings.Contains(m[i], string(c)) {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return 0
}
