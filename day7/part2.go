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

func calcSingleDist(n1 int, n2 int) int {
	d := n2 - n1
	if d < 0 {
		d = -d
	}
	return (d * (d + 1)) / 2
}

func calcDist(crabs []int, d int) int {
	r := 0
	for _, v := range crabs {
		r += calcSingleDist(v, d)
	}
	return r
}

func main() {
	strLines := ReadFile()
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
