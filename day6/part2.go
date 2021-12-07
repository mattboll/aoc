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
	fish := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, line := range strLines {
		if line == "" {
			break
		}
		r1 := strings.SplitN(line, ",", -1)
		for _, v := range r1 {
			n, _ := strconv.Atoi(v)
			fish[n] = fish[n] + 1
		}
	}

	for i := 0; i < 256; i++ {
		zeroFish := fish[0]
		fish[0] = fish[1]
		fish[1] = fish[2]
		fish[2] = fish[3]
		fish[3] = fish[4]
		fish[4] = fish[5]
		fish[5] = fish[6]
		fish[6] = fish[7] + zeroFish
		fish[7] = fish[8]
		fish[8] = zeroFish
	}
	sum := fish[0] + fish[1] + fish[2] + fish[3] + fish[4] + fish[5] + fish[6] + fish[7] + fish[8]

	fmt.Printf("fish length : %d \n", sum)
}
