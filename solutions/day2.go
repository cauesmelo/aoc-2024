package solutions

import (
	"math"

	"github.com/cauesmelo/aoc-2024/util"
)

func checkSafety(numbers []int) bool {
	var direction = ""

	for idx, n := range numbers {
		if idx == 0 {
			continue
		}

		prevN := numbers[idx-1]

		if prevN == n {
			return false
		}

		if direction == "" {
			if prevN < n {
				direction = "up"
			} else {
				direction = "down"
			}
		} else if (direction == "up") && prevN > n {
			return false
		} else if (direction == "down") && prevN < n {
			return false
		}

		diff := int(math.Abs(float64(prevN) - float64(n)))

		if diff > 3 {
			return false
		}
	}

	return true
}

func (AOC) Day2_part1() int {
	lines := util.GetInput(2, false)

	count := 0
	for _, line := range lines {
		numbers := util.GetNumbers(line)

		if checkSafety(numbers) {
			count++
		}
	}

	return count
}

func getSliceExcludingIdx(numbers []int, idx int) []int {
	slice := make([]int, 0)
	for i, n := range numbers {
		if i == idx {
			continue
		}

		slice = append(slice, n)
	}

	return slice
}

func (AOC) Day2_part2() int {
	lines := util.GetInput(2, false)

	count := 0
	for _, line := range lines {
		numbers := util.GetNumbers(line)

		combination := make([][]int, 0)

		for idx := range numbers {
			combination = append(combination, getSliceExcludingIdx(numbers, idx))
		}

		for _, c := range combination {
			if checkSafety(c) {
				count++
				break
			}
		}
	}

	return count
}
