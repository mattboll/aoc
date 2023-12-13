package main

import (
	"fmt"

	"github.com/mattboll/aoc/lib"
)

func main() {
	lib.RunAndMeasure(part1)
	lib.RunAndMeasure(part2)
}

type Galaxy struct {
	x        int
	y        int
	isGalaxy bool
	isEmpty  bool
	char     rune
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
func part1() {
	strLines := lib.ReadFile()

	s := 0
	linesToAdd := []int{}
	rowsToAdd := []int{}

	for key, line := range strLines {
		hasG := false
		for _, v := range line {
			if v == '#' {
				hasG = true
				break
			}
		}
		if !hasG {
			linesToAdd = append(linesToAdd, key)
		}
	}
	for i := len(linesToAdd) - 1; i >= 0; i-- {
		strLines = append(strLines[:linesToAdd[i]+1], strLines[linesToAdd[i]:]...)
	}

	for i := 0; i < len(strLines[0]); i++ {
		hasG := false
		for j := 0; j < len(strLines); j++ {
			if strLines[j][i] == '#' {
				hasG = true
				break
			}
		}
		if !hasG {
			rowsToAdd = append(rowsToAdd, i)
		}
	}

	for i := len(rowsToAdd) - 1; i >= 0; i-- {
		for j := 0; j < len(strLines); j++ {
			strLines[j] = strLines[j][:rowsToAdd[i]+1] + strLines[j][rowsToAdd[i]:]
		}
	}

	g := []Galaxy{}
	for i := 0; i < len(strLines); i++ {
		for j := 0; j < len(strLines[0]); j++ {
			if strLines[i][j] == '#' {
				g = append(g, Galaxy{x: j, y: i})
			}
		}
	}
	for i := 0; i < len(g); i++ {
		for j := i + 1; j < len(g); j++ {
			s += absDiffInt(g[i].x, g[j].x) + absDiffInt(g[i].y, g[j].y)
		}
	}

	fmt.Println(s)
}

func part2() {
	strLines := lib.ReadFile()

	galaxies := make([][]Galaxy, len(strLines))

	linesToAdd := []int{}
	rowsToAdd := []int{}

	for i, line := range strLines {
		galaxies[i] = make([]Galaxy, len(line))
		for j, r := range line {
			if r == '#' {
				galaxies[i][j] = Galaxy{x: j, y: i, isGalaxy: true, isEmpty: false}
			} else {
				galaxies[i][j] = Galaxy{x: j, y: i, isGalaxy: false, isEmpty: false}
			}
		}
	}

	for key, line := range galaxies {
		hasG := false
		for _, v := range line {
			if v.isGalaxy {
				hasG = true
				break
			}
		}
		if !hasG {
			linesToAdd = append(linesToAdd, key)
		}
	}
	for i := 0; i < len(galaxies[0]); i++ {
		hasG := false
		for _, line := range galaxies {
			if line[i].isGalaxy {
				hasG = true
				break
			}
		}
		if !hasG {
			rowsToAdd = append(rowsToAdd, i)
		}
	}

	for i := 0; i < len(linesToAdd); i++ {
		for _, line := range galaxies[linesToAdd[i]] {
			galaxies[line.y][line.x].isEmpty = true
		}
	}
	for i := 0; i < len(rowsToAdd); i++ {
		for _, line := range galaxies {
			e := line[rowsToAdd[i]]
			galaxies[e.y][e.x].isEmpty = true
		}
	}

	g := []Galaxy{}
	for i := 0; i < len(galaxies); i++ {
		for j := 0; j < len(galaxies[i]); j++ {
			if galaxies[i][j].isGalaxy {
				g = append(g, galaxies[i][j])
			}
		}
	}

	s := 0
	for i := 0; i < len(g); i++ {
		for j := i + 1; j < len(g); j++ {

			for y := g[i].y; y < g[j].y; y++ {
				if galaxies[y][g[i].x].isEmpty {
					s += 1000000
				} else {
					s += 1
				}
				if i == 1 {
					galaxies[y][g[i].x].char = '/'
				}
			}

			b := g[i].x
			e := g[j].x
			if b > e {
				b, e = e, b
			}

			for x := b; x < e; x++ {
				if galaxies[g[i].y][x].isEmpty {
					s += 1000000
				} else {
					s += 1
				}
				if i == 1 {
					galaxies[g[i].y][x].char = '/'
				}
			}

		}
	}

	for i := 0; i < len(galaxies); i++ {
		for j := 0; j < len(galaxies[i]); j++ {
			if galaxies[i][j].isGalaxy {
				fmt.Print("#")
			} else if galaxies[i][j].char == '/' {
				fmt.Print("/")
			} else if galaxies[i][j].isEmpty {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println(s)

}

// part 2 382286690609 too low
// part 2 382286640086 too low
// part 2 210360459686
// part 2 210360459686
