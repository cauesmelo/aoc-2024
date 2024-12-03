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

func (AOC) Day2_part2() int {
	// lines := util.GetInput(1, true)

	return 0
}
