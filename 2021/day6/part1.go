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
	fish := []int{}
	for _, line := range strLines {
		if line == "" {
			break
		}
		r1 := strings.SplitN(line, ",", -1)
		for _, v := range r1 {
			n, _ := strconv.Atoi(v)
			fish = append(fish, n)
		}
	}

	for i := 0; i < 80; i++ {
		for k, v := range fish {
			if v == 0 {
				fish[k] = 6
				fish = append(fish, 8)
			} else {
				fish[k] = v - 1
			}
		}
	}
	fmt.Printf("fish length : %d \n", len(fish))
}
