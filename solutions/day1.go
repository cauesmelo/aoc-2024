package solutions

import (
	"math"
	"sort"

	"github.com/cauesmelo/aoc-2024/util"
)

func (AOC) Day1_part1() int {
	lines := util.GetInput(1, false)

	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for idx, line := range lines {
		numbers := util.GetNumbers(line)

		left[idx] = numbers[0]
		right[idx] = numbers[1]
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for idx := range lines {
		dist := int(math.Abs(float64(left[idx] - right[idx])))
		sum += dist
	}

	return sum
}

func (AOC) Day1_part2() int {
	lines := util.GetInput(1, false)

	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for idx, line := range lines {
		numbers := util.GetNumbers(line)

		left[idx] = numbers[0]
		right[idx] = numbers[1]
	}

	sum := 0
	for _, l := range left {
		score := 0

		for _, r := range right {
			if l == r {
				score += l
			}
		}

		sum += score
	}

	return sum
}
