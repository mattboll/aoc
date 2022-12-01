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
		l := strings.SplitN(r[1], " ", -1)
		for _, v := range l {
			s := len(v)
			if s == 2 || s == 4 || s == 3 || s == 7 {
				k++
			}
		}
	}
	fmt.Printf("k : %d \n", k)
}
