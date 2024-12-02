package day02

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Part1() (string, error) {
	file, err := os.OpenFile("solutions/day02/input.txt", os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}

	answer, puzzleErr := solvePart1(bufio.NewReader(file))
	return fmt.Sprintf("%d", answer), puzzleErr
}

func solvePart1(file *bufio.Reader) (uint32, error) {
	listScanner := bufio.NewScanner(file)
	safeCount := 0
	for listScanner.Scan() {
		levelsData := listScanner.Bytes()
		levelsRaw := bytes.Split(levelsData, []byte{' '})
		levels := encode(levelsRaw)

		if (isAllIncreasing(levels) || isAllDecreasing(levels)) && diffIsAtMostThree(levels) {
			safeCount++
		}
	}

	return uint32(safeCount), nil
}

func encode(rawLevels [][]byte) []uint8 {
	result := make([]uint8, len(rawLevels))
	var number uint8
	for index, level := range rawLevels {
		for _, b := range level {
			number = number*10 + b - '0'
		}
		result[index] = number
		number = uint8(0)
	}

	return result
}

func isAllIncreasing(report []uint8) bool {
	result := false
	previousValue := report[0]
	for index, value := range report {
		if index == 0 {
			continue
		}
		result = value > previousValue
		previousValue = value
		if result == false {
			break
		}
	}
	return result
}

func isAllDecreasing(report []uint8) bool {
	result := false
	previousValue := report[0]
	for index, value := range report {
		if index == 0 {
			continue
		}
		result = value < previousValue
		previousValue = value
		if result == false {
			break
		}
	}
	return result
}

func diffIsAtMostThree(report []uint8) bool {
	previousValue := report[0]
	diff := uint8(0)
	for index, value := range report {
		if index == 0 {
			continue
		}
		if value > previousValue {
			diff = value - previousValue
		} else {
			diff = previousValue - value
		}

		if diff > 3 {
			return false
		}
		previousValue = value
	}

	return true
}
