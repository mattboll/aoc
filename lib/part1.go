package lib

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func ReadFile() []string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fileContent := string(file)
	slicedContent := strings.Split(fileContent, "\n")
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
