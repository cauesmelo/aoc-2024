package solutions

import (
	"strconv"
	"strings"

	"github.com/cauesmelo/aoc-2024/util"
)

type coord struct {
	x int
	y int
}

type robot struct {
	pos coord
	vel coord
}

func parseRobots(lines []string) []robot {
	robots := make([]robot, 0)

	for _, line := range lines {
		splits := strings.Split(line, " ")

		vals := make([][]int, 0)

		for _, split := range splits {
			s := strings.Split(split, "=")

			strVals := strings.Split(s[1], ",")

			v1, _ := strconv.Atoi(strVals[0])
			v2, _ := strconv.Atoi(strVals[1])

			vals = append(vals, []int{v1, v2})
		}

		r := robot{
			pos: coord{
				x: vals[0][0],
				y: vals[0][1],
			},
			vel: coord{
				x: vals[1][0],
				y: vals[1][1],
			},
		}

		robots = append(robots, r)
	}

	return robots
}

func walkRobot(size coord, r robot, seconds int) coord {
	x := (r.pos.x + r.vel.x*seconds) % size.x
	y := (r.pos.y + r.vel.y*seconds) % size.y

	if x < 0 {
		x += size.x
	}

	if y < 0 {
		y += size.y
	}

	return coord{
		x, y,
	}
}

func getQuadrant(c coord, size coord) int {
	halfX := size.x / 2
	halfY := size.y / 2

	if c.x < halfX && c.y < halfY {
		return 1
	}

	if c.x > halfX && c.y < halfY {
		return 2
	}

	if c.x > halfX && c.y > halfY {
		return 3
	}

	if c.x < halfX && c.y > halfY {
		return 4
	}

	return 0
}

func (AOC) Day14_part1() int {
	lines := util.GetInput(14, false)

	robots := parseRobots(lines)

	grid := coord{
		x: 101,
		y: 103,
	}

	finalCoords := make([]coord, 0)

	for _, robot := range robots {
		finalCoords = append(finalCoords, walkRobot(grid, robot, 100))
	}

	quadrants := []int{0, 0, 0, 0}

	for _, c := range finalCoords {
		q := getQuadrant(c, grid)

		if q == 0 {
			continue
		}

		quadrants[q-1]++
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func (AOC) Day14_part2() int {
	// lines := util.GetInput(13, false)

	return 0
}
