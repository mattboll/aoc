package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mattboll/aoc2021/lib"
)

func calcDist(crabs []int, d int) int {
	r := 0
	for _, v := range crabs {
		if v > d {
			r += v - d
		} else {
			r += d - v
		}
	}
	return r
}

func main() {
	strLines := lib.ReadFile()
	crabs := []int{}
	min := 1000
	max := 0
	for _, line := range strLines {
		if line == "" {
			break
		}
		r1 := strings.SplitN(line, ",", -1)
		for _, v := range r1 {
			n, _ := strconv.Atoi(v)
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
			crabs = append(crabs, n)
		}
	}

	resi := 0
	vali := 0
	for i := min; i <= max; i++ {
		c := calcDist(crabs, i)
		if resi > c || resi == 0 {
			resi = c
			vali = i
		}
	}
	fmt.Printf("i res : %d %d \n", vali, resi)
}
