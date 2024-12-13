package solutions

import (
	"strconv"

	"github.com/cauesmelo/aoc-2024/util"
)

func calcPathRec(lines [][]int, curr int, x int, y int, p2 bool) int {
	max := len(lines[0]) - 1

	if curr == 9 {
		if !p2 {
			lines[y][x] = -1
		}
		return 1
	}

	sum := 0

	if x > 0 && lines[y][x-1] == curr+1 {
		sum += calcPathRec(lines, curr+1, x-1, y, p2)
	}

	if x < max && lines[y][x+1] == curr+1 {
		sum += calcPathRec(lines, curr+1, x+1, y, p2)
	}

	if y < max && lines[y+1][x] == curr+1 {
		sum += calcPathRec(lines, curr+1, x, y+1, p2)
	}

	if y > 0 && lines[y-1][x] == curr+1 {
		sum += calcPathRec(lines, curr+1, x, y-1, p2)
	}

	return sum
}

func calcPath(lines []string, p2 bool) int {

	arr := make([][]int, 0)

	for _, line := range lines {
		arrLine := make([]int, 0)
		for _, c := range line {
			if string(c) == "." {
				arrLine = append(arrLine, -1)
				continue
			}

			n, _ := strconv.Atoi(string(c))
			arrLine = append(arrLine, n)
		}
		arr = append(arr, arrLine)
	}

	sum := 0

	for y, line := range arr {
		for x, n := range line {
			if n == 0 {
				if p2 {
					sum += calcPathRec(arr, 0, x, y, p2)
				} else {
					arrCp := make([][]int, len(arr))
					for i, v := range arr {
						arrCp[i] = make([]int, len(v))
						copy(arrCp[i], v)
					}

					sum += calcPathRec(arrCp, 0, x, y, p2)
				}
			}
		}
	}

	return sum
}

func (AOC) Day10_part1() int {
	lines := util.GetInput(10, false)

	return calcPath(lines, false)
}

func (AOC) Day10_part2() int {
	lines := util.GetInput(10, false)

	return calcPath(lines, true)
}
