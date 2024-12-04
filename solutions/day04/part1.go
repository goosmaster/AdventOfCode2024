package day04

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func Part1() (string, error) {
	file, err := os.OpenFile("solutions/day04/input.txt", os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}

	answer, puzzleErr := solvePart1(bufio.NewReader(file))
	return fmt.Sprintf("%d", answer), puzzleErr
}

func solvePart1(file *bufio.Reader) (uint32, error) {
	memoryScanner := bufio.NewScanner(file)
	startBuffered := false
	total := uint32(0)
	var lines [][]byte
	for memoryScanner.Scan() {
		line := memoryScanner.Bytes()
		buff := []byte{'.', '.', '.', '.'}
		line = append(buff, line...)
		line = append(line, buff...)
		if startBuffered == false {
			var bufferedZone []byte
			for j := 0; j < len(line); j++ {
				bufferedZone = append(bufferedZone, '.')
			}
			for i := 0; i < 4; i++ {
				lines = append(lines, bufferedZone)
			}
			startBuffered = true
		}

		total += findHorizontal(line)
		slices.Reverse(line)
		total += findHorizontal(line)
		slices.Reverse(line)
		lines = append(lines, line)
	}

	var bufferedZone []byte
	for j := 0; j < len(lines[0]); j++ {
		bufferedZone = append(bufferedZone, '.')
	}
	for i := 0; i < 4; i++ {
		lines = append(lines, bufferedZone)
	}

	verticalLines := make([][]byte, len(lines[0]))
	for _, line := range lines {
		for i, character := range line {
			verticalLines[i] = append(verticalLines[i], character)
		}
	}
	for _, verticalLine := range verticalLines {
		total += findVertically(verticalLine)
		slices.Reverse(verticalLine)
		total += findVertically(verticalLine)
	}

	total += findDiagonally(lines)

	return total, nil
}

func findHorizontal(line []byte) uint32 {
	return countDuplicates(line, []byte{'X', 'M', 'A', 'S'})
}
func findVertically(line []byte) uint32 {
	return countDuplicates(line, []byte{'X', 'M', 'A', 'S'})
}
func findDiagonally(lines [][]byte) uint32 {
	count := uint32(0)
	for y, line := range lines {
		for x, character := range line {
			if character == 'X' {
				diagonalLines := constructDiagonalStrings(lines, y, x, 4)
				for _, diagonalLine := range diagonalLines {
					count += countDuplicates(diagonalLine, []byte{'X', 'M', 'A', 'S'})
				}
			}
		}
	}
	return count
}

func constructDiagonalStrings(lines [][]byte, y, x, length int) [][]byte {
	var diagonalLines [][]byte
	for dir := 0; dir <= 3; dir++ {
		var diagonalLine []byte
		for charOffset := 0; charOffset < length; charOffset++ {
			// out of bounds
			width := len(lines[0])
			height := len(lines)

			if dir == 0 { // left-up
				if y-length < 0 || x-length < 0 {
					continue
				}
				diagonalLine = append(diagonalLine, lines[y-charOffset][x-charOffset])
			} else if dir == 1 { // right-up
				if y-length < 0 || x+length > width {
					continue
				}
				diagonalLine = append(diagonalLine, lines[y-charOffset][x+charOffset])
			} else if dir == 2 { // left-down
				if y+length > height || x-length < 0 {
					continue
				}
				diagonalLine = append(diagonalLine, lines[y+charOffset][x-charOffset])
			} else if dir == 3 { // right-down
				if y+length > height || x+length > width {
					continue
				}
				diagonalLine = append(diagonalLine, lines[y+charOffset][x+charOffset])
			}
		}
		var duplicatedLine []byte
		diagonalLines = append(diagonalLines, diagonalLine)

		for i := len(diagonalLine) - 1; i >= 0; i-- {
			duplicatedLine = append(duplicatedLine, diagonalLine[i])
		}

		diagonalLines = append(diagonalLines, duplicatedLine)
	}

	return diagonalLines
}

func countDuplicates(haystack, needle []byte) uint32 {
	count := uint32(0)
	start := 0

	for {
		index := bytes.Index(haystack[start:], needle)
		if index == -1 {
			break
		}
		count++
		start += index + len(needle)
	}

	return count
}
