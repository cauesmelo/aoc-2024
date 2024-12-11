package solutions

import (
	"strconv"
	"strings"

	"github.com/cauesmelo/aoc-2024/util"
)

func mountDisk(digits string) string {
	disk := ""
	currN := 0
	isFile := true

	for i := 0; i < len(digits); i++ {
		nTimes, _ := strconv.Atoi(string(digits[i]))

		if isFile {
			n := strconv.Itoa(currN)
			for j := 0; j < nTimes; j++ {
				disk += n + ";"
			}
			currN++
		} else {
			for j := 0; j < nTimes; j++ {
				disk += "." + ";"
			}
		}

		isFile = !isFile
	}

	return disk
}

func compactDisk(disk string) string {
	compactDisk := ""

	parsedDisk := strings.Split(disk, ";")
	parsedDisk = parsedDisk[:len(parsedDisk)-1]
	endPointer := len(parsedDisk) - 1

	for i := 0; i < len(parsedDisk); i++ {
		if i > endPointer {
			break
		}

		if parsedDisk[i] == "." {
			for {
				ch := parsedDisk[endPointer]

				if ch != "." {
					compactDisk += string(parsedDisk[endPointer]) + ";"
					endPointer--
					break
				}

				endPointer--

				if endPointer == i-1 {
					return compactDisk
				}
			}

			if endPointer <= i {
				break
			}
		} else {
			compactDisk += string(parsedDisk[i]) + ";"
		}
	}

	return compactDisk
}

func getChecksum(compactDisk string) int {
	parsedDisk := strings.Split(compactDisk, ";")
	parsedDisk = parsedDisk[:len(parsedDisk)-1]

	sum := 0

	for i := 0; i < len(parsedDisk); i++ {
		digit, _ := strconv.Atoi(string(parsedDisk[i]))
		sum += i * digit
	}

	return sum
}

func (AOC) Day9_part1() int {
	lines := util.GetInput(9, false)

	disk := mountDisk(lines[0])
	compactDisk := compactDisk(disk)

	return getChecksum(compactDisk)
}

func (AOC) Day9_part2() int {
	// lines := util.GetInput(1, true)

	return 0
}
