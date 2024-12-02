package day02

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Part2() (string, error) {
	file, err := os.OpenFile("solutions/day02/input.txt", os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}

	answer, puzzleErr := solvePart2(bufio.NewReader(file))
	return fmt.Sprintf("%d", answer), puzzleErr
}

func solvePart2(file *bufio.Reader) (uint32, error) {
	listScanner := bufio.NewScanner(file)
	safeCount := 0
	for listScanner.Scan() {
		levelsData := listScanner.Bytes()
		levelsRaw := bytes.Split(levelsData, []byte{' '})
		levels := encode(levelsRaw)

		for index := range levels {
			isSorted := isAllIncreasingTolerated(levels, index) || isAllDecreasingTolerated(levels, index)
			if isSorted == false {
				continue
			}
			if diffIsAtMostThreeTolerated(levels, index) {
				safeCount++
				break
			}
		}
	}

	return uint32(safeCount), nil
}

func isAllIncreasingTolerated(report []uint8, skipIndex int) bool {
	result := false
	previousValue := report[0]
	first := 0
	if skipIndex == 0 {
		first = 1
		previousValue = report[1]
	}
	for index, value := range report {
		if index == first || index == skipIndex {
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

func isAllDecreasingTolerated(report []uint8, skipIndex int) bool {
	result := false
	previousValue := report[0]
	first := 0
	if skipIndex == 0 {
		first = 1
		previousValue = report[1]
	}
	for index, value := range report {
		if index == first || index == skipIndex {
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

func diffIsAtMostThreeTolerated(report []uint8, skipIndex int) bool {
	previousValue := report[0]
	first := 0
	if skipIndex == 0 {
		first = 1
		previousValue = report[1]
	}
	diff := uint8(0)
	for index, value := range report {
		if index == first || index == skipIndex {
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
