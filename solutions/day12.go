package solutions

import (
	"github.com/cauesmelo/aoc-2024/util"
)

type position struct {
	plant   rune
	visited bool
}

func mapPlants(lines []string) [][]position {
	plants := make([][]position, len(lines))

	for y, line := range lines {
		plants[y] = make([]position, len(line))
		for x, ch := range line {
			plants[y][x] = position{ch, false}
		}
	}

	return plants
}

func recPlant(plants [][]position, x, y int) (int, int) {
	if x < 0 || y < 0 || x >= len(plants[0]) || y >= len(plants) {
		return 0, 0
	}

	if plants[y][x].visited {
		return 0, 0
	}

	plants[y][x].visited = true

	area := 1
	per := 4
	r := plants[y][x].plant

	if y > 0 && plants[y-1][x].plant == r {
		per -= 1

		if !plants[y-1][x].visited {
			a, p := recPlant(plants, x, y-1)
			area += a
			per += p
		}
	}

	if y < len(plants)-1 && plants[y+1][x].plant == r {
		per -= 1

		if !plants[y+1][x].visited {
			a, p := recPlant(plants, x, y+1)
			area += a
			per += p
		}
	}

	if x < len(plants[0])-1 && plants[y][x+1].plant == r {
		per -= 1

		if !plants[y][x+1].visited {
			a, p := recPlant(plants, x+1, y)
			area += a
			per += p
		}
	}

	if x > 0 && plants[y][x-1].plant == r {
		per -= 1

		if !plants[y][x-1].visited {
			a, p := recPlant(plants, x-1, y)
			area += a
			per += p
		}
	}

	return area, per
}

func calculatePlants(plants [][]position) int {
	sum := 0

	for y, line := range plants {
		for x, pos := range line {
			if !pos.visited {
				area, per := recPlant(plants, x, y)
				sum += area * per
			}
		}
	}

	return sum
}

func (AOC) Day12_part1() int {
	lines := util.GetInput(12, false)

	plantsMatrix := mapPlants(lines)

	return calculatePlants(plantsMatrix)
}

func (AOC) Day12_part2() int {
	// lines := util.GetInput(1, true)

	return 0
}
