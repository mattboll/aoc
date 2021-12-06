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

type Num struct {
	b bool
	i int
}

func main() {
	strLines := ReadFile()
	strList := strings.SplitN(strLines[0], ",", -1)
	picked := []int{}
	for _, s := range strList {
		if n, err := strconv.Atoi(s); err == nil {
			picked = append(picked, n)
		}
	}

	nums := []Num{}

	for _, line := range strLines[1:] {
		row := strings.SplitN(line, " ", -1)
		for _, s := range row {
			if n, err := strconv.Atoi(s); err == nil {
				nums = append(nums, Num{i: n})
			}
		}
	}

	for _, n := range picked {
		for k, _ := range nums {
			if nums[k].i == n {
				nums[k].b = true
			}
		}

		nbCards := len(nums) / 25
		for k := nbCards - 1; k >= 0; k-- {
			if cardOk(nums, k) {
				if len(nums) <= 25 {
					show(nums, 0, n)
					return
				}
				nums = append(nums[:(k)*25], nums[(k+1)*25:]...)
			}
		}
	}
}

func cardOk(nums []Num, k int) bool {
	p := k * 25
	for i := 0; i < 5; i++ {
		found := true
		for j := 0; j < 5; j++ {
			if !nums[p+i+j*5].b {
				found = false
			}
		}
		if found {
			return true
		}
	}
	for i := 0; i < 5; i++ {
		found := true
		for j := 0; j < 5; j++ {
			if !nums[p+i*5+j].b {
				found = false
			}
		}
		if found {
			return true
		}
	}
	return false
}

func show(nums []Num, k int, n int) {
	fmt.Printf("%v \n", nums)
	sum := 0
	p := k * 25
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !nums[p+i+j*5].b {
				sum += nums[p+i+j*5].i
			}
		}
	}
	fmt.Printf("n : %d \n", n)
	fmt.Printf("sum : %d \n", sum)
	fmt.Printf("res : %d \n", n*sum)
}
