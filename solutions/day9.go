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

type slot struct {
	digit string
	size  int
	used  bool
}

func compactDiskv2(disk string) string {
	parsedDisk := strings.Split(disk, ";")
	parsedDisk = parsedDisk[:len(parsedDisk)-1]

	slots := make([]*slot, 0)

	for _, item := range parsedDisk {
		if len(slots) == 0 {
			slots = append(slots, &slot{
				digit: item,
				size:  len(item),
				used:  false,
			})

			continue
		}

		lastSlotDigit := slots[len(slots)-1].digit

		if item == lastSlotDigit {
			slots[len(slots)-1].size++
		} else {
			slots = append(slots, &slot{
				digit: item,
				size:  1,
				used:  false,
			})
		}
	}

	compactedSlots := make([]slot, 0)

	for sIdx := 0; sIdx < len(slots); sIdx++ {
		s := slots[sIdx]

		if s.digit != "." && !s.used {
			s.used = true
			compactedSlots = append(compactedSlots, *s)
			continue
		}

		freeSpace := s.size

		for i := len(slots) - 1; i > 0; i-- {
			if slots[i].used {
				continue
			}

			if slots[i].size > freeSpace {
				continue
			}

			if slots[i].digit == "." {
				continue
			}

			compactedSlots = append(compactedSlots, *slots[i])
			freeSpace -= slots[i].size
			slots[i].used = true

			if freeSpace <= 0 {
				break
			}
		}

		if freeSpace > 0 {
			compactedSlots = append(compactedSlots, slot{
				digit: ".",
				size:  freeSpace,
				used:  true,
			})
		}
	}

	compactDisk := ""

	for _, slot := range compactedSlots {
		for i := 0; i < slot.size; i++ {
			compactDisk += slot.digit + ";"
		}
	}

	return compactDisk
}

func mountDiskv2(digits string) string {
	disk := ""
	currN := 0
	isFile := true

	for i := 0; i < len(digits); i++ {
		nTimes, _ := strconv.Atoi(string(digits[i]))

		if nTimes > 0 {

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
		}

		isFile = !isFile
	}

	return disk
}

func (AOC) Day9_part2() int {
	lines := util.GetInput(9, false)

	disk := mountDiskv2(lines[0])
	compactDisk := compactDiskv2(disk)

	return getChecksum(compactDisk)
}
