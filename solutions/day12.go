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

type coordinate struct {
	x, y int
}

func recPlantv2(plants [][]position, x, y int) (int, int, []coordinate) {
	if x < 0 || y < 0 || x >= len(plants[0]) || y >= len(plants) {
		return 0, 0, nil
	}

	if plants[y][x].visited {
		return 0, 0, nil
	}

	plants[y][x].visited = true

	area := 1
	per := 4
	coords := []coordinate{{x, y}}
	r := plants[y][x].plant

	if y > 0 && plants[y-1][x].plant == r {
		per -= 1

		if !plants[y-1][x].visited {
			a, p, c := recPlantv2(plants, x, y-1)
			area += a
			per += p
			coords = append(coords, c...)
		}
	}

	if y < len(plants)-1 && plants[y+1][x].plant == r {
		per -= 1

		if !plants[y+1][x].visited {
			a, p, c := recPlantv2(plants, x, y+1)
			area += a
			per += p
			coords = append(coords, c...)
		}
	}

	if x < len(plants[0])-1 && plants[y][x+1].plant == r {
		per -= 1

		if !plants[y][x+1].visited {
			a, p, c := recPlantv2(plants, x+1, y)
			area += a
			per += p
			coords = append(coords, c...)
		}
	}

	if x > 0 && plants[y][x-1].plant == r {
		per -= 1

		if !plants[y][x-1].visited {
			a, p, c := recPlantv2(plants, x-1, y)
			area += a
			per += p
			coords = append(coords, c...)
		}
	}

	return area, per, coords
}

func calculateSides(coords []coordinate, length int) int {

	leftImpact := make([]coordinate, 0)
	for line := 0; line < length; line++ {
		lastImpactCol := -1

		for col := 0; col < length; col++ {
			for _, coord := range coords {
				if coord.x == col && coord.y == line {
					if lastImpactCol == -1 {
						leftImpact = append(leftImpact, coord)
						lastImpactCol = col
						break
					} else if col-lastImpactCol > 1 {
						leftImpact = append(leftImpact, coord)
						lastImpactCol = col
						break
					} else {
						lastImpactCol = col
					}
				}
			}
		}
	}

	rightImpact := make([]coordinate, 0)
	for line := 0; line < length; line++ {
		lastImpactCol := -1

		for col := length - 1; col >= 0; col-- {
			for _, coord := range coords {
				if coord.x == col && coord.y == line {
					if lastImpactCol == -1 {
						rightImpact = append(rightImpact, coord)
						lastImpactCol = col
						break
					} else if lastImpactCol-col > 1 {
						rightImpact = append(rightImpact, coord)
						lastImpactCol = col
						break
					} else {
						lastImpactCol = col
					}
				}
			}
		}
	}

	topImpact := make([]coordinate, 0)
	for col := 0; col < length; col++ {
		lastImpactLine := -1

		for line := 0; line < length; line++ {
			for _, coord := range coords {
				if coord.x == col && coord.y == line {
					if lastImpactLine == -1 {
						topImpact = append(topImpact, coord)
						lastImpactLine = line
						break
					} else if line-lastImpactLine > 1 {
						topImpact = append(topImpact, coord)
						lastImpactLine = line
						break
					} else {
						lastImpactLine = line
					}
				}
			}
		}
	}

	downImpact := make([]coordinate, 0)
	for col := 0; col < length; col++ {
		lastImpactLine := -1

		for line := length - 1; line >= 0; line-- {
			for _, coord := range coords {
				if coord.x == col && coord.y == line {
					if lastImpactLine == -1 {
						downImpact = append(downImpact, coord)
						lastImpactLine = line
						break
					} else if lastImpactLine-line > 1 {
						downImpact = append(downImpact, coord)
						lastImpactLine = line
						break
					} else {
						lastImpactLine = line
					}
				}
			}
		}
	}

	leftCount := 0
	for col := 0; col < length; col++ {
		impactsOnCol := make([]coordinate, 0)

		for _, impact := range leftImpact {
			if impact.x == col {
				impactsOnCol = append(impactsOnCol, impact)
			}
		}

		lastImpact := -1

		for _, impact := range impactsOnCol {
			if lastImpact == -1 {
				lastImpact = impact.y
				leftCount++
				continue
			}

			if impact.y-lastImpact > 1 {
				leftCount++
			}

			lastImpact = impact.y
		}
	}

	rightCount := 0
	for col := length - 1; col >= 0; col-- {
		impactsOnCol := make([]coordinate, 0)

		for _, impact := range rightImpact {
			if impact.x == col {
				impactsOnCol = append(impactsOnCol, impact)
			}
		}

		lastImpact := -1

		for _, impact := range impactsOnCol {
			if lastImpact == -1 {
				lastImpact = impact.y
				rightCount++
				continue
			}

			if impact.y-lastImpact > 1 {
				rightCount++
			}

			lastImpact = impact.y
		}
	}

	topCount := 0
	for line := 0; line < length; line++ {
		impactsOnLine := make([]coordinate, 0)

		for _, impact := range topImpact {
			if impact.y == line {
				impactsOnLine = append(impactsOnLine, impact)
			}
		}

		lastImpact := -1

		for _, impact := range impactsOnLine {
			if lastImpact == -1 {
				lastImpact = impact.x
				topCount++
				continue
			}

			if impact.x-lastImpact > 1 {
				topCount++
			}

			lastImpact = impact.x
		}
	}

	downCount := 0
	for line := length - 1; line >= 0; line-- {
		impactsOnLine := make([]coordinate, 0)

		for _, impact := range downImpact {
			if impact.y == line {
				impactsOnLine = append(impactsOnLine, impact)
			}
		}

		lastImpact := -1

		for _, impact := range impactsOnLine {
			if lastImpact == -1 {
				lastImpact = impact.x
				downCount++
				continue
			}

			if impact.x-lastImpact > 1 {
				downCount++
			}

			lastImpact = impact.x
		}
	}

	return leftCount + rightCount + topCount + downCount
}

func calculatePlantsv2(plants [][]position) int {
	sum := 0

	for y, line := range plants {
		for x, pos := range line {
			if !pos.visited {
				area, _, coords := recPlantv2(plants, x, y)
				sides := calculateSides(coords, len(plants[0]))
				sum += area * sides

			}
		}
	}

	return sum
}

func (AOC) Day12_part2() int {
	lines := util.GetInput(12, false)

	plantsMatrix := mapPlants(lines)

	return calculatePlantsv2(plantsMatrix)
}
