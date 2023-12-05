package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/mattboll/aoc/lib"
)

func main() {
	fmt.Println("--2023 day 04 solution--")
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

	for _, line := range strLines {
		winNums := []int{}
		c := strings.Split(line, ":")
		if len(c) != 2 {
			break
		}
		nums := strings.Split(c[1], "|")
		winners := strings.Fields(nums[0])
		for _, v := range winners {
			n, _ := strconv.Atoi(v)
			winNums = append(winNums, n)
		}
		havingNumbers := strings.Fields(nums[1])
		ts := 0
		for _, v := range havingNumbers {
			n, _ := strconv.Atoi(v)
			if slices.Contains(winNums, n) {
				if ts == 0 {
					ts = 1
				} else {
					ts = ts * 2
				}
			}
		}
		s += ts
	}
	fmt.Printf("%d\n", s)
}

func part2() {
	strLines := lib.ReadFile()
	cards := make([]int, 197)

	for i, line := range strLines {
		winNums := []int{}
		c := strings.Split(line, ":")
		if len(c) != 2 {
			break
		}
		nums := strings.Split(c[1], "|")
		winners := strings.Fields(nums[0])
		for _, v := range winners {
			n, _ := strconv.Atoi(v)
			winNums = append(winNums, n)
		}
		havingNumbers := strings.Fields(nums[1])
		ts := 0
		for _, v := range havingNumbers {
			n, _ := strconv.Atoi(v)
			if slices.Contains(winNums, n) {
				ts++
			}
		}
		for j := 0; j < ts; j++ {
			cards[i+j+1] = cards[i+j+1] + cards[i] + 1
		}
	}
	sum := 0
	for _, num := range cards {
		sum += num + 1
	}

	fmt.Printf("%d\n", sum)
}

// part1: 22193

// part2:
