package solutions

import (
	"strconv"
	"strings"

	"github.com/cauesmelo/aoc-2024/util"
)

type equation struct {
	result  int
	factors []int
}

func parseEquations(lines []string) []equation {
	equations := make([]equation, 0)

	for _, line := range lines {
		splittedLine := strings.Split(line, ":")

		res, _ := strconv.Atoi(splittedLine[0])
		factors := util.GetNumbers(splittedLine[1])

		eq := equation{
			result:  res,
			factors: factors,
		}

		equations = append(equations, eq)
	}

	return equations
}

func isEquationPossible(eq *equation, idx int, acc int) bool {
	if idx == 0 {
		return isEquationPossible(eq, idx+1, eq.factors[idx])
	}

	if acc == eq.result && idx == len(eq.factors) {
		return true
	}

	if idx == len(eq.factors) {
		return false
	}

	if acc > eq.result {
		return false
	}

	return isEquationPossible(eq, idx+1, acc+eq.factors[idx]) || isEquationPossible(eq, idx+1, acc*eq.factors[idx])
}

func getSumOfPossibleEquations(equations []equation) int {
	sum := 0

	for _, eq := range equations {
		if isEquationPossible(&eq, 0, 0) {
			sum += eq.result
		}
	}

	return sum
}

func (AOC) Day7_part1() int {
	lines := util.GetInput(7, false)
	eqs := parseEquations(lines)

	return getSumOfPossibleEquations(eqs)
}

func isEquationPossible2(eq *equation, idx int, acc int) bool {
	if idx == 0 {
		return isEquationPossible2(eq, idx+1, eq.factors[idx])
	}

	if acc == eq.result && idx == len(eq.factors) {
		return true
	}

	if idx == len(eq.factors) {
		return false
	}

	if acc > eq.result {
		return false
	}

	if isEquationPossible2(eq, idx+1, acc+eq.factors[idx]) {
		return true
	}

	if isEquationPossible2(eq, idx+1, acc*eq.factors[idx]) {
		return true
	}

	v := strconv.Itoa(acc)
	v2 := strconv.Itoa(eq.factors[idx])
	vfs := v + v2
	vfi, _ := strconv.Atoi(vfs)

	return isEquationPossible2(eq, idx+1, vfi)
}

func getSumOfPossibleEquations2(equations []equation) int {
	sum := 0

	for _, eq := range equations {
		if isEquationPossible2(&eq, 0, 0) {
			sum += eq.result
		}
	}

	return sum
}

func (AOC) Day7_part2() int {
	lines := util.GetInput(7, false)
	eqs := parseEquations(lines)

	return getSumOfPossibleEquations2(eqs)
}
