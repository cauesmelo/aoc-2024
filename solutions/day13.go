package solutions

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/cauesmelo/aoc-2024/util"
)

// btn A cost 3
// btn B cost 1

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

func parseTarget(str string) clawMachineTarget {
	re := regexp.MustCompile(`X\=(\d+),\s*Y\=(\d+)`)

	matches := re.FindStringSubmatch(str)

	xVal, _ := strconv.Atoi(matches[1])
	yVal, _ := strconv.Atoi(matches[2])

	return clawMachineTarget{
		x: xVal,
		y: yVal,
	}
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
			target: parseTarget(lines[i+2]),
		})
	}

	totalCost := 0

	for _, machine := range machines {
		totalCost += optimizeClaw(machine)
	}

	return totalCost
}

func (AOC) Day13_part2() int {
	// lines := util.GetInput(1, true)

	return 0
}
