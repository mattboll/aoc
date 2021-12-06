package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile() []string {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fileContent := string(file)
	slicedContent := strings.Split(fileContent, "\n")
	return slicedContent
}

func main() {
	strLines := ReadFile()

	a := [1000][1000]int{}

	for _, line := range strLines {
		if line == "" {
			break
		}
		r1 := strings.SplitN(line, " -> ", -1)
		rs1 := strings.SplitN(r1[0], ",", -1)
		rs2 := strings.SplitN(r1[1], ",", -1)
		x1, _ := strconv.Atoi(rs1[0])
		y1, _ := strconv.Atoi(rs1[1])
		x2, _ := strconv.Atoi(rs2[0])
		y2, _ := strconv.Atoi(rs2[1])

		plusX := 1
		plusY := 1
		if x1 == x2 {
			plusX = 0
		}
		if x1 > x2 {
			plusX = -1
		}
		if y1 == y2 {
			plusY = 0
		}
		if y1 > y2 {
			plusY = -1
		}

		x := x1
		y := y1

		for {
			a[x][y] = a[x][y] + 1
			x = x + plusX
			y = y + plusY
			if x == x2 && y == y2 {
				a[x][y] = a[x][y] + 1
				break
			}
		}
	}
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if a[i][j] > 1 {
				count++
			}
		}
	}
	fmt.Printf("count : %d \n", count)
}
