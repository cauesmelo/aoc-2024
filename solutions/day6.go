package solutions

import (
	"github.com/cauesmelo/aoc-2024/util"
)

type direction struct {
	x int
	y int
}

func getStartPos(lab []string) (int, int) {
	for i, line := range lab {
		for j, c := range line {
			if c == '^' {
				return j, i
			}
		}
	}

	panic("No start position found")
}

func walkTheLab(lab []string) [][]rune {
	dir := direction{
		x: 0,
		y: -1, // UP
	}

	x, y := getStartPos(lab)
	labPath := [][]rune{}
	stepCount := 0

	for _, line := range lab {
		labPath = append(labPath, []rune(line))
	}

	for {
		stepCount++

		destX := x + dir.x
		destY := y + dir.y

		if destX == len(labPath) || destY == len(labPath[0]) || destX < 0 || destY < 0 {
			break
		}

		destination := labPath[destY][destX]

		if destination == '#' {
			if dir.x == -1 {
				dir.x = 0
				dir.y = -1
				continue
			}

			if dir.x == 1 {
				dir.x = 0
				dir.y = 1
				continue
			}

			if dir.y == 1 {
				dir.y = 0
				dir.x = -1
				continue
			}

			if dir.y == -1 {
				dir.y = 0
				dir.x = 1
				continue
			}
		} else {
			x += dir.x
			y += dir.y
			labPath[y][x] = 'X'
		}

		if x == len(labPath) || y == len(labPath[0]) || x < 0 || y < 0 {
			break
		}
	}

	return labPath
}

func countPath(labPath [][]rune) int {
	count := 0

	for _, line := range labPath {
		for _, c := range line {
			if c == 'X' {
				count++
			}
		}
	}

	return count
}

func (AOC) Day6_part1() int {
	lines := util.GetInput(6, false)

	labPath := walkTheLab(lines)
	return countPath(labPath)
}

func (AOC) Day6_part2() int {
	// lines := util.GetInput(1, true)

	return 0
}
