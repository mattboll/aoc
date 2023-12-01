package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"

	"github.com/mattboll/aoc/lib"
)

func main() {
	part2()
}

func numb(s string) int {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, word := range words {
		match, _ := regexp.MatchString(".*"+word, s)
		if match {
			return i + 1
		}
	}
	return 0
}

func part2() {
	strLines := lib.ReadFile()
	m := make([]int, len(strLines))
	for i, line := range strLines {
		runeLine := []rune(line)
		for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
			runeLine[i], runeLine[j] = runeLine[j], runeLine[i]
		}
		s1 := getChar(line) + getChar(string(runeLine))
		n, _ := strconv.Atoi(s1)
		m[i] = n
	}

	sum := 0
	for _, num := range m {
		sum += num
	}
	fmt.Printf("%v\n", sum)
	// res = 54431
}

func getChar(line string) string {
	s1 := ""
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			return string(ch)
		}
		s1 = s1 + string(ch)
		if r := numb(s1); r > 0 {
			return strconv.Itoa(r)
		}
	}
	return s1
}
