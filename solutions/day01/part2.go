package day01

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func Part2() (string, error) {
	file, err := os.OpenFile("solutions/day01/input.txt", os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}

	answer, puzzleErr := solvePart2(bufio.NewReader(file))
	return fmt.Sprintf("%d", answer), puzzleErr
}

func solvePart2(file *bufio.Reader) (uint32, error) {
	listScanner := bufio.NewScanner(file)
	lineIndex := uint16(0)
	leftList := make(map[uint16]uint32, 1000)
	rightList := make(map[uint16]uint32, 1000)

	for listScanner.Scan() {
		lists := listScanner.Bytes()
		listParts := bytes.Split(lists, []byte{' ', ' ', ' '})

		appendPart(lineIndex, &listParts[0], &leftList)
		appendPart(lineIndex, &listParts[1], &rightList)

		if lineIndex == math.MaxUint16 {
			return 0, fmt.Errorf("reached end of uint16")
		}
		lineIndex++
	}

	total := uint32(0)
	for _, leftValue := range leftList {
		count := uint32(0)
		for _, rightValue := range rightList {
			if leftValue == rightValue {
				count++
			}
		}
		total += leftValue * count
	}

	return total, nil
}
