package main

import (
	"fmt"
	"strconv"

	"github.com/mattboll/aoc2021/lib"
)

func main() {
	part1()
}

func part1() {
	strLines := lib.ReadFile()
	m := [][]int{}
	for _, line := range strLines {
		row := []int{}
		if line == "" {
			break
		}

		for _, v := range line {
			n, _ := strconv.Atoi(string(v))
			row = append(row, n)
		}

		m = append(m, row)
	}

	res := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if (i == 0 || m[i-1][j] > m[i][j]) && (j == 0 || m[i][j-1] > m[i][j]) && (i == len(m)-1 || m[i+1][j] > m[i][j]) && (j == len(m[i])-1 || m[i][j+1] > m[i][j]) {
				res = res + m[i][j] + 1
			}
		}
	}
	fmt.Printf("%d\n", res)
}
