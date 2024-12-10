package solutions

import (
	"fmt"
	"strings"

	"github.com/cauesmelo/aoc-2024/util"
)

func getDiagonalsRightToLeft(matrix []string) []string {
	diagonals := make([]string, 0)

	for col := 0; col < len(matrix); col++ {
		res := ""

		for line := 0; line < len(matrix); line++ {
			if col-line < 0 {
				break
			}

			res += string(matrix[line][col-line])
		}

		diagonals = append(diagonals, res)
	}

	for line := 1; line < len(matrix); line++ {
		res := ""

		for col := len(matrix) - 1; col >= 0; col-- {
			diff := len(matrix) - 1 - col

			if diff+line >= len(matrix) {
				break
			}

			res += string(matrix[line+diff][col])

		}

		diagonals = append(diagonals, res)
	}

	return diagonals
}

func getDiagonalLeftToRight(matrix []string) []string {
	diagonals := make([]string, 0)

	for i := len(matrix) - 1; i >= 0; i-- {
		res := ""

		for j := 0; j < len(matrix); j++ {
			if i+j >= len(matrix) {
				break
			}
			res += string(matrix[i+j][j])
		}

		diagonals = append(diagonals, res)
	}

	for column := 1; column < len(matrix); column++ {
		res := ""

		for line := 0; line < len(matrix); line++ {
			if column+line >= len(matrix) {
				break
			}

			res += string(matrix[line][column+line])
		}

		diagonals = append(diagonals, res)
	}

	return diagonals
}

func rotateMatrix(matrix []string) []string {
	size := len(matrix)
	newMatrix := make([]string, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newMatrix[j] += string(matrix[i][j])
		}
	}

	return newMatrix
}

func createDebugMatrix(size int) []string {
	matrix := make([]string, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i] += string(rune('A' + i*size + j))
		}
	}

	return matrix
}

func debugPrintMatrix(matrix []string, label string) {
	fmt.Printf("\n%s ==========", label)
	for _, line := range matrix {
		fmt.Printf("\n")
		for _, c := range line {
			fmt.Printf("%c\t", c)
		}
	}

	fmt.Printf("\n==========\n")
}

func countSubstr(s, pat string) int {
	var c int
	for d := range s {
		if strings.HasPrefix(s[d:], pat) {
			c++
		}
	}
	return c
}

func countXmas(line string) int {
	xmasCount := countSubstr(line, "XMAS")
	samxCount := countSubstr(line, "SAMX")

	return xmasCount + samxCount
}

func (AOC) Day4_part1() int {
	lines := util.GetInput(4, false)

	allLines := append([]string{}, lines...)
	allLines = append(allLines, rotateMatrix(lines)...)
	allLines = append(allLines, getDiagonalLeftToRight(lines)...)
	allLines = append(allLines, getDiagonalsRightToLeft(lines)...)

	sum := 0
	for _, line := range allLines {
		sum += countXmas(line)
	}

	return sum
}

func hasMASInX(matrix []string) bool {

	mCount := 0
	sCount := 0

	searchPositions := [][2]int{
		{0, 0},
		{0, 2},
		{2, 0},
		{2, 2},
	}

	for _, pos := range searchPositions {
		if matrix[pos[0]][pos[1]] == 'M' {
			mCount++
		}

		if matrix[pos[0]][pos[1]] == 'S' {
			sCount++
		}
	}

	if mCount != 2 || sCount != 2 {
		return false
	}

	sampleWord := string(matrix[0][0]) + string(matrix[1][1]) + string(matrix[2][2])

	if sampleWord == "SAS" || sampleWord == "MAM" {
		return false
	}

	return true
}

func countMASInX(matrix []string) int {

	count := 0

	maxIdx := len(matrix) - 1

	for line := 0; line < len(matrix); line++ {
		for col := 0; col < len(matrix); col++ {
			if line-1 < 0 || col-1 < 0 || line+1 > maxIdx || col+1 > maxIdx {
				continue
			}

			if string(matrix[line][col]) != "A" {
				continue
			}

			cutoutMatrix := make([]string, 0)
			for i := line - 1; i <= line+1; i++ {
				cutoutMatrix = append(cutoutMatrix, matrix[i][col-1:col+2])
			}

			if hasMASInX(cutoutMatrix) {
				count++
			}
		}
	}

	return count
}

func (AOC) Day4_part2() int {
	lines := util.GetInput(4, false)

	return countMASInX(lines)
}
