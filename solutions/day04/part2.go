package day04

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() (string, error) {
	file, err := os.OpenFile("solutions/day04/input.txt", os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}

	answer, puzzleErr := solvePart2(bufio.NewReader(file))
	return fmt.Sprintf("%d", answer), puzzleErr
}

func solvePart2(file *bufio.Reader) (uint32, error) {
	memoryScanner := bufio.NewScanner(file)
	startBuffered := false
	total := uint32(0)
	var lines [][]byte
	for memoryScanner.Scan() {
		line := memoryScanner.Bytes()
		buff := []byte{'.', '.', '.'}
		line = append(buff, line...)
		line = append(line, buff...)
		if startBuffered == false {
			var bufferedZone []byte
			for j := 0; j < len(line); j++ {
				bufferedZone = append(bufferedZone, '.')
			}
			for i := 0; i < 3; i++ {
				lines = append(lines, bufferedZone)
			}
			startBuffered = true
		}

		lines = append(lines, line)
	}

	var bufferedZone []byte
	for j := 0; j < len(lines[0]); j++ {
		bufferedZone = append(bufferedZone, '.')
	}
	for i := 0; i < 3; i++ {
		lines = append(lines, bufferedZone)
	}

	total += findX(lines)

	return total, nil
}

func findX(lines [][]byte) uint32 {
	count := uint32(0)
	for y, line := range lines {
		for x, character := range line {
			if character == 'A' {
				count += constructX(lines, y, x)
			}
		}
	}
	return count
}

func constructX(lines [][]byte, y, x int) uint32 {
	count := uint32(0)

	if lines[y-1][x-1] == 'M' && lines[y+1][x+1] == 'S' && lines[y+1][x-1] == 'M' && lines[y-1][x+1] == 'S' {
		// from top left to bottom right MAS && from bottom left to top right MAS
		count++
	}
	if lines[y-1][x+1] == 'M' && lines[y+1][x-1] == 'S' && lines[y+1][x+1] == 'M' && lines[y-1][x-1] == 'S' {
		// from top right to bottom left MAS && from bottom right to top left MAS
		count++
	}

	if lines[y+1][x+1] == 'M' && lines[y-1][x-1] == 'S' && lines[y+1][x-1] == 'M' && lines[y-1][x+1] == 'S' {
		// from bottom right to top left MAS && from bottom left to top right MAS
		count++
	}
	if lines[y-1][x+1] == 'M' && lines[y+1][x-1] == 'S' && lines[y-1][x-1] == 'M' && lines[y+1][x+1] == 'S' {
		// from top right to bottom left MAS && from top left to bottom right MAS
		count++
	}

	return count
}
