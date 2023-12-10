package main

import (
	"fmt"

	"github.com/mattboll/aoc/lib"
)

func main() {
	lib.RunAndMeasure(part1)
	lib.RunAndMeasure(part2)
}

type Pipe struct {
	x int
	y int
	r rune
	d int
}

func part1() {
	strLines := lib.ReadFile()
	var matrix [][]Pipe
	matrix = make([][]Pipe, len(strLines))
	for i := 0; i < len(strLines); i++ {
		matrix[i] = make([]Pipe, len(strLines[0]))
	}

	for y, line := range strLines {
		for x := 0; x < len(line); x++ {
			matrix[x][y] = Pipe{x: x, y: y, r: rune(line[x]), d: 0}
		}
	}

	var spot Pipe
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[x][y].r == 'S' {
				spot = matrix[x][y]
				break
			}
		}
		if matrix[spot.x][spot.y].r == 'S' {
			break
		}
	}
	var n1, n2 Pipe
	// home start is not first/last line/row
	if matrix[spot.x][spot.y-1].r == 'F' || matrix[spot.x][spot.y-1].r == '|' || matrix[spot.x][spot.y-1].r == '7' {
		n1 = matrix[spot.x][spot.y-1]
	}
	if matrix[spot.x-1][spot.y].r == '-' || matrix[spot.x-1][spot.y].r == 'L' || matrix[spot.x-1][spot.y].r == 'F' {
		if n1.r == 0 {
			n1 = matrix[spot.x-1][spot.y]
		} else {
			n2 = matrix[spot.x-1][spot.y]
		}
	}
	if matrix[spot.x+1][spot.y].r == '-' || matrix[spot.x+1][spot.y].r == '7' || matrix[spot.x+1][spot.y].r == 'J' {
		if n1.r == 0 {
			n1 = matrix[spot.x+1][spot.y]
		} else {
			n2 = matrix[spot.x+1][spot.y]
		}
	}
	if n2.r == 0 {
		n2 = matrix[spot.x][spot.y+1]
	}

	s := getValue(&matrix, &spot, &spot, &n1, &n2)

	fmt.Println(s)
}

func nextValue(matrix *[][]Pipe, pn *Pipe, n *Pipe) *Pipe {
	if n.r == '-' {
		if pn.x < n.x {
			return &(*matrix)[n.x+1][n.y]
		}
		return &(*matrix)[n.x-1][n.y]
	}
	if n.r == '|' {
		if pn.y < n.y {
			return &(*matrix)[n.x][n.y+1]
		}
		return &(*matrix)[n.x][n.y-1]
	}
	if n.r == 'L' {
		if pn.y < n.y {
			return &(*matrix)[n.x+1][n.y]
		}
		return &(*matrix)[n.x][n.y-1]
	}
	if n.r == '7' {
		if pn.x < n.x {
			return &(*matrix)[n.x][n.y+1]
		}
		return &(*matrix)[n.x-1][n.y]
	}
	if n.r == 'F' {
		if n.x < pn.x {
			return &(*matrix)[n.x][n.y+1]
		}
		return &(*matrix)[n.x+1][n.y]
	}
	if n.r == 'J' {
		if pn.x < n.x {
			return &(*matrix)[n.x][n.y-1]
		}
		return &(*matrix)[n.x-1][n.y]
	}

	return nil
}

func getValue(matrix *[][]Pipe, pn1 *Pipe, pn2 *Pipe, n1 *Pipe, n2 *Pipe) int {
	n1.d = pn1.d + 1
	n2.d = pn2.d + 1
	if n1.x == n2.x && n1.y == n2.y {
		return n1.d
	}

	nn1 := nextValue(matrix, pn1, n1)
	nn2 := nextValue(matrix, pn2, n2)
	if nn1 == nil || nn2 == nil {
		return n1.d
	}

	return getValue(matrix, n1, n2, nn1, nn2)

}

func part2() {
}
