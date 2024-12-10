package solutions

import (
	"github.com/cauesmelo/aoc-2024/util"
)

type rule struct {
	before int
	after  int
}

type update = []int

func readInput(lines []string) ([]rule, []update) {
	rules := []rule{}
	updates := []update{}

	readingRules := true
	for _, line := range lines {
		if line == "" {
			readingRules = false
			continue
		}

		if readingRules {
			numbers := util.GetNumbersSep(line, "|")

			rules = append(rules, rule{
				before: numbers[0],
				after:  numbers[1],
			})
		} else {
			numbers := util.GetNumbersSep(line, ",")
			updates = append(updates, numbers)
		}
	}

	return rules, updates
}

func getUpdatesInRightOrder(rules []rule, updates []update) []update {
	correctUpdates := updates

	for _, rule := range rules {
		for updateIdx := len(correctUpdates) - 1; updateIdx >= 0; updateIdx-- {
			update := correctUpdates[updateIdx]
			beforePos := -1
			afterPos := -1

			for i, change := range update {
				if change == rule.before {
					beforePos = i
				}

				if change == rule.after {
					afterPos = i
				}
			}

			if beforePos == -1 && afterPos == -1 {
				continue
			}

			if beforePos != -1 && afterPos != -1 {
				if beforePos > afterPos {
					correctUpdates = append(correctUpdates[:updateIdx], correctUpdates[updateIdx+1:]...)
				}
			}
		}
	}

	return correctUpdates
}

func sumMiddlePageNumberFromUpdates(updates []update) int {
	sum := 0

	for _, update := range updates {
		middlePos := len(update) / 2
		sum += update[middlePos]
	}

	return sum
}

func (AOC) Day5_part1() int {
	lines := util.GetInput(5, false)

	rules, updates := readInput(lines)
	correctUpdates := getUpdatesInRightOrder(rules, updates)
	sumOfMiddlePageNumbers := sumMiddlePageNumberFromUpdates(correctUpdates)

	return sumOfMiddlePageNumbers
}

func (AOC) Day5_part2() int {
	// lines := util.GetInput(1, true)

	return 0
}
