package util

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetNumbersSep(line string, sep string) []int {
	parts := strings.Split(line, sep)

	numbers := make([]int, 0)

	for _, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, val)
	}

	return numbers
}

func GetNumbers(line string) []int {
	re := regexp.MustCompile(`-?\d+`)

	f := re.FindAllStringIndex(line, -1)

	numbers := make([]int, 0)

	for _, match := range f {
		valueStr := line[match[0]:match[1]]
		val, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, val)
	}

	return numbers
}

func GetInput(day int, test bool) []string {
	partN := 2
	if test {
		partN = 1
	}

	fileName := fmt.Sprintf("d%d_%d.txt", day, partN)

	dat, err := os.ReadFile("./input/" + fileName)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	return lines
}
