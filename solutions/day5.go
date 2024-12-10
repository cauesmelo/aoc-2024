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

func getUpdatesInRightOrder(rules []rule, updates []update) ([]update, []update) {
	correctUpdates := updates
	wrongUpdates := []update{}

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
					wrongUpdates = append(wrongUpdates, update)
				}
			}
		}
	}

	return correctUpdates, wrongUpdates
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
	correctUpdates, _ := getUpdatesInRightOrder(rules, updates)
	sumOfMiddlePageNumbers := sumMiddlePageNumberFromUpdates(correctUpdates)

	return sumOfMiddlePageNumbers
}

func fixWrongUpdate(updates []update, rules []rule) ([]update, int) {
	correctedUpdates := updates
	fixCount := 0

	for _, rule := range rules {
		for updateIdx := len(correctedUpdates) - 1; updateIdx >= 0; updateIdx-- {
			update := correctedUpdates[updateIdx]
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
					correctedUpdate := []int{}

					for _, change := range update {
						if change == rule.after {
							correctedUpdate = append(correctedUpdate, update[beforePos])
							correctedUpdate = append(correctedUpdate, rule.after)
						} else if change != rule.before {
							correctedUpdate = append(correctedUpdate, change)
						}
					}

					newCorrectedUpdates := [][]int{}

					newCorrectedUpdates = append(newCorrectedUpdates, correctedUpdates[:updateIdx]...)
					newCorrectedUpdates = append(newCorrectedUpdates, correctedUpdate)
					newCorrectedUpdates = append(newCorrectedUpdates, correctedUpdates[updateIdx+1:]...)

					correctedUpdates = newCorrectedUpdates
					fixCount++
				}
			}
		}
	}

	return correctedUpdates, fixCount
}

func (AOC) Day5_part2() int {
	lines := util.GetInput(5, false)

	rules, updates := readInput(lines)
	_, wrongUpdates := getUpdatesInRightOrder(rules, updates)
	fixedUpdates, fixCount := fixWrongUpdate(wrongUpdates, rules)

	c := 0
	for {
		if fixCount == 0 {
			break
		}

		if c > 3000 {
			break
		}

		fixedUpdates, fixCount = fixWrongUpdate(fixedUpdates, rules)
		c++
	}

	sum := sumMiddlePageNumberFromUpdates(fixedUpdates)

	return sum
}
