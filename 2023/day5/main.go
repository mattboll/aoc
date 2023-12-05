package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mattboll/aoc/lib"
)

func main() {
	fmt.Println("--2023 day 05 solution--")
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
	seedList := []int{}
	soilMapping := make([][]SeedMap, 7)
	soilMappingPos := -1

	for _, line := range strLines {
		if strings.Contains(line, "seeds") {
			parts := strings.Split(line, ":")
			nums := strings.Fields(parts[1])
			for _, v := range nums {
				n, _ := strconv.Atoi(v)
				seedList = append(seedList, n)
			}
		} else {
			if len(line) > 3 {
				if strings.Contains(line, ":") {
					soilMappingPos++
				} else {
					if soilMapping[soilMappingPos] == nil {
						soilMapping[soilMappingPos] = make([]SeedMap, 0)
					}
					nums := strings.Fields(line)
					s := SeedMap{}
					s.destination, _ = strconv.Atoi(nums[0])
					s.source, _ = strconv.Atoi(nums[1])
					s.length, _ = strconv.Atoi(nums[2])
					soilMapping[soilMappingPos] = append(soilMapping[soilMappingPos], s)
				}
			}
		}
	}
	for i := 0; i < len(seedList); i++ {
		l, _ := getLocation(seedList[i], soilMapping)
		if minimalLocation > l {
			minimalLocation = l
		}
	}
	fmt.Printf("%d\n", minimalLocation)
}

func getLocation(seed int, soilMapping [][]SeedMap) (int, int) {
	r := seed
	minDist := math.MaxInt
	var m int
	for i := 0; i < len(soilMapping); i++ {
		r, m = getNewVal(r, soilMapping[i])
		if m < minDist {
			minDist = m
		}
	}

	return r, minDist
}

func getNewVal(seed int, seeds []SeedMap) (int, int) {
	for i := 0; i < len(seeds); i++ {
		v := seeds[i]
		if v.source <= seed && v.source+v.length > seed {
			r := v.destination + (seed - v.source)
			return r, (v.source + v.length) - seed
		}
	}
	return seed, math.MaxInt
}

var minimalLocation = math.MaxInt

func part2() {
	minimalLocation = math.MaxInt
	strLines := lib.ReadFile()
	seedRange := []SeedRange{}
	soilMapping := make([][]SeedMap, 7)
	soilMappingPos := -1

	for _, line := range strLines {
		if strings.Contains(line, "seeds") {
			parts := strings.Split(line, ":")
			nums := strings.Fields(parts[1])
			start := true
			s := SeedRange{}
			for _, v := range nums {
				n, _ := strconv.Atoi(v)
				if start {
					s = SeedRange{}
					s.start = n
					start = false
				} else {
					s.length = n
					seedRange = append(seedRange, s)
					start = true
				}
			}
		} else {
			if len(line) > 3 {
				if strings.Contains(line, ":") {
					soilMappingPos++
				} else {
					if soilMapping[soilMappingPos] == nil {
						soilMapping[soilMappingPos] = make([]SeedMap, 0)
					}
					nums := strings.Fields(line)
					s := SeedMap{}
					s.destination, _ = strconv.Atoi(nums[0])
					s.source, _ = strconv.Atoi(nums[1])
					s.length, _ = strconv.Atoi(nums[2])
					soilMapping[soilMappingPos] = append(soilMapping[soilMappingPos], s)
				}
			}
		}
	}
	for j := 0; j < len(seedRange); j++ {
		for i := seedRange[j].start; i < seedRange[j].start+seedRange[j].length; {
			l, minDist := getLocation(i, soilMapping)
			if minimalLocation > l {
				minimalLocation = l
			}
			if minDist < math.MaxInt {
				i = lib.Min(i+minDist, seedRange[j].start+seedRange[j].length)
			} else {
				i++
			}
		}
	}
	fmt.Printf("%d\n", minimalLocation)
}

// part1: 218 513 636

// part2: 81956384
