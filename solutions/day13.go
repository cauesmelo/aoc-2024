package solutions

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/cauesmelo/aoc-2024/util"
)

type clawMachineButton struct {
	xDelta int
	yDelta int
	cost   int
}

type clawMachineTarget struct {
	x int
	y int
}

type clawMachine struct {
	a      clawMachineButton
	b      clawMachineButton
	target clawMachineTarget
}

func parseButton(str string) clawMachineButton {
	re := regexp.MustCompile(`X\+(\d+),\s*Y\+(\d+)`)

	matches := re.FindStringSubmatch(str)

	xVal, _ := strconv.Atoi(matches[1])
	yVal, _ := strconv.Atoi(matches[2])

	cost := 3

	if strings.Contains(str, "Button B") {
		cost = 1
	}

	return clawMachineButton{
		xDelta: xVal,
		yDelta: yVal,
		cost:   cost,
	}
}

func parseTarget(str string, p2 bool) clawMachineTarget {
	re := regexp.MustCompile(`X\=(\d+),\s*Y\=(\d+)`)

	matches := re.FindStringSubmatch(str)

	m1 := matches[1]
	m2 := matches[2]

	xVal, _ := strconv.Atoi(m1)
	yVal, _ := strconv.Atoi(m2)

	c := clawMachineTarget{
		x: xVal,
		y: yVal,
	}

	if p2 {
		c.x = c.x + 10000000000000
		c.y = c.y + 10000000000000
	}

	return c
}

func optimizeClaw(machine clawMachine) int {
	minCost := 0

	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			xPos := machine.a.xDelta*a + machine.b.xDelta*b
			yPos := machine.a.yDelta*a + machine.b.yDelta*b

			if xPos == machine.target.x && yPos == machine.target.y {
				c := a*machine.a.cost + b*machine.b.cost

				if c < minCost || minCost == 0 {
					minCost = c
				}
			}
		}
	}

	return minCost
}

func (AOC) Day13_part1() int {
	lines := util.GetInput(13, false)

	machines := make([]clawMachine, 0)

	for i := 0; i < len(lines); i += 4 {
		machines = append(machines, clawMachine{
			a:      parseButton(lines[i]),
			b:      parseButton(lines[i+1]),
			target: parseTarget(lines[i+2], false),
		})
	}

	totalCost := 0

	for _, machine := range machines {
		totalCost += optimizeClaw(machine)
	}

	return totalCost
}

func optimizeClawv2(machine clawMachine) int {
	ca1 := float64(machine.target.x)*float64(machine.b.yDelta) - float64(machine.target.y)*float64(machine.b.xDelta)
	ca2 := float64(machine.a.xDelta)*float64(machine.b.yDelta) - float64(machine.a.yDelta)*float64(machine.b.xDelta)
	ca := ca1 / ca2
	cb := (float64(machine.target.x) - float64(machine.a.xDelta)*ca) / float64(machine.b.xDelta)

	isSolution := ca == float64(int(ca)) && cb == float64(int(cb))

	if isSolution {
		return int(ca)*machine.a.cost + int(cb)*machine.b.cost
	}

	return 0
}

func (AOC) Day13_part2() int {
	lines := util.GetInput(13, false)

	machines := make([]clawMachine, 0)

	for i := 0; i < len(lines); i += 4 {
		machines = append(machines, clawMachine{
			a:      parseButton(lines[i]),
			b:      parseButton(lines[i+1]),
			target: parseTarget(lines[i+2], true),
		})
	}

	totalCost := 0

	for _, machine := range machines {
		totalCost += optimizeClawv2(machine)
	}

	return totalCost
}
