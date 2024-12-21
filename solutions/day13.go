package solutions

import (
	"math"
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

func optimizeClawRec(machine clawMachine, x, y, cost int) int {
	if x > machine.target.x || y > machine.target.y {
		return int(math.Inf(1))
	}

	costA := optimizeClawRec(machine, x+machine.a.xDelta, y+machine.a.yDelta, cost+machine.a.cost)
	costB := optimizeClawRec(machine, x+machine.b.xDelta, y+machine.b.yDelta, cost+machine.b.cost)

	if costA < costB {
		return costA
	}

	return costB
}

func optimizeClaw(machine clawMachine) int {
	return optimizeClawRec(machine, 0, 0, 0)
}

func (AOC) Day13_part1() int {
	lines := util.GetInput(13, true)

	machines := make([]clawMachine, 0)

	for i := 0; i < len(lines); i += 4 {
		machines = append(machines, clawMachine{
			a:      parseButton(lines[i]),
			b:      parseButton(lines[i+1]),
			target: parseTarget(lines[i+2]),
		})
	}

	return optimizeClaw(machines[0])
}

func (AOC) Day13_part2() int {
	// lines := util.GetInput(1, true)

	return 0
}
