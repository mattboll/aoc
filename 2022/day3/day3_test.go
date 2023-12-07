package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mattboll/aoc/lib"
)

func TestDay3Task1(t *testing.T) {
	lib.RunAndMeasure(part1)
}

func TestDay3Task2(t *testing.T) {
	lib.RunAndMeasure(part2)
}

func part1() {
	strLines := lib.ReadFile()
	s := 0
	for _, line := range strLines {
		r := getSame(line[0:len(line)/2], line[len(line)/2:])
		s += getValue(r)
	}
	fmt.Println(s)
}

func getValue(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}
	return int(r - 'A' + 27)
}

func getSame(s1, s2 string) rune {
	for _, r := range s1 {
		if strings.Contains(s2, string(r)) {
			return r
		}
	}
	fmt.Println("ERROR")
	return 'a'
}

func getSame3(s1, s2, s3 string) rune {
	for _, r := range s1 {
		if strings.Contains(s2, string(r)) {
			if strings.Contains(s3, string(r)) {
				return r
			}
		}
	}
	fmt.Println("ERROR")
	return 'a'
}

func part2() {
	strLines := lib.ReadFile()
	s := 0
	for i := 0; i < len(strLines)-2; i += 3 {
		r := getSame3(strLines[i], strLines[i+1], strLines[i+2])
		s += getValue(r)
	}
	fmt.Println(s)
}
