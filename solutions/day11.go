package solutions

import (
	"strconv"

	"github.com/cauesmelo/aoc-2024/util"
)

func splitIntInHalf(n int) []int {
	s := strconv.Itoa(n)

	half := len(s) / 2

	p1 := s[:half]
	p2 := s[half:]

	np1, _ := strconv.Atoi(p1)
	np2, _ := strconv.Atoi(p2)

	return []int{np1, np2}
}

func isIntEvenDigits(n int) bool {
	return len(strconv.Itoa(n))%2 == 0
}

func blink(stones []int) []int {
	newStones := make([]int, 0)

	for i := 0; i < len(stones); i++ {
		s := stones[i]

		if s == 0 {
			newStones = append(newStones, 1)
			continue
		}

		if isIntEvenDigits(s) {
			splitted := splitIntInHalf(s)
			newStones = append(newStones, splitted...)
			continue
		}

		newStones = append(newStones, s*2024)
	}

	return newStones
}

func (AOC) Day11_part1() int {
	lines := util.GetInput(11, false)

	stones := util.GetNumbers(lines[0])

	blinks := 25

	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}

	return len(stones)
}

func (AOC) Day11_part2() int {
	// lines := util.GetInput(1, true)

	return 0
}
