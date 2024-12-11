package solutions

import (
	"github.com/cauesmelo/aoc-2024/util"
)

func createAntinodeMap(size int) [][]rune {
	antinodeMap := make([][]rune, size)

	for i := range antinodeMap {
		for j := 0; j < size; j++ {
			antinodeMap[i] = append(antinodeMap[i], '.')
		}
	}

	return antinodeMap
}

func projectAntinodesForAntenna(lines []string, antinodeMap [][]rune, x int, y int) {
	antennaCh := rune(lines[y][x])
	for lineIdx, line := range lines {
		for chIdx, ch := range line {
			if ch != antennaCh {
				continue
			}

			xDiff := x - chIdx
			yDiff := y - lineIdx

			if xDiff == 0 && yDiff == 0 {
				continue
			}

			xDelta := x + xDiff
			yDelta := y + yDiff

			if xDelta < 0 || yDelta < 0 || xDelta > len(lines)-1 || yDelta > len(lines)-1 {
				continue
			}

			antinodeMap[yDelta][xDelta] = '#'
		}
	}
}

func projectAntinodes(lines []string, antinodeMap [][]rune, part2 bool) {
	for lineIdx, line := range lines {
		for chIdx, ch := range line {
			if ch != '.' {
				if part2 {
					projectAntinodesForAntenna2(lines, antinodeMap, chIdx, lineIdx)
				} else {
					projectAntinodesForAntenna(lines, antinodeMap, chIdx, lineIdx)
				}
			}
		}
	}
}

func countAntinodes(antinodeMap [][]rune) int {
	count := 0

	for i := range antinodeMap {
		for j := range antinodeMap[i] {
			if antinodeMap[i][j] == '#' {
				count++
			}
		}
	}

	return count
}

func (AOC) Day8_part1() int {
	lines := util.GetInput(8, false)

	antinodeMap := createAntinodeMap(len(lines[0]))

	projectAntinodes(lines, antinodeMap, false)

	return countAntinodes(antinodeMap)
}

func isAntennaTheOnly(lines []string, x int, y int) bool {
	antennaCh := rune(lines[y][x])
	for lineIdx, line := range lines {
		for chIdx, ch := range line {
			if ch != antennaCh {
				continue
			}

			if x == chIdx && y == lineIdx {
				continue
			}

			return false
		}
	}

	return true
}

func projectAntinodesForAntenna2(lines []string, antinodeMap [][]rune, x int, y int) {
	antennaCh := rune(lines[y][x])
	for lineIdx, line := range lines {
		for chIdx, ch := range line {
			if ch != antennaCh {
				continue
			}

			xDiff := x - chIdx
			yDiff := y - lineIdx

			if xDiff == 0 && yDiff == 0 {
				if !isAntennaTheOnly(lines, x, y) {
					antinodeMap[y][x] = '#'
				}

				continue
			}

			xDelta := x + xDiff
			yDelta := y + yDiff

			for {
				if xDelta < 0 || yDelta < 0 || xDelta > len(lines)-1 || yDelta > len(lines)-1 {
					break
				}

				antinodeMap[yDelta][xDelta] = '#'

				xDelta += xDiff
				yDelta += yDiff
			}
		}
	}
}

func (AOC) Day8_part2() int {
	lines := util.GetInput(8, false)

	antinodeMap := createAntinodeMap(len(lines[0]))

	projectAntinodes(lines, antinodeMap, true)

	return countAntinodes(antinodeMap)
}
