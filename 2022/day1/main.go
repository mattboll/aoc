package main

import (
	"fmt"
	"strconv"
	"sort"

	"github.com/mattboll/aoc/lib"
)

func main() {
	part1()
	part2()
}

func part1() {
	strLines := lib.ReadFile()
	m := []int{}
	elf := 0
	m = append(m, 0)
	for _, line := range strLines {
		fmt.Printf("%v\n", m)
		if line == "" {
			elf += 1
			m = append(m, 0)
			continue
		}
		n, _ := strconv.Atoi(line)
		m[elf] += n
	}

	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	fmt.Printf("%d\n", max)
}

func part2() {
	strLines := lib.ReadFile()
	m := []int{}
	elf := 0
	m = append(m, 0)
	for _, line := range strLines {
		if line == "" {
			elf += 1
			m = append(m, 0)
			continue
		}
		n, _ := strconv.Atoi(line)
		m[elf] += n
	}
	fmt.Printf("%v\n", m)

	sort.Ints(m)
	sum := m[len(m)-1] +m[len(m)-2] +m[len(m)-3] 

	fmt.Printf("%v\n", sum)
}
