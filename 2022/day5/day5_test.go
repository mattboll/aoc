package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/mattboll/aoc/lib"
)

func TestDay4Task1(t *testing.T) {
	lib.RunAndMeasure(part1)
}

func TestDay4Task2(t *testing.T) {
	lib.RunAndMeasure(part2)
}

func part1() {
	strLines := lib.ReadFile()
	start := true
	crates := make(map[int][]string)
	first := true
	for _, line := range strLines {
		if strings.Contains(line, " 1   2   3") {
			start = false
			continue
		}
		if line == "" {
			continue
		}
		if start {
			for j := 0; j < len(line); j++ {
				if line[j] >= 'A' && line[j] <= 'Z' {
					crates[j/4] = append(crates[j/4], string(line[j]))
				}
			}
		} else {
			if first {
				for _, v := range crates {
					for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
						v[i], v[j] = v[j], v[i]
					}
				}
				first = false
			}
			re := regexp.MustCompile("[0-9]+")
			f := re.FindAllString(line, -1)
			n, _ := strconv.Atoi(f[0])
			from, _ := strconv.Atoi(f[1])
			to, _ := strconv.Atoi(f[2])
			from--
			to--
			for k := 0; k < n; k++ {
				crates[to] = append(crates[to], crates[from][len(crates[from])-1])
				crates[from] = crates[from][:len(crates[from])-1]
			}
		}
	}
	s := ""
	for i := 0; i < 9; i++ {
		s += crates[i][len(crates[i])-1]
	}
	fmt.Println(s)
}

func part2() {
	strLines := lib.ReadFile()
	start := true
	crates := make(map[int][]string)
	first := true
	for _, line := range strLines {
		if strings.Contains(line, " 1   2   3") {
			start = false
			continue
		}
		if line == "" {
			continue
		}
		if start {
			for j := 0; j < len(line); j++ {
				if line[j] >= 'A' && line[j] <= 'Z' {
					crates[j/4] = append(crates[j/4], string(line[j]))
				}
			}
		} else {
			if first {
				for _, v := range crates {
					for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
						v[i], v[j] = v[j], v[i]
					}
				}
				first = false
			}
			re := regexp.MustCompile("[0-9]+")
			f := re.FindAllString(line, -1)
			n, _ := strconv.Atoi(f[0])
			from, _ := strconv.Atoi(f[1])
			to, _ := strconv.Atoi(f[2])
			from--
			to--
			crates[to] = append(crates[to], crates[from][len(crates[from])-n:]...)
			crates[from] = crates[from][:len(crates[from])-n]
		}
	}
	s := ""
	for i := 0; i < 9; i++ {
		s += crates[i][len(crates[i])-1]
	}
	fmt.Println(s)

}
