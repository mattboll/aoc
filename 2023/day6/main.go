package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mattboll/aoc/lib"
)

func main() {
	fmt.Println("--2023 day 06 solution--")
	start := time.Now()
	part1()
	fmt.Println(time.Since(start))

	start = time.Now()
	part2()
	fmt.Println(time.Since(start))
}

type SeedMap struct {
	destination int
	source      int
	length      int
}

type SeedRange struct {
	start  int
	length int
}

func part1() {
	strLines := lib.ReadFile()
	var time []int
	var distance []int

	for _, line := range strLines {
		p := strings.Split(line, ":")
		if p[0] == "Time" {
			stime := strings.Fields(p[1])
			for s := range stime {
				n, _ := strconv.Atoi(stime[s])
				time = append(time, n)
			}
		} else {
			sdistance := strings.Fields(p[1])
			for s := range sdistance {
				n, _ := strconv.Atoi(sdistance[s])
				distance = append(distance, n)
			}
		}
	}

	var m []int
	for l := 0; l < len(time); l++ {
		c := 0
		for i := 0; i < time[l]; i++ {
			d := i * (time[l] - i)
			if d > distance[l] {
				c = c + 1
			}
		}
		m = append(m, c)
	}
	fmt.Print(m)
	mul := 1
	for i := 0; i < len(m); i++ {
		mul = mul * m[i]
	}

	fmt.Printf("%d\n", mul)
}

func part2() {
	strLines := lib.ReadFile()
	var time int
	var distance int

	for _, line := range strLines {
		p := strings.Split(line, ":")
		if p[0] == "Time" {
			stime := strings.Fields(p[1])
			longname := ""
			for _, s := range stime {
				longname = longname + s
			}
			n, _ := strconv.Atoi(longname)
			time = n
		} else {
			sdistance := strings.Fields(p[1])
			longname := ""
			for _, s := range sdistance {
				longname = longname + s
			}
			n, _ := strconv.Atoi(longname)
			distance = n
		}
	}

	c := 0
	for i := 0; i < time; i++ {
		d := i * (time - i)
		if d > distance {
			c = c + 1
		}
	}
	fmt.Printf("%d\n", c)
}
