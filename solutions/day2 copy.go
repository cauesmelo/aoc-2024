package solutions

import (
	"regexp"
	"strconv"

	"github.com/cauesmelo/aoc-2024/util"
)

func multLine(line string) int {
	pattern := `mul\(([0-9]*),([0-9]*)\)`
	regex := regexp.MustCompile(pattern)

	matches := regex.FindAllStringSubmatch(line, -1)

	sum := 0
	for _, match := range matches {
		a, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		sum += a * b
	}

	return sum
}

func (AOC) Day3_part1() int {
	lines := util.GetInput(3, false)

	sum := 0
	for _, line := range lines {
		sum += multLine(line)
	}

	return sum
}

func (AOC) Day3_part2() int {
	// lines := util.GetInput(3, true)

	return 0
}
