package lib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadFile() []string {
	file, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	fileContent := string(file)
	slicedContent := strings.Split(fileContent, "\n")
	if slicedContent[len(slicedContent)-1] == "" {
		slicedContent = slicedContent[:len(slicedContent)-1]
	}
	return slicedContent
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RunAndMeasure(fn func()) {
	start := time.Now()
	fn()
	fmt.Println(time.Since(start))
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
