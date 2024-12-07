package solutions

import (
	"regexp"
	"strconv"
	"strings"

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

func parseCommands(line string, status bool) (string, bool) {
	active := status

	parsedLine := ""
	currInstruction := ""

	i := -1
	for {
		i++

		if len(line) <= i {
			break
		}

		c := line[i]

		if currInstruction == "" && c != 'd' {
			if active {
				parsedLine += string(c)
			}
			continue
		}

		currInstruction += string(c)

		insDo := strings.Contains("do()", currInstruction)
		insDont := strings.Contains("don't()", currInstruction)

		if insDo && insDont {
			continue
		}

		if insDo {
			active = true
			i += 1
		}

		if insDont {
			active = false
			i += 1
		}

		currInstruction = ""
	}

	return parsedLine, active
}

func (AOC) Day3_part2() int {
	lines := util.GetInput(3, false)

	cmds := make([]string, 0)
	active := true

	for _, line := range lines {
		newCmds, newStatus := parseCommands(line, active)

		cmds = append(cmds, newCmds)
		active = newStatus
	}

	sum := 0
	for _, cmd := range cmds {
		sum += multLine(cmd)
	}

	return sum
}
