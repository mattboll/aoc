package main

import (
	"fmt"
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
	s := 0
	for _, line := range strLines {
		p := strings.Split(line, ",")
		se1 := strings.Split(p[0], "-")
		se2 := strings.Split(p[1], "-")
		e1 := []int{0, 0}
		e2 := []int{0, 0}
		e1[0], _ = strconv.Atoi(se1[0])
		e1[1], _ = strconv.Atoi(se1[1])
		e2[0], _ = strconv.Atoi(se2[0])
		e2[1], _ = strconv.Atoi(se2[1])
		if e1[0] <= e2[0] && e1[1] >= e2[1] {
			s++
		} else if e2[0] <= e1[0] && e2[1] >= e1[1] {
			s++
		}
	}
	fmt.Println(s)
}

func part2() {
	strLines := lib.ReadFile()
	s := 0
	for _, line := range strLines {
		p := strings.Split(line, ",")
		se1 := strings.Split(p[0], "-")
		se2 := strings.Split(p[1], "-")
		e1 := []int{0, 0}
		e2 := []int{0, 0}
		e1[0], _ = strconv.Atoi(se1[0])
		e1[1], _ = strconv.Atoi(se1[1])
		e2[0], _ = strconv.Atoi(se2[0])
		e2[1], _ = strconv.Atoi(se2[1])
		if e1[0] <= e2[0] && e1[1] >= e2[1] {
			s++
		} else if e2[0] <= e1[0] && e2[1] >= e1[1] {
			s++
		} else if e1[0] <= e2[0] && e1[1] >= e2[0] {
			s++
		} else if e1[0] <= e2[1] && e1[1] >= e2[1] {
			s++
		}
	}
	fmt.Println(s)

}
