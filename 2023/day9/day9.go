package main

import (
	"fmt"
	"strings"

	"github.com/mattboll/aoc/lib"
)

func main() {
	lib.RunAndMeasure(part1)
	lib.RunAndMeasure(part2)
}

func computeValue(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := []int{}
	for i := 1; i < len(nums); i++ {
		l = append(l, nums[i]-nums[i-1])
	}
	for i := 1; i < len(l); i++ {
		if l[i-1] != l[i] {
			c := computeValue(l)
			return nums[len(nums)-1] + c
		}
	}
	return nums[len(nums)-1] + l[0]
}

func part1() {
	strLines := lib.ReadFile()
	s := 0

	for _, line := range strLines {
		snums := strings.Fields(line)
		nums := []int{}
		for _, v := range snums {
			nums = append(nums, lib.Atoi(v))
		}
		x := computeValue(nums)
		s += x
	}

	fmt.Println(s)
}

func computePreviousValue(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := []int{}
	for i := 1; i < len(nums); i++ {
		l = append(l, nums[i]-nums[i-1])
	}
	for i := 1; i < len(l); i++ {
		if l[i-1] != l[i] {
			c := computePreviousValue(l)
			return nums[0] - c
		}
	}
	return nums[0] - l[0]
}

func part2() {
	strLines := lib.ReadFile()
	s := 0

	for _, line := range strLines {
		snums := strings.Fields(line)
		nums := []int{}
		for _, v := range snums {
			nums = append(nums, lib.Atoi(v))
		}
		x := computePreviousValue(nums)
		s += x
	}

	fmt.Println(s)
}
